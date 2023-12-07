package main

import (
	_ "embed"
	"flag"
	"fmt"

	"github.com/omarshaarawi/advent-of-code/util"
)

//go:embed data.txt
var data []byte

func main() {
	data, err := util.ReadInput(data)
	if err != nil {
		fmt.Printf("error reading input: %v\n", err)
		return
	}

	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()

	switch part {
	case 1:
		fmt.Println("Running part", part)
		ans := part1(data)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	case 2:
		fmt.Println("Running part", part)
		ans := part2(data)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	default:
		fmt.Println("Invalid part number")
	}
}

func part1(data []string) int {
	return 0
}

func part2(data []string) int {
	return 0
}
