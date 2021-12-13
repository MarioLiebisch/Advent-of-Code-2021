package manual

import (
	"aoc2021/pkg/io"
	"strconv"
	"strings"
)

type Paper map[int]map[int]bool

func ReadPaper(file string, fold_once bool) Paper {
	paper := make(Paper)
	for _, line := range io.ReadLines(file) {
		if line == "" {
			continue
		}
		if line[0] == 'f' { // fold
			axis := line[11]
			index, _ := strconv.Atoi(line[13:])
			if axis == 'x' {
				for _, row := range paper {
					for x := 0; x < index; x++ {
						if row[2*index-x] {
							row[x] = true
						}
					}
					for x := range row {
						if x > index {
							delete(row, x)
						}
					}
				}
			} else if axis == 'y' {
				for y := 0; y < index; y++ {
					fold := paper[2*index-y]
					if fold != nil {
						if paper[y] == nil {
							paper[y] = make(map[int]bool)
						}
						for x, e := range fold {
							if e {
								paper[y][x] = true
							}
						}
						delete(paper, 2*index-y)
					}
				}
			}
			if fold_once {
				return paper
			}
			continue
		}
		data := strings.SplitN(line, ",", 2)
		x, _ := strconv.Atoi(data[0])
		y, _ := strconv.Atoi(data[1])
		if paper[y] == nil {
			paper[y] = make(map[int]bool)
		}
		paper[y][x] = true
	}
	return paper
}

func (p *Paper) Print() {
	width := 0
	height := 0
	for y := range *p {
		if y >= height {
			height = y + 1
		}
		for x := range (*p)[y] {
			if x >= width {
				width = x + 1
			}
		}
	}
	out := ""
	for y := 0; y < height; y++ {
		col := (*p)[y]
		for x := 0; x < width; x++ {
			if col[x] {
				out += "#"
			} else {
				out += "."
			}
		}
		out += "\n"
	}
	println(out)
}

func (p *Paper) Count() int {
	count := 0
	for _, row := range *p {
		for _, e := range row {
			if e {
				count++
			}
		}
	}
	return count
}
