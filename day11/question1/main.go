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

	up    = "up"
	down  = "down"
	left  = "left"
	right = "right"

	turnLeft  = 0
	turnRight = 1

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

type coord struct {
	x, y int
}

type colorMap map[coord]int
type changeMap map[coord]bool

var relativeBase = 0

func main() {
	scanner, err := fileparse.NewScanner("day11/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer scanner.Close()

	vals := scanner.CommaStringParseInt()
	vals = append(vals, make([]int, 9999999)...)

	opLengthMap := getOpLengthMap()

	nextOp := 0
	cMap := make(colorMap)
	chMap := make(changeMap)

	currCoord := coord{x: 0, y: 0}
	cMap[currCoord] = 0
	chMap[currCoord] = false
	currOut := []int{}
	facing := up

	for {
		idx := nextOp
		c := code{op: vals[nextOp]}
		if vals[nextOp] > 4 {
			c = parseOpCode(vals[nextOp])
		}
		if c.op == halt {
			break
		}

		nextOp = nextOp + opLengthMap[c.op]

		out := performOp(idx, &nextOp, vals, c, cMap[currCoord])
		if c.op == display && len(currOut) < 2 {
			currOut = append(currOut, out)
		}
		if len(currOut) == 2 {
			cMap[currCoord] = currOut[0]
			if currOut[0] == 1 {
				chMap[currCoord] = true
			}
			facing = nextDirection(facing, currOut[1])
			currCoord = nextCoord(currCoord, facing)
			if _, ok := cMap[currCoord]; !ok {
				cMap[currCoord] = 0
			}
			currOut = []int{}
		}
	}
	count := 0
	for _, hasChanged := range chMap {
		if hasChanged {
			count++
		}
	}
	fmt.Println(count)
}

func nextDirection(facing string, turn int) (next string) {
	if facing == up && turn == turnLeft {
		next = left
	} else if facing == up && turn == turnRight {
		next = right
	} else if facing == down && turn == turnLeft {
		next = right
	} else if facing == down && turn == turnRight {
		next = left
	} else if facing == left && turn == turnLeft {
		next = down
	} else if facing == left && turn == turnRight {
		next = up
	} else if facing == right && turn == turnLeft {
		next = up
	} else if facing == right && turn == turnRight {
		next = down
	}
	return
}

func nextCoord(c coord, facing string) (next coord) {
	switch facing {
	case up:
		next = coord{
			x: c.x,
			y: c.y + 1,
		}
	case down:
		next = coord{
			x: c.x,
			y: c.y - 1,
		}
	case right:
		next = coord{
			x: c.x + 1,
			y: c.y,
		}
	case left:
		next = coord{
			x: c.x - 1,
			y: c.y,
		}
	}
	return
}

func performOp(idx int, nextOp *int, vals []int, c code, inputVal int) int {
	var out int
	p1, p2, p3 := getParams(vals, c, idx)
	// fmt.Println("idx:", idx, "p1:", p1, "p2:", p2, "p3:", p3, "code:", c)
	switch c.op {
	case add:
		vals[p3] = p1 + p2
	case mult:
		vals[p3] = p1 * p2
	case saveAddr:
		vals[p1] = inputVal
	case display:
		out = p1
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
	return out
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
