package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/johnny88/aoc2019/fileparse"
)

const (
	numPhases = 100
)

type signal []int

func main() {
	scanner, err := fileparse.NewScanner("day16/input.txt")
	if err != nil {
		panic(err)
	}
	defer scanner.Close()
	input := scanner.PCStringParseInt()
	sig := signal{}
	sig = input
	for i := 0; i < numPhases; i++ {
		sums := signal{}
		for i := range sig {
			sum := 0
			pat := repeatingPattern(len(sig), i)
			for j, v := range sig {
				sum += v * pat[j]
			}
			sums = append(sums, int(math.Abs(float64(sum%10))))
		}
		if err != nil {
			panic(err)
		}
		sig = sums
	}
	fmt.Println(sig.toString())
}

func (sig signal) toString() string {
	strArr := []string{}
	for _, v := range sig[0:8] {
		strArr = append(strArr, strconv.Itoa(v))
	}
	s := strings.Join(strArr, "")
	return s

}

func repeatingPattern(length, idx int) []int {
	pattern := []int{0, 1, 0, -1}
	repeatPattern := []int{}
	inc := 0
	count := 0
	for i := 0; i < length+1; i++ {
		if i != 0 {
			repeatPattern = append(repeatPattern, pattern[inc%4])
		}
		if count == idx {
			count = 0
			inc++
		} else {
			count++
		}
	}
	return repeatPattern
}
