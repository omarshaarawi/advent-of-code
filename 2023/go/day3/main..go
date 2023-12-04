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

func main() {
	data, err := util.ReadInput(data)
	if err != nil {
		fmt.Printf("error reading input: %v\n", err)
		return
	}

	sum := sumAdjacentNumbers(data)
	fmt.Println("Sum of all part numbers:", sum)

	gearRatios := findAdjacentNumbersToStars(data)
	fmt.Println("Sum of all gear ratio: ", gearRatios)

}

func sumAdjacentNumbers(schematic []string) int {
	sum := 0
	for y, line := range schematic {
		for x := 0; x < len(line); {
			if line[x] >= '0' && line[x] <= '9' {
				start := x
				for x < len(line) && line[x] >= '0' && line[x] <= '9' {
					x++
				}
				end := x - 1
				number := line[start : end+1]

				if num, err := strconv.Atoi(number); err == nil && isAdjacentToSymbol(schematic, start, end, y) {
					sum += num
				}
			} else {
				x++
			}
		}
	}
	return sum
}

func isAdjacentToSymbol(schematic []string, start, end, y int) bool {
	for x := start; x <= end; x++ {
		if checkAdjacent(schematic, x, y) {
			return true
		}
	}
	return false
}
func checkAdjacent(schematic []string, x, y int) bool {
	directions := []struct{ dx, dy int }{
		{-1, -1}, {0, -1}, {1, -1}, // Above
		{-1, 0}, {1, 0}, // Sides
		{-1, 1}, {0, 1}, {1, 1}, // Below
	}
	for _, dir := range directions {
		nx, ny := x+dir.dx, y+dir.dy
		if ny >= 0 && ny < len(schematic) && nx >= 0 && nx < len(schematic[ny]) && isSymbol(rune(schematic[ny][nx])) {
			return true
		}
	}
	return false
}

func findAdjacentNumbersToStars(schematic []string) int {
	var gearRatios []int
	for y, line := range schematic {
		for x, ch := range line {
			if ch == '*' {
				foundNumbers := findTwoAdjacentNumbers(schematic, x, y)
				if len(foundNumbers) == 2 {
					gearRatio := foundNumbers[0] * foundNumbers[1]
					gearRatios = append(gearRatios, gearRatio)
				}
			}
		}
	}
	return sum(gearRatios)
}

func findTwoAdjacentNumbers(schematic []string, x, y int) []int {
	numberSet := make(map[int]struct{})
	directions := []struct{ dx, dy int }{
		{-1, -1}, {0, -1}, {1, -1}, // Above
		{-1, 0}, {1, 0}, // Sides
		{-1, 1}, {0, 1}, {1, 1}, // Below
	}

	for _, dir := range directions {
		nx, ny := x+dir.dx, y+dir.dy
		if ny >= 0 && ny < len(schematic) && nx >= 0 && nx < len(schematic[ny]) {
			numberStr := extractNumber(schematic, ny, nx)
			if number, err := strconv.Atoi(numberStr); err == nil && numberStr != "" {
				numberSet[number] = struct{}{}
			}
		}
	}

	if len(numberSet) == 2 {
		var numbers []int
		for num := range numberSet {
			numbers = append(numbers, num)
		}
		return numbers
	}
	return nil
}

func extractNumber(schematic []string, y, x int) string {
	line := schematic[y]
	start, end := x, x

	// Check if current position is a digit; if not, return empty string
	if !isDigit(rune(line[x])) {
		return ""
	}

	// Move backwards to find the start of the number
	for start > 0 && isDigit(rune(line[start-1])) {
		start--
	}

	// Move forwards to find the end of the number
	for end < len(line)-1 && isDigit(rune(line[end+1])) {
		end++
	}

	return line[start : end+1]
}
func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func isSymbol(r rune) bool {
	return !strings.ContainsRune("0123456789.", r)
}

func sum(arr []int) int {
	sum := 0
	for _, valueInt := range arr {
		sum += valueInt
	}
	return sum
}
