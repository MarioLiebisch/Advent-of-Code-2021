package main

import (
	"aoc2021/pkg/air"
	"aoc2021/pkg/io"
	"aoc2021/pkg/power"
)

func main() {
	d3s := io.ReadLines("data/sample-3.txt")
	d3i := io.ReadLines("data/input-3.txt")

	println("Solution 1 - Sample:", power.GetPowerConsumption(d3s))
	println("Solution 1 - Input:", power.GetPowerConsumption(d3i))

	println("Solution 2 - Sample:", air.GetLifeSupportRating(d3s))
	println("Solution 2 - Input:", air.GetLifeSupportRating(d3i))
}
