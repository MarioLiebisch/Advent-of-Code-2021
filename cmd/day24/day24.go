package main

import (
	"aoc2021/pkg/alu"
	"fmt"
	"sync"
)

type storedState struct {
	state             alu.State
	largest, smallest uint64
	count             uint64
}

// I originally used a vector of X/Y/Z as key, but considering
// only Z matters, this can be simplified to just using that
type storeKey = int64

func findSerialNumbers() (uint64, uint64, uint64) {
	program := alu.CompileProgram("./data/input-24.txt", true)

	// Initialize a fake "old state" as the initial state
	old_states := make(map[storeKey]storedState)
	old_states[0] = storedState{count: 1}

	// We've got 14 digits to process
	for n := 1; n <= 14; n++ {
		println("Evaluating", len(old_states)*9, "possible results after position", n, "using", len(old_states), "different starting states")
		new_states := make(map[storeKey]storedState)
		for _, state := range old_states {
			// Use at least some parallelization
			var wg sync.WaitGroup
			var lock sync.Mutex
			wg.Add(9)
			for i := 1; i <= 9; i++ {
				smallest_input := state.smallest*10 + uint64(i)
				largest_input := state.largest*10 + uint64(i)
				go func() {
					defer wg.Done()
					var processor alu.ALU

					// We only have to use either the largest or smallest input,
					// since both numbers arrive here with the same state
					// This will exit once the input buffer runs out
					new_state := processor.Run(program, fmt.Sprint(largest_input), &state.state)

					// Lock the list of new states and add the result
					lock.Lock()
					entry, exists := new_states[new_state.Z]
					// If we've reached this state before, make sure to
					// save the smallest/largest number
					if exists {
						if entry.smallest > smallest_input {
							entry.smallest = smallest_input
						}
						if entry.largest < largest_input {
							entry.largest = largest_input
						}
						entry.count += state.count
						new_states[new_state.Z] = entry
					} else {
						new_states[new_state.Z] = storedState{state: new_state, smallest: smallest_input, largest: largest_input, count: state.count}
					}
					// Let other goroutines access the results
					lock.Unlock()
				}()
			}
			wg.Wait()
		}
		// Move new states to old states and repeat
		old_states = new_states
		println("- Resulting states:", len(old_states))
	}

	// We now have all possible states produced by any given number
	// together with their minimum/maximum input value.
	// All that's left to do is read the values for "Z = 0".
	return old_states[0].smallest, old_states[0].largest, old_states[0].count
}

func main() {
	var alus alu.ALU
	tmp := alus.Run(alu.CompileProgram("./data/sample-24.txt", false), "5", nil)
	println("Test: ", tmp.W, tmp.X, tmp.Y, tmp.Z) // should return binary 0101 in W, X, Y, Z
	println()

	smallest, largest, count := findSerialNumbers()
	println()
	println("Solution 1 - Largest valid serial number:", largest)
	println("Solution 2 - Smallest valid serial number:", smallest)
	println("Bonus - Total valid serial numbers:", count)
}
