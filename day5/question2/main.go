package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/johnny88/aoc2019/fileparse"
)

const (
	add         = 1
	mult        = 2
	saveAddr    = 3
	display     = 4
	jumpIfTrue  = 5
	jumpIfFalse = 6
	lessThan    = 7
	equals      = 8
	halt        = 99

	inputVal = 5

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
			fmt.Println("next op:", nextOp, "idx:", idx)
			fmt.Println("Error: next operation index is less than current index")
			break
		}
		fmt.Println("idx:", idx, "op:", op)

		if op == halt && nextOp == op {
			fmt.Println("halt index:", idx)
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
			fmt.Println("add")
			if longCode {
				p1, p2 := getParams(vals, c, idx)
				addFunc(vals, p1, p2, vals[idx+3])
			} else {
				addFunc(vals, vals[vals[idx+1]], vals[vals[idx+2]], vals[idx+3])
			}
		case mult:
			fmt.Println("mult")
			if longCode {
				p1, p2 := getParams(vals, c, idx)
				multFunc(vals, p1, p2, vals[idx+3])
			} else {
				multFunc(vals, vals[vals[idx+1]], vals[vals[idx+2]], vals[idx+3])
			}
		case saveAddr:
			fmt.Println("saveAddr")
			if longCode {
				p1, _ := getParams(vals, c, idx)
				saveToRegister(vals, p1, inputVal)
			} else {
				saveToRegister(vals, vals[idx+1], inputVal)
			}
		case display:
			fmt.Println("display")
			if longCode {
				p1, _ := getParams(vals, c, idx)
				printRegister(vals, p1)
			} else {
				printRegister(vals, vals[idx+1])
			}
		case jumpIfTrue:
			fmt.Println("jumpIfTrue")
			if longCode {
				p1, p2 := getParams(vals, c, idx)
				fmt.Println("params", p1, p2)
				if p1 > 0 {
					nextOp = p2
				}
			} else {
				p1, p2 := vals[idx+1], vals[idx+2]
				fmt.Println("params", p1, p2)
				if p1 > 0 {
					nextOp = p2
				}
			}
		case jumpIfFalse:
			fmt.Println("jumpIfFalse")
			if longCode {
				p1, p2 := getParams(vals, c, idx)
				if p1 == 0 {
					nextOp = p2
				}
			} else {
				p1, p2 := vals[idx+1], vals[idx+2]
				if p1 == 0 {
					nextOp = p2
				}
				printRegister(vals, vals[idx+1])
			}
		case lessThan:
			fmt.Println("lessThan")
			if longCode {
				p1, p2 := getParams(vals, c, idx)
				if p1 < p2 {
					saveToRegister(vals, vals[idx+3], 1)
				} else {
					saveToRegister(vals, vals[idx+3], 0)
				}
			} else {
				p1, p2 := vals[idx+1], vals[idx+2]
				if p1 < p2 {
					saveToRegister(vals, vals[idx+3], 1)
				} else {
					saveToRegister(vals, vals[idx+3], 0)
				}
			}
		case equals:
			fmt.Println("equals")
			if longCode {
				p1, p2 := getParams(vals, c, idx)
				if p1 == p2 {
					saveToRegister(vals, vals[idx+3], 1)
				} else {
					saveToRegister(vals, vals[idx+3], 0)
				}
			} else {
				p1, p2 := vals[idx+1], vals[idx+2]
				if p1 == p2 {
					saveToRegister(vals, vals[idx+3], 1)
				} else {
					saveToRegister(vals, vals[idx+3], 0)
				}
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
		5: 3,
		6: 3,
		7: 4,
		8: 4,
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
	} else {
		param = arr[idx]
	}
	return param

}
