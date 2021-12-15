package main

import "aoc2021/pkg/chitons"

func main() {
	cms := chitons.ReadChitonsMap("./data/sample-15.txt")
	cmi := chitons.ReadChitonsMap("./data/input-15.txt")

	println("Solution 1 - Sample:", cms.GetPathRisk())
	println("Solution 1 - Input:", cmi.GetPathRisk())
	cms.Enlarge(5)
	cmi.Enlarge(5)
	println("Solution 2 - Sample:", cms.GetPathRisk())
	println("Solution 2 - Input:", cmi.GetPathRisk())
}
