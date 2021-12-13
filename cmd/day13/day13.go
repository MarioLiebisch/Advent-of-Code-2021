package main

import "aoc2021/pkg/manual"

func main() {
	papers := manual.ReadPaper("./data/sample-13.txt", true)
	paperi := manual.ReadPaper("./data/input-13.txt", true)

	println("Solution 1 - Sample:", papers.Count())
	println("Solution 1 - Input:", paperi.Count())

	println("Solution 2:")
	paper := manual.ReadPaper("./data/input-13.txt", false)
	paper.Print()
}
