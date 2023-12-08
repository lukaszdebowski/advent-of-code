package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("example-1.txt")
	lines := strings.Split(string(data), "\n")

	result := part1Solution(lines)
	fmt.Println("Result:", result)
}

type Game struct {
	id     int
	rounds []Round
}

type Round struct {
	red   int
	green int
	blue  int
}

func part1Solution(lines []string) int {
	allGames := []Game{}

	for _, line := range lines {
		gameId, _ := strconv.Atoi(strings.Split(line[5:], ":")[0])

		info := strings.Split(line, ": ")[1]
		roundsInfo := strings.Split(info, "; ")
		rounds := []Round{}

		for _, roundInfo := range roundsInfo {
			roundInfo = strings.ReplaceAll(roundInfo, ",", "")
			parts := strings.Split(roundInfo, " ")
			redIndex := slices.Index(parts, "red") - 1
			blueIndex := slices.Index(parts, "blue") - 1
			greenIndex := slices.Index(parts, "green") - 1

			var redValue, blueValue, greenValue int

			if redIndex >= 0 {
				redValue, _ = strconv.Atoi(parts[redIndex])
			}
			if blueIndex >= 0 {
				blueValue, _ = strconv.Atoi(parts[blueIndex])
			}
			if greenIndex >= 0 {
				greenValue, _ = strconv.Atoi(parts[greenIndex])
			}

			rounds = append(rounds, Round{redValue, blueValue, greenValue})
		}

		allGames = append(allGames, Game{gameId, rounds})
	}

	result := 0

	for _, game := range allGames {
		validGame := true

		for _, round := range game.rounds {
			if round.red > 12 || round.green > 13 || round.blue > 14 {
				validGame = false
				break
			}
		}

		if validGame {
			result += game.id
		}

	}

	return result
}

func part2Solution(lines []string) int {
	result := 0

	// for _, line := range lines {

	// }

	return result
}
