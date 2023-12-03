package main

import (
	_ "embed"
	"fmt"
	"github.com/omarshaarawi/advent-of-code/util"
	"regexp"
	"strconv"
)

//go:embed data1.txt
var data1 []byte

func main() {
	data, err := util.ReadInput(data1)
	if err != nil {
		fmt.Printf("error reading input: %v\n", err)
		return
	}
	processGame(data, parseColors)
	processGame(data, parseColors2)
}

func processGame(data []string, processor func(string) int) {
	var sum int
	for i := range data {
		game := processor(data[i])
		sum += game
	}
	fmt.Println(sum)
}

func parseColors2(input string) int {
	gameNumRegex := regexp.MustCompile(`Game (\d+):`)
	match := gameNumRegex.FindStringSubmatch(input)
	if len(match) < 2 {
		fmt.Println("Game number not found")
		return 0
	}
	colorNumRegex := regexp.MustCompile(`(\d+) (\w+)`)
	matches := colorNumRegex.FindAllStringSubmatch(input, -1)
	var red int
	var green int
	var blue int
	for _, m := range matches {
		if len(m) < 3 {
			continue
		}

		number, err := strconv.Atoi(m[1])
		if err != nil {
			fmt.Printf("Invalid number in pair: %v\n", m)
			continue
		}

		switch m[2] {
		case "red":
			if red <= number {
				red = number
			}
		case "green":
			if green <= number {
				green = number
			}
		case "blue":
			if blue <= number {
				blue = number
			}
		}
	}

	return red * green * blue
}

func parseColors(input string) int {
	gameNumRegex := regexp.MustCompile(`Game (\d+):`)
	match := gameNumRegex.FindStringSubmatch(input)
	if len(match) < 2 {
		fmt.Println("Game number not found")
		return 0
	}

	gameNumber, _ := strconv.Atoi(match[1])

	colorNumRegex := regexp.MustCompile(`(\d+) (\w+)`)
	matches := colorNumRegex.FindAllStringSubmatch(input, -1)
	for _, m := range matches {
		if len(m) < 3 {
			continue
		}

		number, err := strconv.Atoi(m[1])
		if err != nil {
			fmt.Printf("Invalid number in pair: %v\n", m)
			continue
		}

		// 12 red cubes, 13 green cubes, and 14 blue cubes
		switch m[2] {
		case "red":
			if number > 12 {
				return 0
			}
		case "green":
			if number > 13 {
				return 0
			}
		case "blue":
			if number > 14 {
				return 0
			}
		}
	}

	return gameNumber
}
