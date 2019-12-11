package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/johnny88/aoc2019/fileparse"
)

const (
	add           = 1
	mult          = 2
	saveAddr      = 3
	display       = 4
	jumpIfTrue    = 5
	jumpIfFalse   = 6
	lessThan      = 7
	equals        = 8
	adjustRelBase = 9
	halt          = 99

	inputVal = 1

	positionMode  = 0
	immediateMode = 1
	relativeMode  = 2
)

type code struct {
	op         int
	firstMode  int
	secondMode int
	thirdMode  int
}

var relativeBase = 0

func main() {
	scanner, err := fileparse.NewScanner("day9/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer scanner.Close()

	vals := scanner.CommaStringParseInt()
	vals = append(vals, make([]int, 9999999)...)

	opLengthMap := getOpLengthMap()

	nextOp := 0

	for {
		op := vals[nextOp]
		idx := nextOp
		if op == halt {
			break
		}

		c := code{
			op: op,
		}
		if op > 4 {
			c = parseOpCode(op)
		}
		fmt.Println(op, c)

		nextOp = nextOp + opLengthMap[op]

		performOp(idx, &nextOp, vals, c)
	}
}

func performOp(idx int, nextOp *int, vals []int, c code) {
	p1, p2, p3 := getParams(vals, c, idx)
	// fmt.Println("index:", idx, "\top:", opToString(op), "\tparam 1:", p1, "\tparam 2:", p2, "\tparam 3:", p3)
	switch c.op {
	case add:
		addFunc(vals, p1, p2, p3)
	case mult:
		multFunc(vals, p1, p2, p3)
	case saveAddr:
		saveToRegister(vals, p1, inputVal)
	case display:
		printRegister(vals, p1)
	case jumpIfTrue:
		if p1 > 0 {
			*nextOp = p2
		}
	case jumpIfFalse:
		if p1 == 0 {
			*nextOp = p2
		}
	case lessThan:
		if p1 < p2 {
			saveToRegister(vals, p3, 1)
		} else {
			saveToRegister(vals, p3, 0)
		}
	case equals:
		if p1 == p2 {
			saveToRegister(vals, p3, 1)
		} else {
			saveToRegister(vals, p3, 0)
		}
	case adjustRelBase:
		relativeBase = relativeBase + p1
		fmt.Println(relativeBase)
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
		add:           4,
		mult:          4,
		saveAddr:      2,
		display:       2,
		jumpIfTrue:    3,
		jumpIfFalse:   3,
		lessThan:      4,
		equals:        4,
		adjustRelBase: 2,
	}
}

func addFunc(a []int, p1 int, p2 int, pos int) {
	a[pos] = p1 + p2
}

func multFunc(a []int, p1 int, p2 int, pos int) {
	a[pos] = p1 * p2
}

func saveToRegister(a []int, pos int, val int) {
	a[pos] = val
}

func printRegister(a []int, pos int) {
	fmt.Println(a[pos])
}

func getParams(arr []int, c code, idx int) (p1, p2, p3 int) {
	switch c.op {
	case add, mult, lessThan, equals:
		p1 = getParam(arr, c.firstMode, idx+1)
		p2 = getParam(arr, c.secondMode, idx+2)
		p3 = getParam(arr, 1, idx+3)
	case saveAddr:
		p1 = getParam(arr, 1, idx+1)
		if c.firstMode == relativeMode {
			p1 += relativeBase
		}
	case jumpIfFalse, jumpIfTrue:
		p1 = getParam(arr, c.firstMode, idx+1)
		p2 = getParam(arr, c.secondMode, idx+2)
	case display, adjustRelBase:
		p1 = getParam(arr, c.firstMode, idx+1)
	}
	return
}

func getParam(arr []int, mode int, idx int) int {
	if mode == positionMode {
		return arr[arr[idx]]
	} else if mode == relativeMode {
		return relativeBase + arr[idx]
	}
	return arr[idx]

}

// func getVals(arr []int, c code, longCode bool, idx int) (int, int) {
// if longCode {
// return getParams(arr, c, idx)
// } else if c.op == saveAddr || c.op == display {
// return arr[idx+1], 0
// }
// return arr[idx+1], arr[idx+2]
// }

func opToString(op int) string {
	opStringMap := map[int]string{
		add:           "add",
		mult:          "mult",
		saveAddr:      "saveAddr",
		display:       "display",
		jumpIfTrue:    "jumpIfTrue",
		jumpIfFalse:   "jumpIfFalse",
		lessThan:      "lessThan",
		equals:        "equals",
		adjustRelBase: "adjustRel",
	}
	return opStringMap[op]
}
