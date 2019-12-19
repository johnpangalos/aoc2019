package main

import (
	"fmt"
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
	scanner, err := fileparse.NewScanner("day16/input.txt")
	if err != nil {
		panic(err)
	}
	defer scanner.Close()
	t := ""
	for scanner.Scan() {
		t = scanner.Text()
	}
	var ts [repeat]string

	for i := 0; i < repeat; i++ {
		ts[i] = t
	}
	tFull := strings.Join(ts[:], "")
	sig := signal{}
	for _, v := range tFull {
		i, err := strconv.Atoi(string(v))
		if err != nil {
			panic(err)
		}
		sig = append(sig, i)
	}

	offestArr := sig[0:7]
	offset, err := strconv.Atoi(offestArr.toString())
	if err != nil {
		panic(err)
	}

	sig = sig[offset:]
	for phase := 0; phase < numPhases; phase++ {
		tmp := make(signal, len(sig))
		copy(tmp, sig)

		sum := 0
		for i := len(sig) - 1; i >= 0; i-- {
			sum += sig[i]
			tmp[i] = sum % 10
		}
		sig = tmp
	}
	fmt.Println(sig.toStringOffset(0))
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func (sig signal) toStringOffset(offset int) string {
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

func inc(level, idx int) int {
	return int(float64(idx+1)/float64(level+1)) % 4
}

func multiplier(level, idx int) int {
	pattern := []int{0, 1, 0, -1}
	i := inc(level, idx)
	return pattern[i]
}
