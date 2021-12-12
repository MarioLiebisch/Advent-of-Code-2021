package caves

import (
	"aoc2021/pkg/io"
	"strings"
	"unicode"
)

type Caves map[string][]string
type Path []string
type Paths []Path

func (path *Path) Contains(node string) bool {
	for _, step := range *path {
		if step == node {
			return true
		}
	}
	return false
}

func (path *Path) IsAllowed1(next string) bool {
	// Uppercase is always allowed
	if unicode.IsUpper(rune(next[0])) {
		return true
	}
	// Otherwise see if it's included already
	return !path.Contains(next)
}

func (path *Path) IsAllowed2(next string) bool {
	// Uppercase is always allowed
	if unicode.IsUpper(rune(next[0])) {
		return true
	}

	// Count all lowercase nodes so far
	counts := make(map[string]int)
	for _, step := range *path {
		if unicode.IsLower(rune(step[0])) {
			counts[step]++
		}
	}

	// Not in the path so far
	if counts[next] == 0 {
		return true
	}
	// Already twice in the path
	if counts[next] == 2 {
		return false
	}

	// Look for any other duplicate
	other_twice := false
	for k, v := range counts {
		if k != next && v > 1 {
			other_twice = true
			break
		}
	}
	return !other_twice
}

func ReadCaves(file string) Caves {
	caves := make(Caves)
	for _, line := range io.ReadLines(file) {
		parts := strings.SplitN(line, "-", 2)
		caves[parts[0]] = append(caves[parts[0]], parts[1])
		caves[parts[1]] = append(caves[parts[1]], parts[0])
	}
	return caves
}

func (caves *Caves) GetPathsStep1(from, to string, current Path) Paths {
	var paths Paths
	now := current[len(current)-1]
	for _, next := range (*caves)[now] {
		if next == to {
			// We're at the goal
			paths = append(paths, append(current, next))
		} else if next == from {
			// Don't  go back to the start
		} else if current.IsAllowed1(next) {
			// Take the next step and append all found paths
			paths = append(paths, caves.GetPathsStep1(from, to, append(current, next))...)
		}
	}
	return paths
}

func (caves *Caves) GetPaths1(from, to string) Paths {
	return caves.GetPathsStep1(from, to, Path{from})
}

func (caves *Caves) GetPathsStep2(from, to string, current Path) Paths {
	var paths Paths
	now := current[len(current)-1]
	for _, next := range (*caves)[now] {
		if next == to {
			// We're at the goal
			paths = append(paths, append(current, next))
		} else if next == from {
			// Don't  go back to the start
		} else if current.IsAllowed2(next) {
			// Take the next step and append all found paths
			paths = append(paths, caves.GetPathsStep2(from, to, append(current, next))...)
		}
	}
	return paths
}

func (caves *Caves) GetPaths2(from, to string) Paths {
	return caves.GetPathsStep2(from, to, Path{from})
}
