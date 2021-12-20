package trench

import (
	"aoc2021/pkg/io"
	"aoc2021/pkg/numbers"
)

type ScannerImagePixel bool
type ScannerImageRow map[int]ScannerImagePixel
type ScannerImage map[int]ScannerImageRow

type ScannerData struct {
	algorithm    []ScannerImagePixel
	image        []ScannerImage
	page         int
	num_pages    int
	minx, maxx   int
	miny, maxy   int
	nminx, nmaxx int
	nminy, nmaxy int
	inf_false    ScannerImagePixel
	inf_true     ScannerImagePixel
	inf_value    ScannerImagePixel
}

func ReadScannerData(file string) ScannerData {
	var sd ScannerData
	sd.page = 0
	sd.num_pages = 2

	lines := io.ReadLines(file)

	for _, v := range lines[0] {
		if v == '#' {
			sd.algorithm = append(sd.algorithm, true)
		} else {
			sd.algorithm = append(sd.algorithm, false)
		}
	}
	sd.image = make([]ScannerImage, sd.num_pages)
	for i := 0; i < sd.num_pages; i++ {
		sd.image[i] = make(ScannerImage)
	}
	j := 0
	for i := 2; i < len(lines); i++ {
		row := make(ScannerImageRow)
		for x, v := range lines[i] {
			if v == '#' {
				row[x] = true
			}
		}
		sd.image[sd.page][j] = row
		j++
	}

	sd.inf_false = sd.algorithm[0]
	sd.inf_true = sd.algorithm[len(sd.algorithm)-1]

	sd.maxx = len(lines[2])
	sd.nmaxx = sd.maxx
	sd.maxy = len(lines) - 2
	sd.nmaxy = sd.maxy

	return sd
}

func (sd *ScannerData) Get(x, y int) ScannerImagePixel {
	if x < sd.minx || x > sd.maxx || y < sd.miny || y > sd.maxy {
		return sd.inf_value
	}
	if sd.image[sd.page][y] == nil {
		return false
	}
	return sd.image[sd.page][y][x]
}

func (sd *ScannerData) Set(x, y int, value ScannerImagePixel) {
	if value {
		if x < sd.nminx {
			sd.nminx = x
		}
		if x > sd.nmaxx {
			sd.nmaxx = x
		}
		if y < sd.nminy {
			sd.nminy = y
		}
		if y > sd.nmaxy {
			sd.nmaxy = y
		}

		if sd.image[(sd.page+1)%sd.num_pages][y] == nil {
			sd.image[(sd.page+1)%sd.num_pages][y] = make(ScannerImageRow)
		}
		sd.image[(sd.page+1)%sd.num_pages][y][x] = value
	}
}

func (sd *ScannerData) Flip() {
	sd.image[sd.page] = make(ScannerImage)
	sd.page = (sd.page + 1) % sd.num_pages
	sd.minx = sd.nminx
	sd.maxx = sd.nmaxx
	sd.miny = sd.nminy
	sd.maxy = sd.nmaxy
	if sd.inf_value {
		sd.inf_value = sd.inf_true
	} else {
		sd.inf_value = sd.inf_false
	}
}

func (sd *ScannerData) Reconstruct(iterations int) {
	for i := 0; i < iterations; i++ {
		for y := sd.miny - 1; y <= sd.maxy+1; y++ {
			for x := sd.minx - 1; x <= sd.maxx+1; x++ {
				idx := ""
				for _y := y - 1; _y <= y+1; _y++ {
					for _x := x - 1; _x <= x+1; _x++ {
						if sd.Get(_x, _y) {
							idx += "1"
						} else {
							idx += "0"
						}
					}
				}
				sd.Set(x, y, sd.algorithm[numbers.BinToInt(idx)])
			}
		}
		sd.Flip()
	}
}

func (sd *ScannerData) String() string {
	str := ""
	for y := sd.miny; y <= sd.maxy; y++ {
		for x := sd.minx; x <= sd.maxx; x++ {
			if sd.Get(x, y) {
				str += "#"
			} else {
				str += "."
			}
		}
		str += "\n"
	}
	return str
}

func (sd *ScannerData) Count() int {
	count := 0
	for _, row := range sd.image[sd.page] {
		count += len(row)
	}
	return count
}
