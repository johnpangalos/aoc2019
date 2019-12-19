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
	repeat    = 10000
)

type signal []int

func main() {
	scanner, err := fileparse.NewScanner("day16/question2/test1.txt")
	if err != nil {
		panic(err)
	}
	defer scanner.Close()
	input := scanner.PCStringParseInt()
	sig := signal{}
	sig = input
	offestArr := sig[0:7]
	offset, err := strconv.Atoi(offestArr.toString())
	if err != nil {
		panic(err)
	}

	sig = append(sig[(offset-1)%len(sig):], sig[:(offset-1)%len(sig)]...)
	for i := 0; i < numPhases; i++ {
		sums := signal{}
		for i := range sig {
			sum := 0
			pat := repeatingPattern(len(sig), i, offset)
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

func (sig signal) toStringOffset(offset int) string {
	fmt.Println(len(sig), offset)
	strArr := []string{}
	for _, v := range sig[offset : offset+8] {
		strArr = append(strArr, strconv.Itoa(v))
	}
	s := strings.Join(strArr, "")
	return s
}
func (sig signal) toString() string {
	strArr := []string{}
	for _, v := range sig {
		strArr = append(strArr, strconv.Itoa(v))
	}
	s := strings.Join(strArr, "")
	return s
}

func repeatingPattern(length, idx, offset int) []int {
	pattern := []int{0, 1, 0, -1}
	repeatPattern := []int{}
	inc := int(math.Floor(float64(offset) / 4))
	count := (offset % (idx + 1*4)) % 4
	fmt.Println(inc, count)
	for i := 0; i < length+1; i++ {
		repeatPattern = append(repeatPattern, pattern[inc%4])
		if count == idx {
			count = 0
			inc++
		} else {
			count++
		}
	}
	return repeatPattern
}
