package polymers

import (
	"aoc2021/pkg/io"
	"strings"
)

type PolymerInsertions map[string]string

type PolymerInstructions struct {
	Template Polymer
	Rules    PolymerInsertions
}

type Polymer struct {
	chains map[string]int64
	counts map[byte]int64
}

func ReadPolymerInstructions(file string) PolymerInstructions {
	var res PolymerInstructions
	res.Rules = make(PolymerInsertions)
	first := true
	for _, line := range io.ReadLines(file) {
		if first {
			first = false
			res.Template.chains = make(map[string]int64)
			res.Template.counts = make(map[byte]int64)
			for i := 0; i < len(line); i++ {
				res.Template.counts[line[i]]++
			}
			for i := 0; i < len(line)-1; i++ {
				res.Template.chains[line[i:i+2]]++
			}
			continue
		}
		if line == "" {
			continue
		}
		data := strings.SplitN(line, " -> ", 2)
		res.Rules[data[0]] = data[1]
	}
	return res
}

func (p *Polymer) Apply(pi *PolymerInstructions) {
	// Create a new polymer and copy the existing counts
	var np Polymer
	np.chains = make(map[string]int64)
	np.counts = make(map[byte]int64)
	for k, v := range p.counts {
		np.counts[k] = v
	}
	// Now process all chains and add the new chains to the new polymer
	for seg, count := range p.chains {
		add := pi.Rules[seg]
		new1 := string(seg[0]) + add
		new2 := add + string(seg[1])

		np.chains[new1] += count
		np.chains[new2] += count
		np.counts[add[0]] += count
	}
	// Move everything back
	*p = np
}

func (p *Polymer) Solve(pi *PolymerInstructions, applications int) int64 {
	for i := 0; i < applications; i++ {
		p.Apply(pi)
	}

	least_common := p.counts['N']
	most_common := p.counts['N']
	for _, c := range p.counts {
		if c < least_common {
			least_common = c
		}
		if c > most_common {
			most_common = c
		}
	}
	return most_common - least_common
}
