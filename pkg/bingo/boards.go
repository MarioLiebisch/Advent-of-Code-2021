package bingo

import (
	"strconv"
	"strings"
)

type Draws []int

type BoardField struct {
	Value  int
	Ticked bool
}

type BoardRow []BoardField
type Board []BoardRow
type Boards []Board

func (d *Draws) Fill(input string) int {
	*d = nil
	for _, v := range strings.Split(input, ",") {
		iv, _ := strconv.Atoi(v)
		*d = append(*d, iv)
	}
	return len(*d)
}

func (r *BoardRow) Tick(number int) {
	for i := 0; i < len(*r); i++ {
		if (*r)[i].Value == number {
			(*r)[i].Ticked = true
		}
	}
}

func (b *Board) Tick(number int) {
	for i := 0; i < len(*b); i++ {
		(*b)[i].Tick(number)
	}
}

func (bs *Boards) Tick(number int) {
	for i := 0; i < len(*bs); i++ {
		(*bs)[i].Tick(number)
	}
}

func (bs *Boards) FindWinner() int {
	height := len((*bs)[0])
	width := len((*bs)[0][0])
	for i := 0; i < len(*bs); i++ {
		// Check  rows
		for y := 0; y < height; y++ {
			winner := true
			for x := 0; x < width; x++ {
				if !(*bs)[i][y][x].Ticked {
					winner = false
					break
				}
			}
			if winner {
				// We've got a winner
				// Don't bother finding another one
				return i
			}
		}
		// Check columns
		for x := 0; x < width; x++ {
			winner := true
			for y := 0; y < height; y++ {
				if !(*bs)[i][y][x].Ticked {
					winner = false
					break
				}
			}
			if winner {
				// We've got a winner
				// Don't bother finding another one
				return i
			}
		}
	}
	return -1
}

func (b *Board) GetScore() int {
	height := len(*b)
	width := len((*b)[0])
	score := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if !(*b)[y][x].Ticked {
				score += (*b)[y][x].Value
			}
		}
	}
	return score
}
