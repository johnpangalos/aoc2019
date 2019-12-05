package main

import (
	"fmt"
	"strconv"

	"github.com/johnny88/aoc2019/fileparse"
)

const (
	add      = 1
	mult     = 2
	saveAddr = 3
	display  = 4
	halt     = 99

	positionMode  = 0
	immediateMode = 1
)

type code struct {
	op         int
	firstMode  int
	secondMode int
	thirdMode  int
}

func main() {
	scanner, err := fileparse.NewScanner("day5/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer scanner.Close()

	vals := scanner.CommaStringParseInt()

	opLengthMap := getOpLengthMap()

	nextOp := 0
	for idx, op := range vals {
		if nextOp < idx {
			fmt.Println("Error: next operation index is less than current index")
		}
		if op == halt || nextOp < idx {
			break
		}
		if idx != nextOp {
			continue
		}

		nextOp = nextOp + opLengthMap[op]
		var c code
		if op > 4 {
			c := parseOpCode(op)
		}
	}
}

func parseOpCode(op int) code {
	strconv.Itoa(op)
	return code{}
}

func getOpLengthMap() map[int]int {
	return map[int]int{
		1: 4,
		2: 4,
		3: 2,
		4: 2,
	}
}

func add(a []int, p1 int, p2 int, pos int) int {
}
