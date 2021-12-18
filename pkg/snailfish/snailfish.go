package snailfish

import (
	"aoc2021/pkg/io"
	"fmt"
)

type SFNumber []int
type SFNumbers []SFNumber

func ReadSnailFishNumbers(file string) SFNumbers {
	var sfns SFNumbers
	for _, line := range io.ReadLines(file) {
		var sfn SFNumber
		for _, c := range line {
			switch c {
			case '[':
				sfn = append(sfn, -1)
			case ']':
				sfn = append(sfn, -2)
			case ',':
				sfn = append(sfn, -3)
			default:
				sfn = append(sfn, int(c-'0'))
			}
		}
		sfns = append(sfns, sfn)
	}
	return sfns
}

func (sfn *SFNumber) Add(other SFNumber) {
	*sfn = append(append(append(append(SFNumber{-1}, (*sfn)...), -3), other...), -2)
	sfn.Reduce()
}

func (sfn *SFNumber) magnitude(offset int) (int, int) {
	v := (*sfn)[offset]
	switch v {
	case -1: // [
		ml, ll := sfn.magnitude(offset + 1)
		mr, lr := sfn.magnitude(offset + ll + 2)
		return ml*3 + mr*2, ll + lr + 3
	case -2: // ]
	case -3: // ,
	default:
		return v, 1
	}
	return 0, 0
}

func (sfn *SFNumber) Magnitude() int {
	m, _ := sfn.magnitude(0)
	return m
}

func (sfn *SFNumber) String() string {
	str := ""
	for _, c := range *sfn {
		switch c {
		case -1:
			str += "["
		case -2:
			str += "]"
		case -3:
			str += ","
		default:
			str += fmt.Sprint(c)
		}
	}
	return str
}

func (sfn *SFNumber) Reduce() {
	for {
		changed := false
		depth := 0
		for i := 0; i < len(*sfn) && !changed; i++ {
			c := (*sfn)[i]
			switch c {
			case -1: // [
				depth++
			case -2: // ]
				depth--
			case -3: // ,
			default:
				if depth >= 5 {
					left := c
					right := (*sfn)[i+2]
					*sfn = append(append((*sfn)[:i-1], 0), (*sfn)[i+4:]...)
					for j := i - 2; j > 0; j-- {
						if (*sfn)[j] >= 0 {
							(*sfn)[j] += left
							break
						}
					}
					for j := i; j < len(*sfn); j++ {
						if (*sfn)[j] >= 0 {
							(*sfn)[j] += right
							break
						}
					}
					changed = true
				}
			}
		}
		if changed {
			continue
		}

		for i := 0; i < len(*sfn) && !changed; i++ {
			c := (*sfn)[i]
			switch c {
			case -1: // [
			case -2: // ]
			case -3: // ,
			default:
				if c > 9 {
					var nsfn SFNumber
					nsfn = append(nsfn, (*sfn)[:i]...)
					nsfn = append(nsfn, -1, c/2, -3, (c+1)/2, -2)
					nsfn = append(nsfn, (*sfn)[i+1:]...)
					*sfn = nsfn
					changed = true
				}
			}
		}
		if !changed {
			break
		}
	}
}
