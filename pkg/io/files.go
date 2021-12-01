package io

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadLines(file string) []string {
	var lines []string

	fp, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	return lines
}

func ReadIntegers(file string) []int {
	var numbers []int
	for _, line := range ReadLines(file) {
		number, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalln(err)
		}
		numbers = append(numbers, number)
	}
	return numbers
}
