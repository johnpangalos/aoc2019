package main

import (
	"fmt"

	"github.com/johnny88/aoc2019/fileparse"
)

func main() {
	scanner, err := fileparse.NewScanner("day2/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer scanner.Close()

	vals := scanner.CommaStringParseInt()
	vals[1] = 12
	vals[2] = 2

	for idx, val := range vals {
		if (idx+1)%4 != 0 {
			continue
		}
		if vals[idx-3] == 99 {
			break
		}
		calc(vals, vals[idx-3], vals[idx-2], vals[idx-1], val)
	}
	fmt.Println(vals[0])
}

func calc(vals []int64, op int64, mem1 int64, mem2 int64, mem3 int64) {
	switch op {
	case 1:
		vals[mem3] = vals[mem1] + vals[mem2]
	case 2:
		vals[mem3] = vals[mem1] * vals[mem2]
	}
}
