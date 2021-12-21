package dd

type DeterministicDie struct {
	state int
}

func (d *DeterministicDie) Roll() int {
	d.state++
	// No need to reset state as +-100 will still
	// cause the same fields to be reached
	return d.state
}

func (d *DeterministicDie) Count() int {
	return d.state
}
