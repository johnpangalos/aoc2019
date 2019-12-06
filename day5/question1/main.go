package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/johnny88/aoc2019/fileparse"
)

const (
	add      = 1
	mult     = 2
	saveAddr = 3
	display  = 4
	halt     = 99

	inputVal = 1

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
			fmt.Println(vals)
			fmt.Println("next op:", nextOp, "idx:", idx)
			fmt.Println("Error: next operation index is less than current index")
		}

		if op == halt || nextOp < idx {
			break
		}

		if idx != nextOp {
			continue
		}

		longCode := false
		c := code{}

		if op > 4 {
			c = parseOpCode(op)
			op = c.op
			longCode = true
		}

		nextOp = nextOp + opLengthMap[op]

		switch op {
		case add:
			var p1, p2 int
			if longCode {
				p1, p2 = getValParams(vals, c, idx)
			} else {
				p1, p2 = vals[vals[idx+1]], vals[vals[idx+2]]
			}
			addFunc(vals, p1, p2, vals[idx+3])
		case mult:
			var p1, p2 int
			if longCode {
				p1, p2 = getValParams(vals, c, idx)
			} else {
				p1, p2 = vals[vals[idx+1]], vals[vals[idx+2]]
			}
			multFunc(vals, p1, p2, vals[idx+3])
		case saveAddr:
			var p1 int
			if longCode {
				p1, _ = getValParams(vals, c, idx)
			} else {
				p1 = vals[idx+1]
			}
			saveToRegister(vals, p1)
		case display:
			var p1 int
			if longCode {
				p1, _ = getValParams(vals, c, idx)
			} else {
				p1 = vals[idx+1]
			}
			printRegister(vals, p1)
		}
	}
}

func parseOpCode(op int) code {
	opStr := strconv.Itoa(op)
	for len(opStr) < 5 {
		opStr = strings.Join([]string{"0", opStr}, "")
	}
	opArr := strings.SplitN(opStr, "", 4)

	op, _ = strconv.Atoi(opArr[3])
	firstMode, _ := strconv.Atoi(opArr[2])
	secondMode, _ := strconv.Atoi(opArr[1])
	thirdMode, _ := strconv.Atoi(opArr[0])

	return code{
		op:         op,
		firstMode:  firstMode,
		secondMode: secondMode,
		thirdMode:  thirdMode,
	}
}

func getOpLengthMap() map[int]int {
	return map[int]int{
		1: 4,
		2: 4,
		3: 2,
		4: 2,
	}
}

func addFunc(a []int, p1 int, p2 int, pos int) {
	a[pos] = p1 + p2
}

func multFunc(a []int, p1 int, p2 int, pos int) {
	a[pos] = p1 * p2
}

func saveToRegister(a []int, pos int) {
	a[pos] = inputVal
}

func printRegister(a []int, pos int) {
	fmt.Println(a[pos])
}

func getValParams(arr []int, c code, idx int) (int, int) {
	p1 := getValParam(arr, c.firstMode, idx+1)
	if c.op == saveAddr || c.op == display {
		return p1, 0
	}
	p2 := getValParam(arr, c.secondMode, idx+2)
	return p1, p2
}

func getValParam(arr []int, mode int, idx int) int {
	var param int

	if mode == positionMode {
		param = arr[arr[idx]]
	} else {
		param = arr[idx]
	}
	return param
}
