package height

import (
	"aoc2021/pkg/io"
	"aoc2021/pkg/numbers"
	"sort"
)

type HeightMap [][]int

type LowPoints []numbers.Vector2

func Load(file string) HeightMap {
	var hm HeightMap
	for _, line := range io.ReadLines(file) {
		var l []int
		for _, c := range line {
			l = append(l, int(c-'0'))
		}
		hm = append(hm, l)
	}
	return hm
}

func (hm *HeightMap) IsLowPoint(x, y int) bool {
	current := (*hm)[y][x]
	left := 10
	right := 10
	top := 10
	bottom := 10
	if x > 0 {
		left = (*hm)[y][x-1]
	}
	if y > 0 {
		top = (*hm)[y-1][x]
	}
	if x < len((*hm)[0])-1 {
		right = (*hm)[y][x+1]
	}
	if y < len(*hm)-1 {
		bottom = (*hm)[y+1][x]
	}
	if left <= current || right <= current || top <= current || bottom <= current {
		return false
	}
	return true
}

func (hm *HeightMap) LowPoints() (LowPoints, int) {
	height := len(*hm)
	width := len((*hm)[0])
	var pos LowPoints
	var risk int
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if hm.IsLowPoint(x, y) {
				pos = append(pos, numbers.Vector2{
					X: x,
					Y: y,
				})
				risk += (*hm)[y][x] + 1
			}
		}
	}

	return pos, risk
}

func (hm *HeightMap) FindBasins() []int {
	height := len(*hm)
	width := len((*hm)[0])
	basinmap := make([][]int, height)
	for i := range basinmap {
		basinmap[i] = make([]int, width)
	}

	// Not the most efficient approach, but let's just stick with flood fill

	// Fill initial low poins
	lps, _ := hm.LowPoints()
	for i, lp := range lps {
		// Save index + 1, so we have 0 for empty
		basinmap[lp.Y][lp.X] = i + 1
	}

	for flow := true; flow; {
		flow = false
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				// Skip borders and unfilled locations
				if (*hm)[y][x] == 9 || basinmap[y][x] == 0 {
					continue
				}
				// Fill to top
				if y > 0 && (*hm)[y-1][x] != 9 && basinmap[y-1][x] == 0 {
					basinmap[y-1][x] = basinmap[y][x]
					flow = true
				}
				// Fill to bottom
				if y < height-1 && (*hm)[y+1][x] != 9 && basinmap[y+1][x] == 0 {
					basinmap[y+1][x] = basinmap[y][x]
					flow = true
				}
				// Fill to left
				if x > 0 && (*hm)[y][x-1] != 9 && basinmap[y][x-1] == 0 {
					basinmap[y][x-1] = basinmap[y][x]
					flow = true
				}
				// Fill to bottom
				if x < width-1 && (*hm)[y][x+1] != 9 && basinmap[y][x+1] == 0 {
					basinmap[y][x+1] = basinmap[y][x]
					flow = true
				}
			}
		}
		if !flow {
			break
		}
	}

	// Count areas per ID
	basins := make([]int, len(lps))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			id := basinmap[y][x] - 1
			if id >= 0 {
				basins[id]++
			}
		}
	}
	sort.Slice(basins, func(i, j int) bool { return basins[i] > basins[j] })
	return basins
}
