package main

import "aoc2021/pkg/polymers"

func main() {
	inss := polymers.ReadPolymerInstructions("./data/sample-14.txt")
	insi := polymers.ReadPolymerInstructions("./data/input-14.txt")

	ps := inss.Template
	pi := insi.Template

	println("Solution 1 - Sample:", ps.Solve(&inss, 10))
	println("Solution 1 - Input:", pi.Solve(&insi, 10))

	ps = inss.Template
	pi = insi.Template

	println("Solution 2 - Sample:", ps.Solve(&inss, 40))
	println("Solution 2 - Input:", pi.Solve(&insi, 40))
}
