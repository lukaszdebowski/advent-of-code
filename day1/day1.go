package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	data, _ := os.ReadFile("example.txt")
	lines := strings.Split(string(data), "\n")

	result := 0

	for _, line := range lines {
		firstDigit := getFirstDigit(line)
		lastDigit := getLastDigit(line)

		value := 10*firstDigit + lastDigit
		result += value
	}

	fmt.Println("Result:", result)
}

func getFirstDigit(line string) int {
	for _, char := range line {
		if unicode.IsNumber(char) {
			number, _ := strconv.Atoi(string(char))
			fmt.Println("Found first number:", number)
			return number
		}
	}
	panic("No number found in line")
}

func getLastDigit(line string) int {
	for i := len(line) - 1; i >= 0; i-- {
		if unicode.IsNumber(rune(line[i])) {
			number, _ := strconv.Atoi(string(line[i]))
			fmt.Println("Found last number:", number)
			return number
		}
	}
	panic("No number found in line")

}
