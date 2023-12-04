package main

import (
	_ "embed"
	"fmt"
	"github.com/omarshaarawi/advent-of-code/util"
	"strconv"
	"strings"
)

//go:embed data.txt
var data []byte

type Card struct {
	id           int
	totalMatches int
	copies       int
}

func main() {
	data, err := util.ReadInput(data)
	if err != nil {
		fmt.Printf("error reading input: %v\n", err)
		return
	}

	totalPoints(data)
	totalScratchCards(data)
}

func totalScratchCards(data []string) {
	cards := make([]Card, len(data))

	for index, line := range data {
		split := strings.Split(line, ":")
		card, _ := strconv.Atoi(strings.Replace(split[0], "Card ", "", -1))
		matches := strings.Split(split[1], "|")
		winning := toInt(strings.Split(matches[0], " "))
		numbers := toInt(strings.Split(matches[1], " "))
		totalMatches := countMatches(winning, numbers)

		cards[index] = Card{
			id:           card,
			totalMatches: totalMatches,
			copies:       1,
		}
	}

	for i := 0; i < len(cards); i++ {
		for j := 1; j <= cards[i].totalMatches && (i+j) < len(cards); j++ {
			cards[i+j].copies += cards[i].copies
		}
	}

	totalCards := 0
	for _, card := range cards {
		totalCards += card.copies
	}

	fmt.Println("Total number of cards:", totalCards)
}

func totalPoints(data []string) {

	var total int

	for _, line := range data {
		split := strings.Split(line, ":")
		matches := strings.Split(split[1], "|")
		winning := toInt(strings.Split(matches[0], " "))
		numbers := toInt(strings.Split(matches[1], " "))
		total += calculateScore(countMatches(winning, numbers))
	}
	fmt.Println(total)
}

func calculateScore(matches int) int {
	if matches == 0 {
		return 0
	}
	score := 1
	for i := 1; i < matches; i++ {
		score *= 2
	}
	return score
}

func toInt(strs []string) []int {
	var ints []int
	for _, s := range strs {
		num, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		ints = append(ints, num)
	}
	return ints
}

func countMatches(l1, l2 []int) int {
	var matchCount int
	for _, num2 := range l2 {
		for _, num1 := range l1 {
			if num1 == num2 {
				matchCount++
				break
			}
		}
	}
	return matchCount
}
