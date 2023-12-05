package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var categories = []string{"seed", "soil", "fertilizer", "water", "light", "temperature", "humidity", "location"}

func main() {
	data, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")

	seeds := extractSeeds(lines)
	seedToSoilMap := extractSourceToDestinationMap(lines, "seed", "soil")
	soilToFertilizerMap := extractSourceToDestinationMap(lines, "soil", "fertilizer")
	fertilizerToWaterMap := extractSourceToDestinationMap(lines, "fertilizer", "water")
	waterToLightMap := extractSourceToDestinationMap(lines, "water", "light")
	lightToTemperatureMap := extractSourceToDestinationMap(lines, "light", "temperature")
	temperatureToHumidityMap := extractSourceToDestinationMap(lines, "temperature", "humidity")
	humidityToLocationMap := extractSourceToDestinationMap(lines, "humidity", "location")

	locations := []int{}

	for _, seed := range seeds {
		soil, soilFound := seedToSoilMap[seed]
		if !soilFound {
			soil = seed
		}

		fertilizer, fertilizerFound := soilToFertilizerMap[soil]
		if !fertilizerFound {
			fertilizer = soil
		}
		water, waterFound := fertilizerToWaterMap[fertilizer]
		if !waterFound {
			water = fertilizer
		}
		light, lightFound := waterToLightMap[water]
		if !lightFound {
			light = water
		}
		temperature, temperatureFound := lightToTemperatureMap[light]
		if !temperatureFound {
			temperature = light
		}
		humidity, humidityFound := temperatureToHumidityMap[temperature]
		if !humidityFound {
			humidity = temperature
		}
		location, locationFound := humidityToLocationMap[humidity]
		if !locationFound {
			location = humidity
		}

		locations = append(locations, location)
		fmt.Println("Seed", seed, "soil", soil, "fertilizer", fertilizer, "water", water, "light", light, "temperature", temperature, "humidity", humidity, "location", location)

	}

	fmt.Println("Lowest location is", slices.Min(locations))

}

func extractSeeds(lines []string) []int {
	stringSeeds := strings.Split(lines[0], " ")[1:]
	intSeeds := make([]int, len(stringSeeds))
	for i, seed := range stringSeeds {
		intSeeds[i], _ = strconv.Atoi(seed)
	}

	return intSeeds
}

func extractSourceToDestinationMap(lines []string, source, destination string) map[int]int {
	result := make(map[int]int)

	startLineIndex := slices.IndexFunc(lines, func(line string) bool {
		return strings.Contains(line, fmt.Sprintf("%s-to-%s map", source, destination))
	}) + 1

	var endLineIndex int

	if destination == categories[len(categories)-1] {
		// last category
		endLineIndex = len(lines) - 1
	} else {
		endLineIndex = slices.IndexFunc(lines, func(line string) bool {
			return strings.Contains(line, fmt.Sprintf(
				"%s-to-%s map",
				destination,
				categories[slices.Index(categories, destination)+1],
			))
		}) - 2
	}

	mapLines := lines[startLineIndex : endLineIndex+1]

	for _, line := range mapLines {
		destinationRangeStart, _ := strconv.Atoi(strings.Split(line, " ")[0])
		sourceRangeStart, _ := strconv.Atoi(strings.Split(line, " ")[1])
		rangeLength, _ := strconv.Atoi(strings.Split(line, " ")[2])
		for i := 0; i < rangeLength; i++ {
			result[sourceRangeStart+i] = destinationRangeStart + i
		}
	}
	return result
}
