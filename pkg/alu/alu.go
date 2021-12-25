package alu

import (
	"aoc2021/pkg/io"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type operation byte
type operandtype byte

type operand struct {
	value      int64
	identifier byte
}

func (o *operand) set(identifier byte) {
	o.identifier = identifier
	o.value = 0
}

func (o *operand) setValue(value int64) {
	o.identifier = 0
	o.value = value
}

type instruction struct {
	op     operation
	v1, v2 operand
}

type Program []instruction

const (
	inp operation = iota
	add
	mul
	div
	mod
	eql
	res
	neq
	mov
	m26
	d26
	m25
	inc
)

func CompileProgram(file string, optimize bool) Program {
	var p Program

	varset := make(map[byte]bool)
	varset[0] = true // literal values are always set

	for _, line := range io.ReadLines(file) {
		var ins instruction
		data := strings.SplitN(line, " ", 3)
		switch data[0] {
		case "inp":
			ins.op = inp
		case "add":
			ins.op = add
		case "mul":
			ins.op = mul
		case "div":
			ins.op = div
		case "mod":
			ins.op = mod
		case "eql":
			ins.op = eql
		}
		switch data[1] {
		case "w":
			ins.v1.set('w')
		case "x":
			ins.v1.set('x')
		case "y":
			ins.v1.set('y')
		case "z":
			ins.v1.set('z')
		}

		if len(data) > 2 {
			switch data[2] {
			case "w":
				ins.v2.set('w')
			case "x":
				ins.v2.set('x')
			case "y":
				ins.v2.set('y')
			case "z":
				ins.v2.set('z')
			default:
				v, _ := strconv.Atoi(data[2])
				ins.v2.setValue(int64(v))
			}
		} else {
			ins.v2.set(' ')
		}

		// Optimization (only single-pass for now)
		// Next time I should think abou whether optimizing is even the correct approach,
		// given the fact you're supposed to try 14 digit numbers…
		if optimize {
			if ins.op == mod && ins.v2.value == 26 {
				ins.op = m26
				ins.v2.set(' ')
			} else if ins.op == div && ins.v2.value == 26 {
				ins.op = d26
				ins.v2.set(' ')
			} else if ins.op == add && ins.v2.value == 1 {
				ins.op = inc
				ins.v2.set(' ')
			}

			if !varset[ins.v2.identifier] || (ins.v2.identifier == 0 && ins.v2.value == 0) {
				// Second operand is zero or not set yet
				switch ins.op {
				case add:
					continue // do nothing
				case mul:
					// remove all previous lines until the variable is actually used
					if varset[ins.v1.identifier] {
						for i := len(p) - 1; i >= 0; i-- {
							tins := p[i]
							if tins.v2.identifier == ins.v1.identifier {
								break
							} else if tins.v1.identifier == ins.v1.identifier {
								// This assignment is negated
								newp := make(Program, 0)
								newp = append(newp, p[:i]...)
								if i != len(p) {
									newp = append(newp, p[i+1:]...)
								}
								p = newp
							}
							if i == 0 {
								// Reached the start without this being used
								varset[ins.v1.identifier] = false
							}
						}
					}

					if varset[ins.v1.identifier] {
						ins.op = res // reset
						varset[ins.v1.identifier] = false
					} else {
						continue
					}
				case eql:
					last := p[len(p)-1]
					if last.v1.identifier == ins.v1.identifier {
						// This negates the previous comparison
						if last.op == eql {
							p[len(p)-1].op = neq
							// Special case:
							// Comparing input w with previous "out of bounds" value
							if len(p) > 2 {
								last2 := p[len(p)-2]
								if last2.op == mov && last2.v1.identifier == last.v1.identifier && last2.v2.identifier == 0 && last2.v2.value > 9 {
									p = p[:len(p)-1]
									p[len(p)-1].op = mov
									p[len(p)-1].v2.value = 1
								}
							}
							continue
						} else if last.op == neq {
							p[len(p)-1].op = eql
							continue
						}
					}
				}
			} else if ins.v2.identifier == 0 && ins.v2.value == 1 {
				// Second operand is one
				switch ins.op {
				case mul:
					continue // do nothing
				case div:
					continue // do nothing
				case mod:
					ins.op = res // reset
				}
			}

			if !varset[ins.v1.identifier] {
				// First operand not set
				switch ins.op {
				case mul:
					continue // do nothing
				case mod:
					continue // do nothing
				case m26:
					continue // do nothing
				case div:
					continue // do nothing
				case d26:
					continue // do nothing
				case add:
					// This is only "move"
					// Remove a preceding reset (if any)
					if len(p) > 1 {
						last := p[len(p)-1]
						if last.op == res && last.v1.identifier == ins.v1.identifier {
							if ins.v2.value == 25 {
								p[len(p)-1].op = m25
								p[len(p)-1].v2.set(' ')
							} else {
								p[len(p)-1].op = mov
								p[len(p)-1].v2 = ins.v2
							}
							varset[ins.v1.identifier] = true
							continue
						}
					}
					ins.op = mov
				}
			}

			if len(p) > 1 {
				last := p[len(p)-1]
				switch ins.op {
				case add:
					// Merge add/move with values
					if (last.op == add || last.op == mov) && last.v2.identifier == 0 && ins.v2.identifier == 0 {
						last.v2.value += ins.v2.value
						continue
					}
				}
			}

			varset[ins.v1.identifier] = ins.op != res
		}

		p = append(p, ins)
	}
	return p
}

func (o *operand) String() string {
	if o.identifier == 0 {
		return fmt.Sprint(o.value)
	}
	return string(o.identifier)
}

func (o *operation) String() string {
	switch *o {
	case inp:
		return "inp"
	case add:
		return "add"
	case mul:
		return "mul"
	case div:
		return "div"
	case mod:
		return "mod"
	case eql:
		return "eql"
	case res:
		return "res*"
	case neq:
		return "neq*"
	case mov:
		return "mov*"
	case m26:
		return "m26*"
	case d26:
		return "d26*"
	case m25:
		return "m25*"
	case inc:
		return "inc*"
	default:
		return "Missing Operand String (" + fmt.Sprint(*o) + ")!"
	}
}

func (p *Program) Listing() string {
	lst := ""
	for _, ins := range *p {
		lst += fmt.Sprintf("%s %s %s\n", ins.op.String(), ins.v1.String(), ins.v2.String())
	}
	return lst
}

type State struct {
	W, X, Y, Z  int64
	ins_pointer int
	inp_pointer int
}

type ALU struct {
	state   State
	input   string
	program Program
	halt    bool
}

func (alu *ALU) get(o operand) int64 {
	switch o.identifier {
	case 'w':
		return alu.state.W
	case 'x':
		return alu.state.X
	case 'y':
		return alu.state.Y
	case 'z':
		return alu.state.Z
	default:
		return o.value
	}
}

func (alu *ALU) set(o *operand, value int64) {
	switch o.identifier {
	case 'w':
		alu.state.W = value
	case 'x':
		alu.state.X = value
	case 'y':
		alu.state.Y = value
	case 'z':
		alu.state.Z = value
	default:
		o.identifier = 0
		o.value = value
	}
}

func (alu *ALU) run() {
	ins := alu.program[alu.state.ins_pointer]
	switch ins.op {
	case inp:
		if alu.state.inp_pointer < len(alu.input) {
			alu.set(&ins.v1, int64(alu.input[alu.state.inp_pointer]-'0'))
		} else {
			alu.halt = true
			return
		}
		alu.state.inp_pointer++
	case add:
		alu.set(&ins.v1, alu.get(ins.v1)+alu.get(ins.v2))
	case mul:
		alu.set(&ins.v1, alu.get(ins.v1)*alu.get(ins.v2))
	case div:
		alu.set(&ins.v1, alu.get(ins.v1)/alu.get(ins.v2))
	case mod:
		alu.set(&ins.v1, alu.get(ins.v1)%alu.get(ins.v2))
	case eql:
		if alu.get(ins.v1) == alu.get(ins.v2) {
			alu.set(&ins.v1, 1)
		} else {
			alu.set(&ins.v1, 0)
		}
	case neq:
		if alu.get(ins.v1) != alu.get(ins.v2) {
			alu.set(&ins.v1, 1)
		} else {
			alu.set(&ins.v1, 0)
		}
	case res: // Custom: Reset
		alu.set(&ins.v1, 0)
	case mov: // Custom: Move
		alu.set(&ins.v1, alu.get(ins.v2))
	case m26: // Custom: Mod 26
		alu.set(&ins.v1, alu.get(ins.v1)%26)
	case d26: // Custom: Div 26
		alu.set(&ins.v1, alu.get(ins.v1)/26)
	case m25:
		alu.set(&ins.v1, 25)
	case inc:
		alu.set(&ins.v1, alu.get(ins.v1)+1)
	default:
		log.Fatalln("Instruction not implemented!")
	}
}

func (alu *ALU) Run(program Program, input string, state *State) State {
	if state == nil {
		alu.state.W, alu.state.X, alu.state.Y, alu.state.Z = 0, 0, 0, 0
	} else {
		alu.state = *state
	}
	alu.state.inp_pointer = 0
	alu.input = input
	alu.program = program
	alu.halt = false
	for alu.state.ins_pointer = 0; alu.state.ins_pointer < len(program) && !alu.halt; alu.state.ins_pointer++ {
		alu.run()
	}
	return alu.state
}
