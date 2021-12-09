package main

import "aoc2021/pkg/height"

func main() {
	hms := height.Load("./data/sample-9.txt")
	hmi := height.Load("./data/input-9.txt")

	_, risks := hms.LowPoints()
	_, riski := hmi.LowPoints()

	println("Solution 1 - Sample:", risks)
	println("Solution 1 - Input:", riski)

	bs := hms.FindBasins()
	bi := hmi.FindBasins()

	println("Solution 2 - Sample:", bs[0]*bs[1]*bs[2])
	println("Solution 2 - Input:", bi[0]*bi[1]*bi[2])
}
