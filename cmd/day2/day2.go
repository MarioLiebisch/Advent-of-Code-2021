package main

import (
	"aoc2021/pkg/io"
	"aoc2021/pkg/navigation"
)

func main() {
	println("Solution 1 - Sample:", navigation.ParseCourseSimple(io.ReadLines("data/sample-2.txt")).Checksum())
	println("Solution 1 - Input:", navigation.ParseCourseSimple(io.ReadLines("data/input-2.txt")).Checksum())

	println("Solution 2 - Sample:", navigation.ParseCourse(io.ReadLines("data/sample-2.txt")).Checksum())
	println("Solution 2 - Input:", navigation.ParseCourse(io.ReadLines("data/input-2.txt")).Checksum())
}
