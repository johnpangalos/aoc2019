package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/johnny88/aoc2019/fileparse"
)

func main() {
	scanner, file, err := fileparse.NewScanner("day1/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var sum float64
	for scanner.Scan() {
		val, err := strconv.ParseFloat(scanner.Text(), 10)

		val = math.Floor(val/3) - 2
		if err != nil {
			fmt.Println(err)
			continue
		}
		sum = sum + val
	}
	fmt.Printf("%.0f\n", sum)
}
