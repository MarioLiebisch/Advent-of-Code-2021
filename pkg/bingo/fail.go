package bingo

func (bs *Boards) Fail(draws []int) (int, int) {
	nb := *bs
	for di := 0; di < len(draws); di++ {
		nb.Tick(draws[di])
		// There might be more than one winner at some point so loop as long as there's one
		for {
			if winner := nb.FindWinner(); winner != -1 {
				if len(nb) == 1 {
					return nb[0].GetScore(), draws[di]
				} else {
					// Remove the winner by swapping it with last in the set
					// then reducing size by 1
					last := len(nb) - 1
					nb[winner] = nb[last]
					nb = nb[:last]
				}
				// Look for more
				continue
			}
			break
		}
	}
	// We should never run out of draws, but correctness forces a return
	return 0, -1
}

func Fail(file string) int {
	draws, boards := Read(file)
	score, draw := boards.Fail(draws)
	return score * draw
}
