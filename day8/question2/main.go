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
	type layer map[int]int
	pic := []layer{}
	for i, _ := range arr {
		if i%layerLength == 0 {
			pic = append(pic, layer{})
		}
		lIdx := int(math.Floor(float64(i) / layerLength))
		rowIdx := int(math.Floor(float64(i-lIdx*layerLength) / 25))
		colIdx := i - rowIdx*25 - lIdx*layerLength
		l := pic[0]
		if _, ok := l[rowIdx]; !ok {

		}

		// d := m[lIdx]
		// d[el]++
	}

	// min := 99999999999999
	// layer := 0
	// for idx, d := range m {
	// if min > d[0] {
	// min = d[0]
	// layer = idx
	// }
	// }
	// fmt.Println(m[layer][1] * m[layer][2])
	fmt.Println()
}
