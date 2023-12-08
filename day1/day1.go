package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")

	result := part2Solution(lines)
	fmt.Println("Result:", result)
}

func part1Solution(lines []string) int {
	result := 0

	for _, line := range lines {
		firstDigit := part1GetFirstDigit(line)
		lastDigit := part1GetLastDigit(line)

		value := 10*firstDigit + lastDigit
		result += value
	}

	return result
}

func part1GetFirstDigit(line string) int {
	for _, char := range line {
		if unicode.IsNumber(char) {
			number, _ := strconv.Atoi(string(char))
			return number
		}
	}
	panic("No number found in line")
}

func part1GetLastDigit(line string) int {
	for i := len(line) - 1; i >= 0; i-- {
		if unicode.IsNumber(rune(line[i])) {
			number, _ := strconv.Atoi(string(line[i]))
			return number
		}
	}
	panic("No number found in line")
}

func part2Solution(lines []string) int {
	result := 0

	for _, line := range lines {
		firstDigit := part2GetFirstDigit(line)
		lastDigit := part2GetLastDigit(line)

		value := 10*firstDigit + lastDigit
		result += value
	}

	return result
}

func part2GetFirstDigit(line string) int {
	for i := 0; i < len(line); i++ {
		substring := line[i:]

		if unicode.IsNumber(rune(substring[0])) {
			number, _ := strconv.Atoi(string(substring[0]))
			return number
		}

		for k, v := range wordToDigit {
			if strings.HasPrefix(substring, k) {
				return v
			}
		}

	}
	panic("No number found in line")
}

func part2GetLastDigit(line string) int {
	for i := len(line) - 1; i >= 0; i-- {
		substring := line[:i+1]

		if unicode.IsNumber(rune(substring[i])) {
			number, _ := strconv.Atoi(string(line[i]))
			return number
		}

		for k, v := range wordToDigit {
			if strings.HasSuffix(substring, k) {
				return v
			}
		}
	}
	panic("No number found in line")
}

var wordToDigit = map[string]int{
	"one":   1,
	"two":   2,
	"four":  4,
	"three": 3,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}
