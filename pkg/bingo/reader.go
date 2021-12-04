package bingo

import (
	"aoc2021/pkg/io"
	"strconv"
	"strings"
)

func Read(file string) (Draws, Boards) {
	var boards Boards
	var draws Draws

	lines := io.ReadLines(file)

	draws.Fill(lines[0])

	var board Board
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			if len(board) > 0 {
				boards = append(boards, board)
				board = nil
			}
		} else {
			var row BoardRow
			for _, v := range strings.Split(line, " ") {
				if v != "" {
					iv, _ := strconv.Atoi(v)
					row = append(row, BoardField{
						Value:  iv,
						Ticked: false,
					})
				}
			}
			board = append(board, row)
		}
	}
	if len(board) > 0 {
		boards = append(boards, board)
	}

	return draws, boards
}
