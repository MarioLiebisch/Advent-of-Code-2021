package bingo

func (bs *Boards) Play(draws []int) (int, int) {
	for di := 0; di < len(draws); di++ {
		bs.Tick(draws[di])
		if winner := bs.FindWinner(); winner != -1 {
			return (*bs)[winner].GetScore(), draws[di]
		}
	}
	// We should never run out of draws, but correctness forces a return
	return 0, -1
}

func Play(file string) int {
	draws, boards := Read(file)
	score, draw := boards.Play(draws)
	return score * draw
}
