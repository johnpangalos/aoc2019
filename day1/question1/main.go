package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/johnny88/aoc2019/fileparse"
)

func main() {

	scanner, err := fileparse.NewScanner("day1/input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	var sum float64
	for scanner.Scan() {
		val, err := strconv.ParseFloat(scanner.Text(), 10)

		fmt.Println(val)
		val = math.Floor(val/3) - 2
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(val)
		sum = sum + val
	}
	fmt.Printf("%.0f\n", sum)
}
