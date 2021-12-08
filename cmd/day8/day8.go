package main

import "aoc2021/pkg/wiring"

func main() {
	sample := wiring.ReadWiringSets("./data/sample-8.txt")
	input := wiring.ReadWiringSets("./data/input-8.txt")

	println("Solution 1 - Sample:", sample.CountPartOne())
	println("Solution 2 - Input:", input.CountPartOne())

	println("Solution 2 - Sample:", sample.SumPartTwo())
	println("Solution 2 - Input:", input.SumPartTwo())
}
