package main

import (
	"fmt"

	"github.com/johnny88/aoc2019/fileparse"
)

func main() {
	scanner, file, err := fileparse.NewScanner("day2/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	vals := scanner.CommaStringParseInt()
	// 19690720
	vals[1] = 38
	vals[2] = 92

	for idx, val := range vals {
		if (idx+1)%4 != 0 {
			continue
		}
		if vals[idx-3] == 99 {
			break
		}
		mem1 := vals[idx-3]
		mem2 := vals[idx-2]
		mem3 := vals[idx-1]

		switch mem1 {
		case 1:
			vals[val] = vals[mem2] + vals[mem3]
		case 2:
			vals[val] = vals[mem2] * vals[mem3]
		}
	}
	fmt.Println(vals[0])
}
