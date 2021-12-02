package navigation

import (
	"log"
	"strconv"
	"strings"
)

func ParseCourseSimple(instructions []string) NavState {
	var state NavState
	for _, line := range instructions {
		data := strings.SplitN(line, " ", 2)
		change, err := strconv.Atoi(data[1])
		if err != nil {
			log.Fatal(err)
		}
		switch data[0] {
		case "forward":
			state.Horizontal += change
		case "down":
			state.Depth += change
		case "up":
			state.Depth -= change
		}
	}
	return state
}

func ParseCourse(instructions []string) NavState {
	var state NavState
	for _, line := range instructions {
		data := strings.SplitN(line, " ", 2)
		change, err := strconv.Atoi(data[1])
		if err != nil {
			log.Fatal(err)
		}
		switch data[0] {
		case "down":
			state.Aim += change
		case "up":
			state.Aim -= change
		case "forward":
			state.Horizontal += change
			state.Depth += state.Aim * change
		}
	}
	return state
}
