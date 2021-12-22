package main

import "aoc2021/pkg/reactor"

func main() {
	var rs, ri reactor.Reactor
	is, ii := reactor.LoadInstructions("./data/sample-22.txt"), reactor.LoadInstructions("./data/input-22.txt")
	rs.Apply(is, true)
	println("Solution 1 - Sample:", rs.GetCount())
	ri.Apply(ii, true)
	println("Solution 1 - Input:", ri.GetCount())

	// Clear results
	rs = rs[:0]
	ri = ri[:0]

	rs.Apply(is, false)
	println("Solution 2 - Sample:", rs.GetCount())
	ri.Apply(ii, false)
	println("Solution 2 - Input:", ri.GetCount())
}
