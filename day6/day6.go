package main

import (
	"fmt"
)

type Race struct {
	time     int
	distance int
}

func main() {
	resultPart1 := getResultFromRaces([]Race{
		{time: 56, distance: 499},
		{time: 97, distance: 2210},
		{time: 77, distance: 1097},
		{time: 93, distance: 1440},
	})

	resultPart2 := getResultFromRaces([]Race{
		{time: 56977793, distance: 499221010971440},
	})

	fmt.Println("resultPart1", resultPart1)
	fmt.Println("resultPart2", resultPart2)

}

func getResultFromRaces(races []Race) int {
	result := 1

	for _, race := range races {
		winningOptions := 0
		for speed := 1; speed < race.time; speed++ {
			if speed*(race.time-speed) > race.distance {
				winningOptions++
			}
		}
		result *= winningOptions
	}

	return result
}
