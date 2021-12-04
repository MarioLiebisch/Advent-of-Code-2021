package main

import "aoc2021/pkg/bingo"

func main() {
	println("Solution 1 - Sample:", bingo.Play("./data/sample-4.txt"))
	println("Solution 1 - Input:", bingo.Play("./data/input-4.txt"))

	println("Solution 2 - Sample:", bingo.Fail("./data/sample-4.txt"))
	println("Solution 2 - Input:", bingo.Fail("./data/input-4.txt"))
}
