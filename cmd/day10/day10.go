package main

import "aoc2021/pkg/chunks"

func main() {
	_, scores, cscores := chunks.ReadChunks("./data/sample-10.txt")
	_, scorei, cscorei := chunks.ReadChunks("./data/input-10.txt")

	println("Solution 1 - Sample:", scores)
	println("Solution 1 - Input:", scorei)

	println("Solution 2 - Sample:", cscores)
	println("Solution 2 - Input:", cscorei)
}
