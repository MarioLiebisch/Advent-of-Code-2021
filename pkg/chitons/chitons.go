package chitons

import (
	"aoc2021/pkg/io"
)

type Chitons []int
type ChitonsMap []Chitons

func ReadChitonsMap(file string) ChitonsMap {
	var cm ChitonsMap
	for _, line := range io.ReadLines(file) {
		var row Chitons
		for _, e := range line {
			row = append(row, int(e-'0'))
		}
		cm = append(cm, row)
	}
	return cm
}

// "Simple" upscaling. References for arrays/members can be pain…
func (cm *ChitonsMap) Enlarge(factor int) {
	rc := len(*cm)
	for i, row := range *cm {
		rl := len(row)
		for j := 1; j < factor; j++ {
			for k := 0; k < rl; k++ {
				nr := row[k] + j
				for nr > 9 {
					nr -= 9
				}
				row = append(row, nr)
			}
		}
		(*cm)[i] = row
	}
	cc := len((*cm)[0])
	for j := 1; j < factor; j++ {
		for i := 0; i < rc; i++ {
			new_row := make(Chitons, cc)
			for k := 0; k < cc; k++ {
				nr := (*cm)[i][k] + j
				for nr > 9 {
					nr -= 9
				}
				new_row[k] = nr
			}
			*cm = append(*cm, new_row)
		}
	}
}

// Couldn't be bothered to implement A* yet again, so let's try mapping it out…
func (cm *ChitonsMap) GetPathRisk() int {
	height := len(*cm)
	width := len((*cm)[0])
	costs := make(map[int]int)

	// Entries to be added to our map
	new_costs := make(map[int]int)
	new_costs[width*height-1] = (*cm)[height-1][width-1]
	// As long as there are new entries
	for len(new_costs) > 0 {
		// Add them to the map
		for pos, nc := range new_costs {
			if costs[pos] == 0 || costs[pos] > nc {
				costs[pos] = nc
			}
		}
		old_costs := new_costs
		new_costs = make(map[int]int)
		// Clear the entries and recalculate all neighbours of previously new entries
		for pos := range old_costs {
			x := pos % width
			y := pos / width
			if x > 0 {
				// This repeats for all 4 directions
				// Get the new costs
				nc := (*cm)[y][x-1] + costs[pos]
				// There are no old costs or the old costs are higher?
				if costs[pos-1] == 0 || costs[pos-1] > nc {
					// There are no new costs yet or the new ones are cheaper?
					if new_costs[pos-1] == 0 || new_costs[pos-1] > nc {
						// Add/Overwrite
						new_costs[pos-1] = nc
					}
				}
			}
			if y > 0 {
				nc := (*cm)[y-1][x] + costs[pos]
				if costs[pos-width] == 0 || costs[pos-width] > nc {
					if new_costs[pos-width] == 0 || new_costs[pos-width] > nc {
						new_costs[pos-width] = nc
					}
				}
			}
			if x < width-1 {
				nc := (*cm)[y][x+1] + costs[pos]
				if costs[pos+1] == 0 || costs[pos+1] > nc {
					if new_costs[pos+1] == 0 || new_costs[pos+1] > nc {
						new_costs[pos+1] = nc
					}
				}
			}
			if y < height-1 {
				nc := (*cm)[y+1][x] + costs[pos]
				if costs[pos+width] == 0 || costs[pos+width] > nc {
					if new_costs[pos+width] == 0 || new_costs[pos+width] > nc {
						new_costs[pos+width] = nc
					}
				}
			}
		}
	}
	// Since we have calculated the least cost/risk for our starting point
	// return it, but remove it's own step costs (which are still included)
	return costs[0] - (*cm)[0][0]
}
