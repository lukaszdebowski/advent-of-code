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
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	seeds := extractSeeds(lines)
	fmt.Println("Seeds length", len(seeds))
	locations := []int{}

	soilMapping := extractMapping(lines, "seed", "soil")
	fertilizerMapping := extractMapping(lines, "soil", "fertilizer")
	waterMapping := extractMapping(lines, "fertilizer", "water")
	lightMapping := extractMapping(lines, "water", "light")
	temperatureMapping := extractMapping(lines, "light", "temperature")
	humidityMapping := extractMapping(lines, "temperature", "humidity")
	locationMapping := extractMapping(lines, "humidity", "location")

	for _, seed := range seeds {
		soil, _ := soilMapping.InRange(seed)
		fertilizer, _ := fertilizerMapping.InRange(soil)
		water, _ := waterMapping.InRange(fertilizer)
		light, _ := lightMapping.InRange(water)
		temperature, _ := temperatureMapping.InRange(light)
		humidity, _ := humidityMapping.InRange(temperature)
		location, _ := locationMapping.InRange(humidity)
		locations = append(locations, location)
	}

	fmt.Println("Smallest location is", slices.Min(locations))

}

func extractSeeds(lines []string) []int {
	stringSeeds := strings.Split(lines[0], " ")[1:]
	intSeeds := []int{}
	// TODO: use goroutines
	for i := 0; i < len(stringSeeds); i = i + 2 {
		beginRange, _ := strconv.Atoi(stringSeeds[i])
		rangeLength, _ := strconv.Atoi(stringSeeds[i+1])

		fmt.Println("beginRange", beginRange, "rangeLength", rangeLength)

		for j := 0; j < rangeLength; j++ {
			intSeeds = append(intSeeds, beginRange+j)
		}
	}

	return intSeeds
}

func extractMapping(lines []string, source, destination string) Mapping {
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

	mapping := Mapping{
		SourceRangeStartValues:      []int{},
		DestinationRangeStartValues: []int{},
		RangeLengthValues:           []int{},
	}

	for _, line := range lines[startLineIndex : endLineIndex+1] {
		lineValues := strings.Split(line, " ")
		destinationRangeStart, _ := strconv.Atoi(lineValues[0])
		sourceRangeStart, _ := strconv.Atoi(lineValues[1])
		rangeLength, _ := strconv.Atoi(lineValues[2])

		mapping.DestinationRangeStartValues = append(mapping.DestinationRangeStartValues, destinationRangeStart)
		mapping.SourceRangeStartValues = append(mapping.SourceRangeStartValues, sourceRangeStart)
		mapping.RangeLengthValues = append(mapping.RangeLengthValues, rangeLength)
	}

	return mapping
}

type Mapping struct {
	SourceRangeStartValues      []int
	DestinationRangeStartValues []int
	RangeLengthValues           []int
}

func (m *Mapping) InRange(value int) (int, bool) {
	for index, sourceRangeStart := range m.SourceRangeStartValues {
		if value >= sourceRangeStart && value <= sourceRangeStart+m.RangeLengthValues[index] {
			result := m.DestinationRangeStartValues[index] + (value - m.SourceRangeStartValues[index])
			return result, true
		}
	}

	return value, false
}
