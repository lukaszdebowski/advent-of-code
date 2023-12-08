package main

import (
	"fmt"
	"os"
	"strings"
)

type Instruction struct {
	left  string
	right string
}

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")

	directions := strings.Split(lines[0], "")
	allInstructions := map[string]Instruction{}

	for _, line := range lines[2:] {
		allInstructions[line[:3]] = Instruction{
			left:  line[7:10],
			right: line[12:15],
		}
	}

	var instructions = []string{}
	for k := range allInstructions {
		if strings.HasSuffix(k, "A") {
			instructions = append(instructions, k)
		}
	}

	steps := []int{}
	for i := 0; i < len(instructions); i++ {
		key := instructions[i]
		step := 0
	inner:
		for {
			instruction := allInstructions[key]
			dir := directions[step%len(directions)]
			switch dir {
			case "L":
				key = instruction.left
			case "R":
				key = instruction.right
			}
			step++
			if strings.HasSuffix(key, "Z") {
				steps = append(steps, step)
				break inner
			}
		}
	}

	result := LCM(steps[0], steps[1], steps[2:]...)
	fmt.Println("result:", result)
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)
	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}
	return result
}
