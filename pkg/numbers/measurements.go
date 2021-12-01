package numbers

func CountIncrements(input []int) int {
	num := 0
	prev := 0
	for i, v := range input {
		if i > 0 && v > prev {
			num++
		}
		prev = v
	}
	return num
}

func MergeWindow(input []int, width int) []int {
	var out []int
	for i := 0; i <= len(input)-width; i++ {
		tmp := 0
		for j := 0; j < width; j++ {
			tmp += input[i+j]
		}
		out = append(out, tmp)
	}
	return out
}
