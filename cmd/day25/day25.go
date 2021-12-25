package main

import "aoc2021/pkg/cucumbers"

func main() {
	ofs := cucumbers.ReadOceanFloor("./data/sample-25.txt")
	ofi := cucumbers.ReadOceanFloor("./data/input-25.txt")

	for n := 1; ; n++ {
		if !ofs.Step() {
			println("Solution 1 - Sample:", n)
			break
		}
	}

	for n := 1; ; n++ {
		if !ofi.Step() {
			println("Solution 1 - Input:", n)
			break
		}
	}

}
