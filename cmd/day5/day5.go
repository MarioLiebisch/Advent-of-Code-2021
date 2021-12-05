package main

import (
	"aoc2021/pkg/vents"
)

func main() {
	ventss := vents.ReadVents("./data/sample-5.txt")
	ventsi := vents.ReadVents("./data/input-5.txt")

	ventmaps := make(vents.VentMap)
	ventmaps.AddVents(ventss, false)
	println("Sample map 1:")
	ventmaps.Print()
	println("Solution 1 - Sample:", ventmaps.CountCrossings())

	println()

	ventmapi := make(vents.VentMap)
	ventmapi.AddVents(ventsi, false)
	println("Solution 1 - Input:", ventmapi.CountCrossings())

	println()

	ventmaps = make(vents.VentMap)
	ventmaps.AddVents(ventss, true)
	println("Sample map 2:")
	ventmaps.Print()
	println("Solution 2 - Sample:", ventmaps.CountCrossings())

	println()

	ventmapi = make(vents.VentMap)
	ventmapi.AddVents(ventsi, true)
	println("Solution 2 - Input:", ventmapi.CountCrossings())
}
