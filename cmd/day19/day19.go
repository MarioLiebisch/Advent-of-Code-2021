package main

import "aoc2021/pkg/scanner"

func main() {
	ss := scanner.ReadScanners("./data/sample-19.txt")
	ss.Align()

	si := scanner.ReadScanners("./data/input-19.txt")
	si.Align()

	println("Solution 1 - Sample: ", ss.GetBeaconCount())
	println("Solution 1 - Input: ", si.GetBeaconCount())

	println("Solution 2 - Sample: ", ss.GetScannerSpan())
	println("Solution 2 - Input: ", si.GetScannerSpan())
}
