package main

import (
	"fmt"
	"math"

	"github.com/johnny88/aoc2019/fileparse"
)

const layerLength = 25 * 6

func main() {
	s, err := fileparse.NewScanner("day8/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer s.Close()

	arr := s.PCStringParseInt()

	// max := 0
	// layer := 0
	for idx, el := range arr {
		if el == 0 {
			lIdx := math.Floor(float64(idx) / layerLength)

		}
	}

}
