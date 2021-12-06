package main

import "aoc2021/pkg/fish"

func main() {
	fls := fish.LoadFish("./data/sample-6.txt")
	fli := fish.LoadFish("./data/input-6.txt")

	println("Solution 1 - Sample:", fls.Process(80).Count())
	println("Solution 1 - Input:", fli.Process(80).Count())

	println("Solution 2 - Sample:", fls.Process(256-80).Count())
	println("Solution 2 - Input:", fli.Process(256-80).Count())
}
