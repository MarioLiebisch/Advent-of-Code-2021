package chunks

import (
	"aoc2021/pkg/io"
	"errors"
	"sort"
)

type Chunks []Chunk

type Chunk struct {
	tag      byte
	children Chunks
}

var ends = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var scores = map[byte]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var comp_scores = map[byte]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

// Small "wrapper" workaround to read multiple chunks in one line that aren't nested in one "root"
func ReadChunkLine(line string) (Chunks, int, int, int, error) {
	var chunks Chunks
	var err error
	offset := 0
	for offset <= len(line) && err == nil {
		chunk, consumed, score, cscore, cerr := ReadChunk(line[offset:])
		chunks = append(chunks, chunk)
		offset += consumed
		if cerr != nil || cscore > 0 {
			return chunks, offset, score, cscore, cerr
		}
	}
	return chunks, offset, 0, 0, nil
}

func ReadChunk(input string) (Chunk, int, int, int, error) {
	var chunk Chunk
	consumed := 0
	token := input[0]
	expected := ends[token]
	if expected == 0 {
		return chunk, consumed, scores[token], 0, errors.New("Unexpected token: " + string(token))
	}
	chunk.tag = token
	// Input only contains the opening tag - add completion score
	if len(input) == 1 {
		return chunk, consumed, 0, comp_scores[expected], nil
	}

	// Consume opening tag
	consumed++
	for next := input[consumed]; next != expected; next = input[consumed] {
		child, cconsumed, score, cscore, err := ReadChunk(input[consumed:])
		// Child didn't parse properly (error or incomplete)
		if err != nil || cscore > 0 {
			return chunk, consumed + cconsumed, score, cscore*5 + comp_scores[expected], err
		}
		chunk.children = append(chunk.children, child)
		consumed += cconsumed
		// Reached end of line?
		if consumed >= len(input) {
			return chunk, consumed - 1, 0, cscore*5 + comp_scores[expected], nil
		}
	}
	// Consume closing tag
	consumed++
	return chunk, consumed, 0, 0, nil
}

func ReadChunks(file string) (Chunks, int, int) {
	var chunks Chunks
	total_score := 0
	var cscores []int
	for _, line := range io.ReadLines(file) {
		if nchunks, _, score, cscore, err := ReadChunkLine(line); err != nil {
			total_score += score
			continue
		} else {
			chunks = append(chunks, nchunks...)
			if cscore > 0 {
				cscores = append(cscores, cscore)
			}
		}
	}
	comp_score := 0
	if len(cscores) > 0 {
		sort.Ints(cscores)
		comp_score = cscores[len(cscores)/2]
	}
	return chunks, total_score, comp_score
}
