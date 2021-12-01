package main

import (
	"aoc2021/pkg/io"
	"aoc2021/pkg/numbers"
)

func main() {
	println("Solution 1 - Sample:", numbers.CountIncrements(io.ReadIntegers("data/sample-1.txt")))
	println("Solution 1 - Input:", numbers.CountIncrements(io.ReadIntegers("data/input-1.txt")))

	println("Solution 2 - Sample:", numbers.CountIncrements(numbers.MergeWindow(io.ReadIntegers("data/sample-1.txt"), 3)))
	println("Solution 2 - Input:", numbers.CountIncrements(numbers.MergeWindow(io.ReadIntegers("data/input-1.txt"), 3)))
}
