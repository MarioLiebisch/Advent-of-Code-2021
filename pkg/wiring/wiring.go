package wiring

import (
	"aoc2021/pkg/io"
	"log"
	"sort"
	"strings"
)

type WiringSet struct {
	digits  Glyphs
	outputs Glyphs
	idcache map[int]string
	mapping map[rune]rune
}

type Glyphs []string

type WiringSets []WiringSet

// Mapping of how the output wires *should* map to digits
var glyphs = map[string]int{
	"abcefg":  0,
	"cf":      1,
	"acdeg":   2,
	"acdfg":   3,
	"bcdf":    4,
	"abdfg":   5,
	"abdefg":  6,
	"acf":     7,
	"abcdefg": 8,
	"abcdfg":  9,
}

func (sets *WiringSets) CountPartOne() int {
	count := 0
	for _, set := range *sets {
		for _, d := range set.Decode() {
			if d == 1 || d == 4 || d == 7 || d == 8 {
				count++
			}
		}
	}
	return count
}

func (sets *WiringSets) SumPartTwo() int {
	sum := 0
	for _, set := range *sets {
		number := 0
		for _, d := range set.Decode() {
			number = number*10 + d
		}
		sum += number
	}
	return sum
}

// Prebuild a full mapping, if it isn't done already
// Wouldn't be necessary, but was nice for testing
func (set *WiringSet) BuildMapping() {
	if set.mapping == nil || len(set.mapping) == 0 {
		set.mapping = make(map[rune]rune)
		set.idcache = make(map[int]string)
		for _, c := range "abcdefg" {
			set.IdentifySegment(c)
		}
	}
}

func (set *WiringSet) Decode() []int {
	var res []int
	set.BuildMapping()
	for _, encoded := range set.outputs {
		var decoded string
		for _, e := range encoded {
			found := false
			for d, m := range set.mapping {
				if m == e {
					decoded += string(d)
					found = true
					break
				}
			}
			if !found {
				log.Fatalln("Didn't find ", e)
			}
		}
		decoded = stringSort(decoded)
		number := glyphs[decoded]
		res = append(res, number)
	}
	return res
}

// Reduce a given string by glyphs in another string
func removeGlyphs(set string, dirty string) string {
	var res string
	for _, c := range set {
		if !strings.ContainsRune(dirty, c) {
			res += string(c)
		}
	}
	return res
}

// Try to use known segments and digits to identify this set's mapping for a given segment
// Building this one up step by step *before* completely reading the first task was kind of
// not the best idea, but that made part 2 easy…
func (set *WiringSet) IdentifySegment(seg rune) rune {
	cache := set.mapping[seg]
	if cache != 0 {
		return cache
	}
	switch seg {
	case 'a':
		one := set.IdentifyDigit(1)
		seven := set.IdentifyDigit(7)
		unique := removeGlyphs(seven, one)
		r := rune(unique[0])
		set.mapping['a'] = r
		return r
	case 'b':
		eight := set.IdentifyDigit(8)
		two := set.IdentifyDigit(2)
		seg_f := set.IdentifySegment('f')
		unique := removeGlyphs(removeGlyphs(eight, two), string(seg_f))
		r := rune(unique[0])
		set.mapping['b'] = r
		return r
	case 'c':
		one := set.IdentifyDigit(1)
		six := set.IdentifyDigit(6)
		unique := removeGlyphs(one, six)
		r := rune(unique[0])
		set.mapping['c'] = r
		return r
	case 'd':
		two := set.IdentifyDigit(2)
		four := set.IdentifyDigit(4)
		seg_c := set.IdentifySegment('c')
		for _, c := range two {
			if strings.ContainsRune(four, c) && c != seg_c {
				set.mapping['d'] = c
				return c
			}
		}
	case 'e':
		eight := set.IdentifyDigit(8)
		nine := set.IdentifyDigit(9)
		unique := removeGlyphs(eight, nine)
		r := rune(unique[0])
		set.mapping['e'] = r
		return r
	case 'f':
		one := set.IdentifyDigit(1)
		six := set.IdentifyDigit(6)
		unique := removeGlyphs(one, six)
		var r rune
		if unique[0] == one[0] {
			r = rune(one[1])
		} else {
			r = rune(one[0])
		}
		set.mapping['f'] = r
		return r
	case 'g':
		// 9 - 4 - a
		nine := set.IdentifyDigit(9)
		four := set.IdentifyDigit(4)
		seg_a := set.IdentifySegment('a')
		unique := removeGlyphs(removeGlyphs(nine, four), string(seg_a))
		r := rune(unique[0])
		set.mapping['g'] = r
		return r
	}
	log.Fatal("Shouldn't ever happen!")
	return 0
}

// Try to use known segments and digits to identify this set's mapping for a given segment
// Since some of the more ambiguous glyphs are never identified this way (nor needed), they
// can just be omitted.
func (set *WiringSet) IdentifyDigit(digit int) string {
	cache := set.idcache[digit]
	if cache != "" {
		return cache
	}
	switch digit {
	case 0:
		eight := set.IdentifyDigit(8)
		seg_d := set.IdentifySegment('d')
		r := removeGlyphs(eight, string(seg_d))
		set.idcache[0] = r
		return r
	case 1:
		for _, dgt := range set.digits {
			if len(dgt) == 2 {
				set.idcache[1] = dgt
				return dgt
			}
		}
	case 2:
		seg_f := set.IdentifySegment('f')
		for _, dgt := range set.digits {
			if len(dgt) == 5 {
				if !strings.ContainsRune(dgt, seg_f) {
					set.idcache[2] = dgt
					return dgt
				}
			}
		}
	case 3: // Not needed to decode
	case 4:
		for _, dgt := range set.digits {
			if len(dgt) == 4 {
				set.idcache[4] = dgt
				return dgt
			}
		}
	case 5: // Not needed to decode
	case 6:
		one := set.IdentifyDigit(1)
		eight := set.IdentifyDigit(8)
		v1 := removeGlyphs(eight, string(one[0]))
		v2 := removeGlyphs(eight, string(one[1]))
		for _, dgt := range set.digits {
			if dgt == v1 || dgt == v2 {
				set.idcache[6] = dgt
				return dgt
			}
		}
	case 7:
		for _, dgt := range set.digits {
			if len(dgt) == 3 {
				set.idcache[7] = dgt
				return dgt
			}
		}
	case 8:
		for _, dgt := range set.digits {
			if len(dgt) == 7 {
				set.idcache[8] = dgt
				return dgt
			}
		}
	case 9:
		zero := set.IdentifyDigit(0)
		six := set.IdentifyDigit(6)
		for _, dgt := range set.digits {
			if len(dgt) == 6 && dgt != zero && dgt != six {
				set.idcache[9] = dgt
				return dgt
			}
		}
	}
	log.Fatal("Something went wrong!")
	return ""
}

// Sort the glyphs in a string by alphabet
func stringSort(str string) string {
	s := []rune(str)
	sort.Slice(s, func(a, b int) bool { return s[a] < s[b] })
	return string(s)
}

func ReadWiringSets(file string) WiringSets {
	var sets WiringSets
	for _, line := range io.ReadLines(file) {
		var ws WiringSet
		parts := strings.SplitN(line, "|", 2)
		for _, digit := range strings.Split(parts[0], " ") {
			if digit != "" {
				digit = stringSort(digit)
				ws.digits = append(ws.digits, digit)
			}
		}
		for _, output := range strings.Split(parts[1], " ") {
			if output != "" {
				ws.outputs = append(ws.outputs, output)
			}
		}
		sets = append(sets, ws)
	}
	return sets
}
