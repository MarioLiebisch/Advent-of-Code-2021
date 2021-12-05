package vents

import (
	"aoc2021/pkg/io"
	"aoc2021/pkg/numbers"
	"fmt"
	"strconv"
	"strings"
)

type Vent struct {
	Start numbers.Vector2
	End   numbers.Vector2
}

type Vents []Vent

type VentMap map[int]map[int]int

func ReadVents(file string) Vents {
	var vents Vents
	for _, line := range io.ReadLines(file) {
		vents = append(vents, FromLine(line))
	}
	return vents
}

func FromLine(line string) Vent {
	points := strings.SplitN(line, " -> ", 2)
	start := strings.SplitN(points[0], ",", 2)
	end := strings.SplitN(points[1], ",", 2)
	var vent Vent
	vent.Start.X, _ = strconv.Atoi(start[0])
	vent.Start.Y, _ = strconv.Atoi(start[1])
	vent.End.X, _ = strconv.Atoi(end[0])
	vent.End.Y, _ = strconv.Atoi(end[1])
	return vent
}

func (v *Vent) IsDiagonal() bool {
	return v.Start.X != v.End.X && v.Start.Y != v.End.Y
}

func (vm *VentMap) AddVents(vs Vents, diagonals bool) {
	for _, v := range vs {
		if diagonals || !v.IsDiagonal() {
			vm.AddVent(v)
		}
	}
}

func (vm *VentMap) Print() {
	out := ""
	xs, ys, xe, ye := 0, 0, 0, 0
	for k, row := range *vm {
		ys = min(ys, k)
		ye = max(ye, k)
		for l := range row {
			xs = min(xs, l)
			xe = max(xe, l)
		}
	}

	for y := ys; y <= ye; y++ {
		if (*vm)[y] == nil {
			for x := xs; x <= xe; x++ {
				out += "0 "
			}
		} else {
			for x := xs; x <= xe; x++ {
				out += fmt.Sprintf("%d ", (*vm)[y][x])
			}
		}
		out += "\n"
	}
	print(out)
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// This could be done more simple, but I had
// the floating point approach for any diagonal
// line, then I get to part two and it's only 45 degree…
func (vm *VentMap) AddVent(v Vent) {
	dx := v.End.X - v.Start.X
	dy := v.End.Y - v.Start.Y
	steps := max(abs(dx), abs(dy))
	dx /= steps
	dy /= steps
	for i := 0; i <= steps; i++ {
		x := v.Start.X + dx*i
		y := v.Start.Y + dy*i
		if (*vm)[y] == nil {
			(*vm)[y] = make(map[int]int)
		}
		(*vm)[y][x]++
	}
}

func (vm *VentMap) CountCrossings() int {
	count := 0
	for _, row := range *vm {
		for _, entry := range row {
			if entry > 1 {
				count++
			}
		}
	}
	return count
}
