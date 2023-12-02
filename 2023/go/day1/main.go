package main

import (
	_ "embed"
	"fmt"
	"github.com/omarshaarawi/advent-of-code/util"
	"strconv"
	"unicode"
)

//go:embed data1.txt
var data1 []byte

//go:embed data2.txt
var data2 []byte

func main() {
	processInput(data1, processDigits)
	processInput(data2, processWords)
}

func sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func processInput(data []byte, processor func(string) int) {
	input, err := util.ReadInput(data)
	if err != nil {
		fmt.Printf("error reading input: %v\n", err)
		return
	}
	total := 0
	for _, line := range input {
		total += processor(line)
	}
	fmt.Println(total)
}

func processDigits(input string) int {
	var first, last string
	for _, r := range input {
		if unicode.IsNumber(r) {
			if first == "" {
				first = string(r)
			}
			last = string(r)
		}
	}
	if combined, err := strconv.Atoi(first + last); err == nil {
		return combined
	}
	return 0
}

func processWords(input string) int {
	numbers := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"ten":   10,
	}
	first, last, err := extractNumbers(input, numbers)
	if err != nil {
		return 0
	}
	if result, err := strconv.Atoi(fmt.Sprintf("%d%d", first, last)); err == nil {
		return result
	}
	return 0
}

func extractNumbers(input string, numbersMap map[string]int) (firstNumber int, lastNumber int, err error) {
	var firstFound bool
	inputLength := len(input)

	for i := 0; i < inputLength; i++ {
		if unicode.IsDigit(rune(input[i])) {
			f, _ := strconv.Atoi(string(input[i]))
			lastNumber = f

			if firstNumber == 0 {
				firstNumber = f
				firstFound = true
			}
			continue
		}
		for word, number := range numbersMap {
			wordLength := len(word)
			if i+wordLength <= inputLength && input[i:i+wordLength] == word {
				if !firstFound {
					firstNumber = number
					firstFound = true
				}
				lastNumber = number
			}
		}
	}

	if !firstFound {
		return 0, 0, fmt.Errorf("no valid number word found in input")
	}

	return firstNumber, lastNumber, nil
}
