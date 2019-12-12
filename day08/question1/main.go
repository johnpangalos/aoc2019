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
	type digitMap map[int]int
	m := []digitMap{}
	for idx, el := range arr {
		if idx%layerLength == 0 {
			m = append(m, digitMap{})
		}
		lIdx := int(math.Floor(float64(idx) / layerLength))
		d := m[lIdx]
		d[el]++
	}

	min := 99999999999999
	layer := 0
	for idx, d := range m {
		if min > d[0] {
			min = d[0]
			layer = idx
		}
	}
	fmt.Println(m[layer][1] * m[layer][2])
}
