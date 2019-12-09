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

	positionMode  = 0
	immediateMode = 1
)

type code struct {
	op         int
	firstMode  int
	secondMode int
	thirdMode  int
}

type state struct {
	idx int
	arr []int
}

func main() {
	settings := generateThrusterSettings()

	max := 0
	file := parseFile()

	for _, sArr := range settings {
		input := 0
		ampState := []state{
			state{idx: 0, arr: copySlice(file)},
			state{idx: 0, arr: copySlice(file)},
			state{idx: 0, arr: copySlice(file)},
			state{idx: 0, arr: copySlice(file)},
			state{idx: 0, arr: copySlice(file)},
		}

		fullStop := false
		count := 0
		for !fullStop {
			currAmp := count % 5
			s := sArr[currAmp]
			sInt, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			input = amplifier(input, sInt, ampState, currAmp, &fullStop)

			if input > max {
				max = input
			}
			count++
		}
	}
	fmt.Println(max)
}

func generateThrusterSettings() [][]string {
	arr := [][]string{}
	for i := 56789; i <= 98765; i++ {
		strArr := strings.Split(strconv.Itoa(i), "")

		hasUniqueDigits := true
		hasDigitUnderFour := false

		digitMap := map[string]int{}

		for _, str := range strArr {
			val, err := strconv.Atoi(str)
			if err != nil {
				panic(err)
			}
			if val <= 4 {
				hasDigitUnderFour = true
				continue
			}

			if _, ok := digitMap[str]; ok {
				hasUniqueDigits = false
				continue
			}

			digitMap[str] = 1
		}

		if !hasUniqueDigits || hasDigitUnderFour {
			continue
		}

		arr = append(arr, strArr)
	}

	return arr
}

func parseFile() []int {
	scanner, err := fileparse.NewScanner("day7/input.txt")
	if err != nil {
		panic(err)
	}
	defer scanner.Close()

	return scanner.CommaStringParseInt()
}

func copySlice(s []int) []int {
	out := make([]int, len(s))
	copy(out, s)
	return out
}

func amplifier(input, setting int, s []state, currAmp int, fullstop *bool) int {
	opLengthMap := getOpLengthMap()

	nextOp := s[currAmp].idx
	output := 0
	vals := s[currAmp].arr

	for {
		op := vals[nextOp]
		idx := nextOp
		if op == halt {
			if currAmp == 4 {
				*fullstop = true
			}
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
		} else {
			if c.op == saveAddr || c.op == display {
				p1, p2 = vals[idx+1], 0
			} else {
				p1, p2 = vals[idx+1], vals[idx+2]
			}
		}

		switch op {
		case add:
			if !longCode {
				p1, p2 = vals[p1], vals[p2]
			}
			addFunc(vals, p1, p2, vals[idx+3])
		case mult:
			if !longCode {
				p1, p2 = vals[p1], vals[p2]
			}
			multFunc(vals, p1, p2, vals[idx+3])
		case saveAddr:
			inputVal := input
			if idx == 0 {
				inputVal = setting
			}
			saveToRegister(vals, p1, inputVal)
		case display:
			output = vals[p1]
			s[currAmp].idx = nextOp
		case jumpIfTrue:
			if p1 > 0 {
				nextOp = p2
				s[currAmp].idx = p2
			}
		case jumpIfFalse:
			if p1 == 0 {
				nextOp = p2
				s[currAmp].idx = p2
			}
		case lessThan:
			if p1 < p2 {
				saveToRegister(vals, vals[idx+3], 1)
			} else {
				saveToRegister(vals, vals[idx+3], 0)
			}
		case equals:
			if p1 == p2 {
				saveToRegister(vals, vals[idx+3], 1)
			} else {
				saveToRegister(vals, vals[idx+3], 0)
			}
		}
		if op == display {
			break
		}
	}
	return output
}

func runOperation() {
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

func getVals(arr []int, c code, longCode bool, idx int) (int, int) {
	var p1, p2 int

	if longCode {
		p1, p2 = getParams(arr, c, idx)
	} else {
		if c.op == saveAddr || c.op == display {
			p1, p2 = arr[idx+1], 0
		} else {
			p1, p2 = arr[idx+1], arr[idx+2]
		}
	}
	return p1, p2
}
