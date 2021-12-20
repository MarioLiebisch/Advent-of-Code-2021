package main

import "aoc2021/pkg/trench"

func main() {
	sds := trench.ReadScannerData("./data/sample-20.txt")
	sds.Reconstruct(2)
	println("Solution 1 - Sample:", sds.Count())
	sds.Reconstruct(48)
	println("Solution 2 - Sample:", sds.Count())

	sdi := trench.ReadScannerData("./data/input-20.txt")
	sdi.Reconstruct(2)
	println("Solution 1 - Input:", sdi.Count())
	sdi.Reconstruct(48)
	println("Solution 2 - Input:", sdi.Count())
}
