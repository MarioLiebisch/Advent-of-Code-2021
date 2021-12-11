package octopuses

import (
	"aoc2021/pkg/io"
	"aoc2021/pkg/numbers"
	"fmt"
)

type Octopus int
type Octopuses []Octopus
type OctoGrid []Octopuses

func ReadFile(file string) OctoGrid {
	var grid OctoGrid
	for _, line := range io.ReadLines(file) {
		var octos Octopuses
		for _, v := range line {
			octos = append(octos, Octopus(v-'0'))
		}
		grid = append(grid, octos)
	}
	return grid
}

func (grid *OctoGrid) Print() {
	height := len(*grid)
	width := len((*grid)[0])
	ret := ""
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			ret += fmt.Sprint((*grid)[y][x])
		}
		ret += "\n"
	}
	println(ret)
}

func (grid *OctoGrid) Inject(x, y int) {
	height := len(*grid)
	width := len((*grid)[0])
	value := (*grid)[y][x]

	// Negative value - flashed already
	if value < 0 {
		return
	}

	// Just increase the energy level, if too low
	if value < 9 {
		(*grid)[y][x] = value + 1
		return
	}

	// Set to flashed and propagate the energy
	(*grid)[y][x] = -1
	for _y := numbers.Max(y-1, 0); _y <= numbers.Min(y+1, height-1); _y++ {
		for _x := numbers.Max(x-1, 0); _x <= numbers.Min(x+1, width-1); _x++ {
			if x != _x || y != _y {
				grid.Inject(_x, _y)
			}
		}
	}
}

func (grid *OctoGrid) Step() int {
	flashes := 0
	height := len(*grid)
	width := len((*grid)[0])

	// Inject energy into all of them
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			grid.Inject(x, y)
		}
	}

	// "Burn out" all that flashed
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if (*grid)[y][x] < 0 {
				(*grid)[y][x] = 0
				flashes++
			}
		}
	}

	return flashes
}

func (grid *OctoGrid) Steps(steps int) int {
	flashes := 0

	for n := 0; n < steps; n++ {
		flashes += grid.Step()
	}

	return flashes
}

func (grid *OctoGrid) FullFlash() int {
	height := len(*grid)
	width := len((*grid)[0])
	count := width * height
	for step := 1; ; step++ {
		if count == grid.Step() {
			return step
		}
	}
}
