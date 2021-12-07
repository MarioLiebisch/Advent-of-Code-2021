package main

import "aoc2021/pkg/crabs"

func main() {
	crabss := crabs.LoadCrabs("./data/sample-7.txt")
	crabsi := crabs.LoadCrabs("./data/input-7.txt")

	println("Solution 1 - Sample:", crabss.AlignConstantCosts())
	println("Solution 1 - Input:", crabsi.AlignConstantCosts())

	println("Solution 2 - Sample:", crabss.AlignIncreasingCosts())
	println("Solution 2 - Input:", crabsi.AlignIncreasingCosts())
}
