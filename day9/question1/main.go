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
	fmt.Println(len(vals))

	opLengthMap := getOpLengthMap()

	nextOp := 0

	for {
		op := vals[nextOp]
		idx := nextOp
		if op == halt {
			break
		}

		longCode := false
		c := code{}

		if op > 4 {
			c = parseOpCode(op)
			op = c.op
			longCode = true
		}

		nextOp = nextOp + opLengthMap[op]

		var p1, p2 int

		if longCode {
			p1, p2 = getParams(vals, c, idx)
		} else if c.op == saveAddr || c.op == display || c.op == relativeBase {
			p1, p2 = vals[idx+1], 0
		} else {
			p1, p2 = vals[idx+1], vals[idx+2]
		}
		performOp(op, idx, p1, p2, &nextOp, vals, longCode)
	}
}

func performOp(op, idx, p1, p2 int, nextOp *int, vals []int, longCode bool) {
	switch op {
	case add:
		fmt.Println("add")
		if !longCode {
			p1, p2 = vals[p1], vals[p2]
		}
		addFunc(vals, p1, p2, vals[idx+3])
	case mult:
		fmt.Println("mult")
		if !longCode {
			p1, p2 = vals[p1], vals[p2]
		}
		multFunc(vals, p1, p2, vals[idx+3])
	case saveAddr:
		fmt.Println("saveAddr")
		saveToRegister(vals, p1, inputVal)
	case display:
		fmt.Println("display")
		printRegister(vals, p1)
	case jumpIfTrue:
		fmt.Println("jumpIfTrue")
		if p1 > 0 {
			*nextOp = p2
		}
	case jumpIfFalse:
		fmt.Println("jumpIfFalse")
		if p1 == 0 {
			*nextOp = p2
		}
	case lessThan:
		fmt.Println("lessThan")
		if p1 < p2 {
			saveToRegister(vals, vals[idx+3], 1)
		} else {
			saveToRegister(vals, vals[idx+3], 0)
		}
	case equals:
		fmt.Println("equals")
		if p1 == p2 {
			saveToRegister(vals, vals[idx+3], 1)
		} else {
			saveToRegister(vals, vals[idx+3], 0)
		}
	case adjustRelBase:
		fmt.Println("adjustRelBase")
		relativeBase = relativeBase + p1
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

func getParams(arr []int, c code, idx int) (int, int) {
	p1 := getParam(arr, c.firstMode, idx+1)
	if c.op == saveAddr || c.op == display {
		return p1, 0
	}
	p2 := getParam(arr, c.secondMode, idx+2)
	return p1, p2
}

func getParam(arr []int, mode int, idx int) int {
	var param int

	if mode == positionMode {
		param = arr[arr[idx]]
	} else if mode == relativeMode {
		param = relativeBase + arr[idx]
	} else {
		param = arr[idx]
	}
	return param

}

func getVals(arr []int, c code, longCode bool, idx int) (int, int) {
	var p1, p2 int

	if longCode {
		p1, p2 = getParams(arr, c, idx)
	} else if c.op == saveAddr || c.op == display {
		p1, p2 = arr[idx+1], 0
	} else {
		p1, p2 = arr[idx+1], arr[idx+2]
	}
	return p1, p2
}
