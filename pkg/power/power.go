package power

import "aoc2021/pkg/numbers"

func GetEpsilon(diagnostics []string) int {
	return numbers.BinToInt(numbers.BinNot(numbers.MostCommonBits(diagnostics)))
}

func GetGamma(diagnostics []string) int {
	return numbers.BinToInt(numbers.MostCommonBits(diagnostics))
}

func GetPowerConsumption(diagnostics []string) int {
	mcb := numbers.MostCommonBits(diagnostics)
	epsilon := numbers.BinToInt(numbers.BinNot(mcb))
	gamma := numbers.BinToInt(mcb)
	return epsilon * gamma
}
