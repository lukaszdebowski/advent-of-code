package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")

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
			greenIndex := slices.Index(parts, "green") - 1
			blueIndex := slices.Index(parts, "blue") - 1

			var redValue, greenValue, blueValue int

			if redIndex >= 0 {
				redValue, _ = strconv.Atoi(parts[redIndex])
			}
			if greenIndex >= 0 {
				greenValue, _ = strconv.Atoi(parts[greenIndex])
			}
			if blueIndex >= 0 {
				blueValue, _ = strconv.Atoi(parts[blueIndex])
			}

			rounds = append(rounds, Round{redValue, greenValue, blueValue})
		}

		allGames = append(allGames, Game{gameId, rounds})
	}

	// result := part1Solution(allGames)
	result := part2Solution(allGames)
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

func part1Solution(allGames []Game) int {
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

func part2Solution(allGames []Game) int {
	result := 0

	for _, game := range allGames {
		reds := []int{}
		greens := []int{}
		blues := []int{}

		for _, round := range game.rounds {
			reds = append(reds, round.red)
			greens = append(greens, round.green)
			blues = append(blues, round.blue)
		}

		maxRed := slices.Max(reds)
		maxGreen := slices.Max(greens)
		maxBlue := slices.Max(blues)

		power := maxRed * maxGreen * maxBlue
		result += power
	}

	return result
}
