package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

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

	inputVal = 2

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
	scanner, err := fileparse.NewScanner("day17/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer scanner.Close()

	vals := scanner.CommaStringParseInt()
	vals = append(vals, make([]int, 9999999)...)
	vals[0] = 2

	opLengthMap := getOpLengthMap()

	nextOp := 0

	currInput := 0
	currOut := 0
	str := ""
	for {
		idx := nextOp
		c := code{op: vals[nextOp]}
		if vals[nextOp] > 4 {
			c = parseOpCode(vals[nextOp])
		}
		if c.op == halt {
			break
		}

		input := []rune{
			'A', ',', 'A', ',', 'B', ',', 'C', ',', 'B', ',', 'C', ',', 'B', ',', 'C', ',', 'B', ',', 'A', '\n',
			'R', ',', '1', '0', ',', 'L', ',', '1', '2', ',', 'R', ',', '6', '\n',
			'R', ',', '6', ',', 'R', ',', '1', '0', ',', 'R', ',', '1', '2', ',', 'R', ',', '6', '\n',
			'R', ',', '1', '0', ',', 'L', ',', '1', '2', ',', 'L', ',', '1', '2', '\n',
			'y', '\n',
		}
		nextOp = nextOp + opLengthMap[c.op]

		performOp(idx, &nextOp, vals, c, input, &currInput, &currOut, &str)
	}
}

func performOp(idx int, nextOp *int, vals []int, c code, input []rune, i, j *int, s *string) {
	p1, p2, p3 := getParams(vals, c, idx)
	switch c.op {
	case add:
		vals[p3] = p1 + p2
	case mult:
		vals[p3] = p1 * p2
	case saveAddr:
		fmt.Println(*i)
		vals[p1] = int(input[*i])
		*i++
	case display:
		if p1 == '.' || p1 == '#' || p1 == '^' || p1 == 'v' || p1 == '>' || p1 == '<' || p1 == '\n' {
			tmp := fmt.Sprintf("%c", rune(p1))
			*s = strings.Join([]string{*s, tmp}, "")
			*j++
			if *j == 1717 {
				time.Sleep(30 * time.Millisecond)
				fmt.Printf(*s)
				*s = ""
				*j = 0
			}
		} else {
			// fmt.Println(p1)
		}
	case jumpIfTrue:
		if p1 > 0 {
			*nextOp = p2
		}
	case jumpIfFalse:
		if p1 == 0 {
			*nextOp = p2
		}
	case lessThan:
		var val int
		if p1 < p2 {
			val = 1
		}
		vals[p3] = val
	case equals:
		var val int
		if p1 == p2 {
			val = 1
		}
		vals[p3] = val
	case adjustRelBase:
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

func getParams(arr []int, c code, idx int) (p1, p2, p3 int) {
	switch c.op {
	case add, mult, lessThan, equals:
		p1 = getParam(arr, c.firstMode, idx+1)
		p2 = getParam(arr, c.secondMode, idx+2)
		p3 = getParam(arr, 1, idx+3)
		if c.thirdMode == relativeMode {
			p3 += relativeBase
		}
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
	p := arr[idx]
	if mode == positionMode {
		return arr[p]
	} else if mode == relativeMode {
		return arr[relativeBase+p]
	}
	return p
}

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
