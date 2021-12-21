package main

import "aoc2021/pkg/dd"

func main() {
	println("Solution 1 - Sample:", dd.PlayPracticeGame(4, 8))
	println("Solution 1 - Input:", dd.PlayPracticeGame(8, 7))

	println("Solution 2 - Sample:", dd.PlayDiracDice(4, 8))
	println("Solution 2 - Input:", dd.PlayDiracDice(8, 7))
}
