package air

import "aoc2021/pkg/numbers"

func GetOxygenGeneratorRating(diagnostics []string) int {
	return numbers.BinToInt(numbers.ReduceByCommonBits(diagnostics, true))
}

func GetCO2ScrubberRating(diagnostics []string) int {
	return numbers.BinToInt(numbers.ReduceByCommonBits(diagnostics, false))
}

func GetLifeSupportRating(diagnostics []string) int {
	return GetOxygenGeneratorRating(diagnostics) * GetCO2ScrubberRating(diagnostics)
}
