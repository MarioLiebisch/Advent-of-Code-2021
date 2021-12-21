package dd

import (
	"aoc2021/pkg/numbers"
)

func PlayPracticeGame(pos1, pos2 int) int {
	var d DeterministicDie
	score1, score2 := 0, 0
	for {
		pos1 = ((pos1 + d.Roll() + d.Roll() + d.Roll() - 1) % 10) + 1
		score1 += pos1
		if score1 >= 1000 {
			return score2 * d.Count()
		}

		pos2 = ((pos2 + d.Roll() + d.Roll() + d.Roll() - 1) % 10) + 1
		score2 += pos2
		if score2 >= 1000 {
			return score1 * d.Count()
		}
	}
}

type DiracDiceGame struct {
	wins1, wins2 int
}

func (dd *DiracDiceGame) Play(pos1, pos2, score1, score2, player, unis int) {
	// Player checks happen in reverse, since we check scores before
	// the current player starts rolling
	if score1 >= 21 {
		// If `unis` universes arrive here, count as `unis` wins
		dd.wins1 += unis
		return
	}
	if score2 >= 21 {
		dd.wins2 += unis
		return
	}

	// 3D3 results in 3..9
	// Number of identical universes spawned by sum of 3 rolls
	// rolls                       | sum | count
	// 111                         |  3  | 1
	// 112 121 211                 |  4  | 3
	// 122 212 221 113 131 311     |  5  | 6
	// 222 123 132 213 231 312 321 |  6  | 7
	// 223 232 322 133 313 331     |  7  | 6
	// 233 323 332                 |  8  | 3
	// 333                         |  9  | 1
	rolls := []int{0, 0, 0, 1, 3, 6, 7, 6, 3, 1}
	if player == 0 {
		for roll := 3; roll <= 9; roll++ {
			np := ((pos1 - 1 + roll) % 10) + 1
			dd.Play(np, pos2, score1+np, score2, 1, unis*rolls[roll])
		}
	} else {
		for roll := 3; roll <= 9; roll++ {
			np := ((pos2 - 1 + roll) % 10) + 1
			dd.Play(pos1, np, score1, score2+np, 0, unis*rolls[roll])
		}
	}
}

func PlayDiracDice(pos1, pos2 int) int {
	var dd DiracDiceGame
	dd.Play(pos1, pos2, 0, 0, 0, 1)
	return numbers.Max(dd.wins1, dd.wins2)
}
