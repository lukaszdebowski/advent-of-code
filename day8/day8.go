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
	instructions := map[string]Instruction{}

	for _, line := range lines[2:] {
		instructions[line[:3]] = Instruction{
			left:  line[7:10],
			right: line[12:15],
		}
	}

	step := 0
	instruction := instructions["AAA"]

	for {
		dir := directions[step%len(directions)]
		var newKey string
		switch dir {
		case "L":
			newKey = instruction.left
		case "R":
			newKey = instruction.right
		}
		instruction = instructions[newKey]
		step++

		if newKey == "ZZZ" {
			break
		}

	}

	fmt.Println("Steps:", step)
}
