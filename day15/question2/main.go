package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/johnny88/aoc2019/fileparse"

	"github.com/fatih/color"
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

	north = 1
	south = 2
	west  = 3
	east  = 4

	wall   = "wall"
	empty  = "empty"
	start  = "start"
	tank   = "tank"
	oxygen = "oxygen"

	foundWall = 0
	moveDone  = 1
	foundTank = 2
)

type code struct {
	op, firstMode, secondMode, thirdMode int
}

type coord struct {
	x, y int
}

type areaMap map[coord]string

var relativeBase = 0

func main() {
	scanner, err := fileparse.NewScanner("day15/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer scanner.Close()
	vals := scanner.CommaStringParseInt()
	runIntcode(vals)
}

func runIntcode(vals []int) {
	vals = append(vals, make([]int, 9999999)...)
	opLengthMap := getOpLengthMap()
	nextOp := 0

	currCoord := coord{x: 0, y: 0}
	aMap := make(areaMap)
	aMap[currCoord] = start
	direction := north

	for {
		idx := nextOp
		c := code{op: vals[nextOp]}
		if vals[nextOp] > 4 {
			c = parseOpCode(vals[nextOp])
		}
		nextOp = nextOp + opLengthMap[c.op]
		out := performOp(direction, idx, &nextOp, vals, c)

		if c.op == display {

			nCoord := nextCoord(direction, currCoord)

			switch out {
			case foundWall:
				if _, ok := aMap[nCoord]; !ok {
					aMap[nCoord] = wall
				}
				direction = nextDirectionCouterClockwise(direction)
			case moveDone:
				if _, ok := aMap[nCoord]; !ok {
					aMap[nCoord] = empty
				}
				direction = nextDirectionClockwise(direction)
				currCoord = nCoord
			case foundTank:
				if _, ok := aMap[nCoord]; !ok {
					aMap[nCoord] = tank
				}
				direction = nextDirectionClockwise(direction)
				currCoord = nCoord
			}

			origin := coord{x: 0, y: 0}
			aMap.printToConsole(currCoord, 1)
			if nCoord == origin {
				break
			}
		}
	}
	var t coord
	for k, v := range aMap {
		if v == tank {
			t = k
		}
	}
	oxygenPaths := []coord{t}
	// don't count the first one...
	count := -1
	tmp := coord{x: 99999999, y: 99999999}
	for len(oxygenPaths) > 0 {
		newOxPaths := []coord{}
		for _, v := range oxygenPaths {
			newOxPaths = append(newOxPaths, emptyAdjactent(aMap, v)...)
			aMap[v] = oxygen
		}
		oxygenPaths = newOxPaths
		count++
		aMap.printToConsole(tmp, 5)
	}
}

func emptyAdjactent(a areaMap, c coord) []coord {
	toCheck := []coord{
		coord{
			x: c.x - 1,
			y: c.y,
		},
		coord{
			x: c.x + 1,
			y: c.y,
		},
		coord{
			x: c.x,
			y: c.y - 1,
		},
		coord{
			x: c.x,
			y: c.y + 1,
		},
	}
	adj := []coord{}
	for _, v := range toCheck {
		if a[v] != empty && a[v] != start {
			continue
		}
		adj = append(adj, v)
	}
	return adj
}

func clearConsole() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func nextDirectionClockwise(direction int) int {
	var nextDirection int
	switch direction {
	case north:
		nextDirection = east
	case south:
		nextDirection = west
	case east:
		nextDirection = south
	case west:
		nextDirection = north
	}

	return nextDirection
}

func nextDirectionCouterClockwise(direction int) int {
	var nextDirection int
	switch direction {
	case north:
		nextDirection = west
	case south:
		nextDirection = east
	case east:
		nextDirection = north
	case west:
		nextDirection = south
	}

	return nextDirection
}

func nextCoord(direction int, c coord) coord {
	newCoord := coord{}
	switch direction {
	case north:
		newCoord = coord{x: c.x, y: c.y + 1}
	case south:
		newCoord = coord{x: c.x, y: c.y - 1}
	case east:
		newCoord = coord{x: c.x + 1, y: c.y}
	case west:
		newCoord = coord{x: c.x - 1, y: c.y}
	}
	return newCoord
}

func performOp(inputVal, idx int, nextOp *int, vals []int, c code) int {
	p1, p2, p3 := getParams(vals, c, idx)
	out := -1
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

func (a areaMap) toString(c coord) string {
	minX := -21
	minY := -19
	maxX := 19
	maxY := 21
	var s string
	// for k := range a {
	// if k.x < minX {
	// minX = k.x
	// }
	// if k.y < minY {
	// minY = k.y
	// }
	// if k.x > maxX {
	// maxX = k.x
	// }
	// if k.y > maxY {
	// maxY = k.y
	// }
	// }
	for j := minY; j < maxY+1; j++ {
		for i := minX; i < maxX+1; i++ {
			c2 := coord{x: i, y: j}
			v := a[c2]
			if c == c2 {
				red := color.New(color.FgRed).SprintFunc()
				s = strings.Join([]string{s, fmt.Sprintf("%s", red("◉"))}, "")
				continue
			}
			switch v {
			case wall:
				s = strings.Join([]string{s, "█"}, "")
			case oxygen:
				blue := color.New(color.FgBlue).SprintFunc()
				s = strings.Join([]string{s, fmt.Sprintf("%s", blue("█"))}, "")
			case empty:
				s = strings.Join([]string{s, "·"}, "")
			case tank:
				s = strings.Join([]string{s, "T"}, "")
			case start:
				s = strings.Join([]string{s, "E"}, "")
			default:
				s = strings.Join([]string{s, "█"}, "")
			}
		}
		s = strings.Join([]string{s, "\n"}, "")
	}
	return s
}

func (a areaMap) printToConsole(c coord, ms int) {
	// clearConsole()
	fmt.Println(a.toString(c))
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

func (a areaMap) emptySpace() {
	totalSpace := 0
	for _, v := range a {
		if v == empty {
			totalSpace++
		}
	}
}
