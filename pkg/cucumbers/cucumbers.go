package cucumbers

import "aoc2021/pkg/io"

const (
	UNSET Tile = 0
	EMPTY Tile = '.'
	RIGHT Tile = '>'
	DOWN  Tile = 'v'
)

type Tile = rune
type Row = []Tile
type Map = []Row
type OceanFloor struct {
	tiles Map
}

func ReadOceanFloor(file string) OceanFloor {
	var of OceanFloor
	for _, line := range io.ReadLines(file) {
		row := make(Row, len(line))
		for x, t := range line {
			row[x] = t
		}
		of.tiles = append(of.tiles, row)
	}
	return of
}

func (of *OceanFloor) Step() bool {
	moved := false
	width := len(of.tiles[0])
	height := len(of.tiles)

	newtiles := make(Map, height)
	for y := 0; y < height; y++ {
		newtiles[y] = make(Row, width)
		for x := 0; x < width; x++ {
			newtiles[y][x] = EMPTY
		}
	}

	for y, row := range of.tiles {
		for x, t := range row {
			if t == RIGHT {
				if of.tiles[y][(x+1)%width] == EMPTY {
					newtiles[y][(x+1)%width] = RIGHT
					moved = true
				} else {
					newtiles[y][x] = RIGHT
				}
			}
		}
	}
	for y, row := range of.tiles {
		for x, t := range row {
			if t == DOWN {
				if newtiles[(y+1)%height][x] == EMPTY && of.tiles[(y+1)%height][x] != DOWN {
					newtiles[(y+1)%height][x] = DOWN
					moved = true
				} else {
					newtiles[y][x] = DOWN
				}
			}
		}
	}
	of.tiles = newtiles
	return moved
}

func (of *OceanFloor) String() string {
	str := ""
	for _, row := range of.tiles {
		for _, t := range row {
			str += string(t)
		}
		str += "\n"
	}
	return str
}
