package crabs

import (
	"aoc2021/pkg/io"
	"aoc2021/pkg/numbers"
	"sort"
	"strconv"
	"strings"
)

type Crabs []int

var costsmap map[int]int

func init() {
	costsmap = make(map[int]int)
	costsmap[0] = 0
	costsmap[1] = 1
}

// Get movement costs for a distance `diff` based
// on a dynamic lookup table
func getCosts(diff int) int {
	if diff == 0 {
		return 0
	}

	value := costsmap[diff]
	if value == 0 {
		value = diff + getCosts(diff-1)
		costsmap[diff] = value
	}
	return value
}

func LoadCrabs(file string) Crabs {
	var cs Crabs
	for _, v := range strings.Split(io.ReadLines(file)[0], ",") {
		iv, _ := strconv.Atoi(v)
		cs = append(cs, iv)
	}
	sort.Ints(cs)
	return cs
}

func (cs *Crabs) AlignConstantCosts() int {
	costs := 0
	median := numbers.Median(*cs)
	for _, crab := range *cs {
		costs += numbers.Abs(crab - median)
	}
	return costs
}

func (cs *Crabs) AlignIncreasingCosts() int {
	best_costs := 0
	mean := numbers.Mean(*cs)
	// Try the mean value, plus above/below to account for rounding
	for i := -1; i <= 1; i++ {
		costs := 0
		for _, crab := range *cs {
			costs += getCosts(numbers.Abs(crab - mean - i))
		}
		if best_costs == 0 || best_costs > costs {
			best_costs = costs
		}
	}
	return best_costs
}
