package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	cards []string
	bid   int
}

type Hands []Hand

func (h Hands) Len() int {
	return len(h)
}

func (h Hands) Less(i, j int) bool {
	elem1 := h[i]
	elem2 := h[j]

	if elem1.getType() < elem2.getType() {
		return true
	} else if elem1.getType() == elem2.getType() {
		for i := 0; i < len(elem1.cards); i++ {
			fmt.Println(elem1.cards[i], cardsMapping[elem1.cards[i]])
			fmt.Println(elem2.cards[i], cardsMapping[elem2.cards[i]])
			if cardsMapping[elem1.cards[i]] < cardsMapping[elem2.cards[i]] {
				return true
			} else if cardsMapping[elem1.cards[i]] == cardsMapping[elem2.cards[i]] {
				continue
			} else {
				return false
			}
		}
	}
	return false

}
func (h Hands) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Hand) getType() int {
	cardsMap := map[string]int{}

	for _, card := range h.cards {
		cardsMap[card]++
	}

	values := make([]int, 0, len(cardsMap))
	for _, v := range cardsMap {
		values = append(values, v)
	}

	slices.Sort(values)

	if slices.Equal(values, []int{5}) {
		return 7
	}
	if slices.Equal(values, []int{1, 4}) {
		return 6
	}
	if slices.Equal(values, []int{2, 3}) {
		return 5
	}
	if slices.Equal(values, []int{1, 1, 3}) {
		return 4
	}
	if slices.Equal(values, []int{1, 2, 2}) {
		return 3
	}
	if slices.Equal(values, []int{1, 1, 1, 2}) {
		return 2
	}
	if slices.Equal(values, []int{1, 1, 1, 1, 1}) {
		return 1
	}
	fmt.Println("ERROR")
	return 0
}

var cardsMapping = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")

	hands := Hands([]Hand{})
	for _, line := range lines {
		cards := strings.Split(strings.Split(line, " ")[0], "")
		bid, _ := strconv.Atoi(strings.Split(line, " ")[1])
		hands = append(hands, Hand{cards: cards, bid: bid})
	}

	sort.Sort(hands)
	result := 0

	for _, hand := range hands {
		fmt.Println(hand)
	}

	for index, hand := range hands {
		value := hand.bid * (index + 1)
		result += value
	}

	fmt.Println(result)
}

// 248143252
