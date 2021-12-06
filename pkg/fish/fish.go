package fish

import (
	"aoc2021/pkg/io"
	"strconv"
	"strings"
)

// Instead of storing individual fishes,
// it's far easier (and way faster!) to
// just store counts by "progress".
type FishList [9]int

func LoadFish(file string) FishList {
	var fl FishList
	for _, v := range strings.Split(io.ReadLines(file)[0], ",") {
		iv, _ := strconv.Atoi(v)
		fl[iv]++
	}
	return fl
}

func (fl *FishList) Cycle() *FishList {
	// Store number of fishes ready to spawn
	spawns := fl[0]
	// Move up all other fishes' progress by one
	for i := 0; i < 8; i++ {
		fl[i] = fl[i+1]
	}
	// "Reset" those spawning to progress "6"
	fl[6] += spawns
	// Add new fish as progress "8"
	fl[8] = spawns
	return fl
}

func (fl *FishList) Process(days int) *FishList {
	for i := 0; i < days; i++ {
		fl.Cycle()
	}
	return fl
}

func (fl *FishList) Count() int {
	count := 0
	// Just add up the individual counts
	for _, c := range fl {
		count += c
	}
	return count
}
