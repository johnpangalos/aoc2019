package main

import (
	"fmt"
	"os"
	"os/exec"
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

	positionMode  = 0
	immediateMode = 1
	relativeMode  = 2

	empty            = 0
	wall             = 1
	block            = 2
	horizontalPaddle = 3
	ball             = 4
)

type coord struct {
	x, y int
}

type board map[coord]int

func main() {
	scanner, err := fileparse.NewScanner("day13/input.txt")
	if err != nil {
		panic(err)
	}
	defer scanner.Close()

	vals := scanner.CommaStringParseInt()
	vals[0] = 2

	err = RunIntcode(0, vals)
	if err != nil {
		panic(err)
	}
}

func createGameObjects(output []int) board {
	currCoord := coord{}
	gameObjects := make(board)
	for idx, v := range output {
		if (idx)%3 == 0 {
			currCoord.x = v
		}
		if (idx)%3 == 1 {
			currCoord.y = v
		}
		if (idx)%3 == 2 {
			gameObjects[currCoord] = v
			currCoord = coord{}
		}
	}
	return gameObjects
}

func (b board) toString() string {
	xMax, yMax := 0, 0
	for c := range b {
		if c.x > xMax {
			xMax = c.x
		}
		if c.y > yMax {
			yMax = c.y
		}
	}
	s := fmt.Sprintf("Score: %d\n", b[coord{x: -1, y: 0}])
	for j := 0; j < yMax+1; j++ {
		for i := 0; i < xMax+1; i++ {
			switch b[coord{x: i, y: j}] {
			case empty:
				s = join(s, " ")
			case wall:
				s = join(s, "▓")
			case block:
				s = join(s, "█")
			case horizontalPaddle:
				s = join(s, "▀")
			case ball:
				s = join(s, "◯")
			}
		}
		s = join(s, "\n")
	}
	return s
}

func join(s, s2 string) string {
	return strings.Join([]string{s, s2}, "")
}

func clearConsole() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func getDir(s string) int {
	var i int
	switch s {
	case "a":
		i = -1
	case "s":
		i = 0
	case "d":
		i = 1
	}
	return i
}

type code struct {
	op         int
	firstMode  int
	secondMode int
	thirdMode  int
}

var relativeBase = 0

// RunIntcode runs the int code computer as stipulated and returns
// and output of an array of integers
func RunIntcode(inputVal int, vals []int) error {
	vals = append(vals, make([]int, 9999999)...)

	opLengthMap := getOpLengthMap()

	nextOp := 0

	output := []int{}
	b := make(board)

	boadCreated := false
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

		out := performOp(idx, inputVal, &nextOp, vals, c, b, output, &boadCreated)
		if c.op == display {
			output = append(output, out)

			if _, ok := b[coord{x: -1, y: 0}]; !ok {
				b = createGameObjects(output)
			} else {
				b.updateBoard(createGameObjects(output))
			}
		}

		if c.op == saveAddr {
			clearConsole()
			fmt.Println(b.toString())
		}

	}
	clearConsole()
	fmt.Println(b.toString())

	return nil
}

func performOp(idx, inputVal int, nextOp *int, vals []int, c code, b board, output []int, boadCreated *bool) int {
	var out int
	p1, p2, p3 := getParams(vals, c, idx)
	switch c.op {
	case add:
		vals[p3] = p1 + p2
	case mult:
		vals[p3] = p1 * p2
	case saveAddr:
		*boadCreated = true
		ballCoord := coord{}
		paddleCoord := coord{}
		for k, v := range b {
			if v == horizontalPaddle {
				paddleCoord = k
			}
			if v == ball {
				ballCoord = k
			}
		}

		i := 0

		if ballCoord.x > paddleCoord.x {
			i = 1
		}
		if ballCoord.x < paddleCoord.x {
			i = -1
		}
		time.Sleep(1 * time.Millisecond)
		vals[p1] = i
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

func (b board) updateBoard(update board) {
	for k, u := range update {
		b[k] = u
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
