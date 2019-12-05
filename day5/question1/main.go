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
			if longCode {
				p1, p2, p3 := getParams(vals, c, idx)
				addFunc(vals, p1, p2, p3)
			} else {
				addFunc(vals, vals[vals[idx+1]], vals[vals[idx+2]], vals[idx+3])
			}
		case mult:
			if longCode {
				p1, p2, p3 := getParams(vals, c, idx)
				multFunc(vals, p1, p2, p3)
			} else {
				multFunc(vals, vals[vals[idx+1]], vals[vals[idx+2]], vals[idx+3])
			}
		case saveAddr:
			if longCode {
				p1, _, _ := getParams(vals, c, idx)
				saveToRegister(vals, p1)
			} else {
				saveToRegister(vals, vals[idx+1])
			}
		case display:
			if longCode {
				p1, _, _ := getParams(vals, c, idx)
				printRegister(vals, p1)
				saveToRegister(vals, p1)
			} else {
				printRegister(vals, vals[idx+1])
			}
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

func getParams(arr []int, c code, idx int) (int, int, int) {
	var p1, p2, p3 int
	if c.firstMode == positionMode {
		p1 = arr[arr[idx+1]]
	} else {
		p1 = arr[idx+1]
	}
	if c.secondMode == positionMode {
		p1 = arr[arr[idx+2]]
	} else {
		p1 = arr[idx+2]
	}
	if c.thirdMode == positionMode {
		p1 = arr[arr[idx+3]]
	} else {
		p1 = arr[idx+3]
	}
	return p1, p2, p3
}
