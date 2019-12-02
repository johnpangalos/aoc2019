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
	defer scanner.Close()

	var sum float64
	for scanner.Scan() {
		val, err := strconv.ParseFloat(scanner.Scanner.Text(), 10)

		val = math.Floor(val/3) - 2
		if err != nil {
			fmt.Println(err)
			continue
		}
		sum = sum + val
	}
	fmt.Printf("%.0f\n", sum)
}
