package main

import (
	_ "embed"
	"fmt"
	"github.com/omarshaarawi/advent-of-code/util"
	"testing"
)

//go:embed data.txt
var data1 []byte

func BenchmarkTotalScratchCards(b *testing.B) {
	data1, err := util.ReadInput(data)
	if err != nil {
		fmt.Printf("error reading input: %v\n", err)
		return
	}
	for i := 0; i < b.N; i++ {
		totalScratchCards(data1)
	}
}
