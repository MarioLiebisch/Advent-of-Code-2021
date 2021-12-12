package main

import "aoc2021/pkg/caves"

func main() {
	samplea := caves.ReadCaves("./data/sample-12a.txt")
	sampleb := caves.ReadCaves("./data/sample-12b.txt")
	samplec := caves.ReadCaves("./data/sample-12c.txt")
	input := caves.ReadCaves("./data/input-12.txt")

	println("Solution 1 - Sample a:", len(samplea.GetPaths1("start", "end")))
	println("Solution 1 - Sample b:", len(sampleb.GetPaths1("start", "end")))
	println("Solution 1 - Sample c:", len(samplec.GetPaths1("start", "end")))
	println("Solution 1 - Input:", len(input.GetPaths1("start", "end")))

	println("Solution 2 - Sample a:", len(samplea.GetPaths2("start", "end")))
	println("Solution 2 - Sample b:", len(sampleb.GetPaths2("start", "end")))
	println("Solution 2 - Sample c:", len(samplec.GetPaths2("start", "end")))
	println("Solution 2 - Input:", len(input.GetPaths2("start", "end")))
}
