package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/johnny88/aoc2019/fileparse"
)

type point struct {
	x, y, count1, count2 int64
	wire1, wire2         bool
}

func main() {
	scanner, err := fileparse.NewScanner("day3/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer scanner.Close()

	first := scanner.CommaStringParse()
	second := scanner.CommaStringParse()
	m := map[string]point{"x0y0": point{x: 0, y: 0, wire1: false, wire2: false}}

	curr := "x0y0"
	count := int64(0)
	for _, op := range first {
		curr, count = logic(op, m, curr, 1, count)
	}

	curr = "x0y0"
	count = int64(0)
	for _, op := range second {
		curr, count = logic(op, m, curr, 2, count)
	}
	min := int64(10000000000)

	for _, val := range m {
		num := val.count1 + val.count2
		if val.wire1 && val.wire2 && (min > num) {
			min = num
		}
	}
	fmt.Println(min)
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func createKey(x int64, y int64) string {
	return strings.Join([]string{
		"x", strconv.FormatInt(x, 10), "y", strconv.FormatInt(y, 10),
	}, "")
}

func logic(op string, m map[string]point, curr string, idx int64, count int64) (string, int64) {
	opArr := strings.SplitN(op, "", 2)

	val, err := strconv.ParseInt(opArr[1], 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	key := curr

	for i := int64(0); i < val; i++ {
		var xval, yval int64
		count1 := m[key].count1
		count2 := m[key].count2
		switch opArr[0] {
		case "R":
			xval = m[key].x + int64(1)
			yval = m[key].y
			key = createKey(xval, m[key].y)
		case "L":
			xval = m[key].x - int64(1)
			yval = m[key].y
			key = createKey(xval, m[key].y)
		case "U":
			xval = m[key].x
			yval = m[key].y + int64(1)
			key = createKey(m[key].x, yval)
		case "D":
			xval = m[key].x
			yval = m[key].y - int64(1)
			key = createKey(m[key].x, yval)
		}
		wire1 := idx == int64(1)
		wire2 := idx == int64(2)
		if idx == int64(1) {
			count1++
		}
		if idx == int64(2) {
			count2++
		}
		if val, ok := m[key]; ok {
			if idx == int64(1) {
				wire1 = true
				wire2 = val.wire2
			}
			if idx == int64(2) {
				wire1 = val.wire1
				wire2 = true
				count1 = val.count1
			}
		}
		m[key] = point{x: xval, y: yval, wire1: wire1, wire2: wire2, count1: count1, count2: count2}
	}
	return key, val
}
