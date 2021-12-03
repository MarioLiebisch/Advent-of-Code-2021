package numbers

import (
	"log"
	"strconv"
)

func MostCommonBits(data []string) string {
	length := len(data)
	llength := len(data[0])
	var ones []int
	for i := 0; i < llength; i++ {
		ones = append(ones, 0)
	}
	for _, entry := range data {
		if entry == "" {
			continue
		}
		for i := 0; i < llength; i++ {
			if entry[i] == '1' {
				ones[i]++
			}
		}
	}
	ret := ""
	for i := 0; i < llength; i++ {
		if ones[i] > length/2 {
			ret += "1"
		} else {
			ret += "0"
		}
	}
	return ret
}

func ReduceByCommonBits(data []string, most bool) string {
	length := len(data[0])
	for i := 0; i < length && len(data) > 1; i++ {
		ones := 0
		for j := 0; j < len(data); j++ {
			if data[j][i] == '1' {
				ones++
			}
		}

		var keep byte = '0'
		if ones >= len(data)-ones {
			if most {
				keep = '1'
			}
		} else if !most {
			keep = '1'
		}

		var ndata []string
		for _, entry := range data {
			if entry[i] == keep {
				ndata = append(ndata, entry)
			}
		}
		data = ndata
	}
	if len(data) == 1 {
		return data[0]
	} else {
		log.Fatal("Shouldn't happen!")
		return ""
	}
}

func BinNot(data string) string {
	ret := ""
	for _, c := range data {
		if c == '1' {
			ret += "0"
		} else {
			ret += "1"
		}
	}
	return ret
}

func BinToInt(data string) int {
	if number, err := strconv.ParseInt(data, 2, 64); err != nil {
		log.Fatal(err)
		return 0
	} else {
		return int(number)
	}
}
