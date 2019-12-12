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
	pic := [][]int{}
	finalPic := []int{}
	numLayers := 0

	for i, el := range arr {
		if i%layerLength == 0 {
			pic = append(pic, []int{})
			numLayers++
		}
		lIdx := int(math.Floor(float64(i) / layerLength))
		pic[lIdx] = append(pic[lIdx], el)
	}

	finalPic = pic[0]
	for j := 1; j < numLayers; j++ {
		for i := 0; i < layerLength; i++ {
			val := finalPic[i]
			if val == 0 || val == 1 {
				continue
			}
			finalPic[i] = pic[j][i]
		}
	}

	fmt.Println(finalPic)
	for i, el := range finalPic {
		if i%25 == 0 {
			fmt.Printf("\n")
		}
		if el == 2 {
			fmt.Printf(" ")
		} else if el == 1 {
			fmt.Printf("█")
		} else {
			fmt.Printf("░")
		}
	}

}
