package main

import "aoc2021/pkg/octopuses"

func main() {
	sample1 := octopuses.ReadFile("./data/sample-11.txt")
	input1 := octopuses.ReadFile("./data/input-11.txt")
	sample2 := octopuses.ReadFile("./data/sample-11.txt")
	input2 := octopuses.ReadFile("./data/input-11.txt")

	println("Solution 1 - Sample:", sample1.Steps(100))
	println("Solution 1 - Input:", input1.Steps(100))

	println("Solution 2 - Sample:", sample2.FullFlash())
	println("Solution 2 - Input:", input2.FullFlash())
}
