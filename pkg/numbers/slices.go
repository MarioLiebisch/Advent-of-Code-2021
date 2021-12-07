package numbers

func Median(values []int) int {
	length := len(values)
	if length == 0 {
		return 0
	}
	return values[length/2]
}

func Mean(values []int) int {
	length := len(values)
	if length == 0 {
		return 0
	}
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum / length
}
