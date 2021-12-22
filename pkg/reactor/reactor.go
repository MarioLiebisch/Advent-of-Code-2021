package reactor

import (
	"aoc2021/pkg/io"
	"aoc2021/pkg/numbers"
	"strconv"
	"strings"
)

type Cuboid struct {
	from, to numbers.Vector3
	toggle   bool
}

func (c *Cuboid) Volume() int {
	return (c.to.X - c.from.X + 1) * (c.to.Y - c.from.Y + 1) * (c.to.Z - c.from.Z + 1)
}

func (c *Cuboid) Intersect(other Cuboid) *Cuboid {
	if c.from.X > other.to.X || c.to.X < other.from.X || c.from.Y > other.to.Y || c.to.Y < other.from.Y || c.from.Z > other.to.Z || c.to.Z < other.from.Z {
		return nil
	}
	return &Cuboid{
		from: numbers.Vector3{
			X: numbers.Max(c.from.X, other.from.X),
			Y: numbers.Max(c.from.Y, other.from.Y),
			Z: numbers.Max(c.from.Z, other.from.Z),
		},
		to: numbers.Vector3{
			X: numbers.Min(c.to.X, other.to.X),
			Y: numbers.Min(c.to.Y, other.to.Y),
			Z: numbers.Min(c.to.Z, other.to.Z),
		},
	}
}

type Reactor []Cuboid

type Instruction = Cuboid

type Instructions []Instruction

func LoadInstructions(file string) Instructions {
	var is Instructions
	for _, line := range io.ReadLines(file) {
		var i Instruction
		var parts []string
		if line[1] == 'n' {
			i.toggle = true
			parts = strings.SplitN(line[3:], ",", 3)
		} else {
			parts = strings.SplitN(line[4:], ",", 3)
		}
		xparts := strings.SplitN(parts[0][2:], "..", 2)
		yparts := strings.SplitN(parts[1][2:], "..", 2)
		zparts := strings.SplitN(parts[2][2:], "..", 2)

		i.from.X, _ = strconv.Atoi(xparts[0])
		i.to.X, _ = strconv.Atoi(xparts[1])
		i.from.Y, _ = strconv.Atoi(yparts[0])
		i.to.Y, _ = strconv.Atoi(yparts[1])
		i.from.Z, _ = strconv.Atoi(zparts[0])
		i.to.Z, _ = strconv.Atoi(zparts[1])

		is = append(is, i)
	}
	return is
}

func (r *Reactor) Apply(instructions Instructions, init bool) {
	for _, instruction := range instructions {
		// Ignore non-initialization steps?
		if init && (instruction.from.X < -50 || instruction.to.X > 50 || instruction.from.Y < -50 || instruction.to.Y > 50 || instruction.from.Z < -50 || instruction.to.Z > 50) {
			continue
		}

		// Save current cuboid count to avoid
		// handling newly added ones
		rcount := len(*r)
		// If this is turning things on, add it
		if instruction.toggle {
			*r = append(*r, instruction)
		}
		// Look for overlaps with previous cuboids
		for i := 0; i < rcount; i++ {
			intersection := instruction.Intersect((*r)[i])
			if intersection != nil {
				// Negate the instruction to account for the
				// overlap (e.g. avoid adding/removing twice)
				intersection.toggle = !(*r)[i].toggle
				*r = append(*r, *intersection)
			}
		}
	}
}

func (r *Reactor) GetCount() int {
	count := 0
	for _, c := range *r {
		if c.toggle {
			count += c.Volume()
		} else {
			count -= c.Volume()
		}
	}
	return count
}
