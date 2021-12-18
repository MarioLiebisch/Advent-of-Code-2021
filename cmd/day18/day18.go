package main

import (
	"aoc2021/pkg/snailfish"
)

func main() {
	sample := snailfish.ReadSnailFishNumbers("./data/sample-18.txt")
	for i := 1; i < len(sample); i++ {
		sample[0].Add(sample[i])
	}

	input := snailfish.ReadSnailFishNumbers("./data/input-18.txt")
	for i := 1; i < len(input); i++ {
		input[0].Add(input[i])
	}

	println("Solution 1 - Sample:", sample[0].String(), sample[0].Magnitude())
	println("Solution 1 - Input:", input[0].String(), input[0].Magnitude())

	sample = snailfish.ReadSnailFishNumbers("./data/sample-18.txt")
	input = snailfish.ReadSnailFishNumbers("./data/input-18.txt")

	largest_magnitude := 0
	for i := 0; i < len(sample)-1; i++ {
		for j := i; j < len(sample); j++ {
			tmp1 := sample[i]
			tmp2 := sample[j]
			tmp1.Add(sample[j])
			tmp2.Add(sample[i])
			m1 := tmp1.Magnitude()
			m2 := tmp2.Magnitude()
			if m1 > largest_magnitude {
				largest_magnitude = m1
			}
			if m2 > largest_magnitude {
				largest_magnitude = m2
			}
		}
	}
	println("Solution 2 - Sample:", largest_magnitude)

	largest_magnitude = 0
	for i := 0; i < len(input)-1; i++ {
		for j := i; j < len(input); j++ {
			tmp1 := input[i]
			tmp2 := input[j]
			tmp1.Add(input[j])
			tmp2.Add(input[i])
			m1 := tmp1.Magnitude()
			m2 := tmp2.Magnitude()
			if m1 > largest_magnitude {
				largest_magnitude = m1
			}
			if m2 > largest_magnitude {
				largest_magnitude = m2
			}
		}
	}
	println("Solution 2 - Input:", largest_magnitude)

}
