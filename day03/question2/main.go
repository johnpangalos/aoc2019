package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/johnny88/aoc2019/fileparse"
)

type point struct {
	x, y, count1, count2 int
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

	moveSet(first, m, 1)
	moveSet(second, m, 2)

	min := 10000000000

	for _, val := range m {
		num := val.count1 + val.count2
		if val.wire1 && val.wire2 && (min > num) {
			min = num
		}
	}
	fmt.Println(min)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func createKey(x int, y int) string {
	return strings.Join([]string{
		"x", strconv.Itoa(x), "y", strconv.Itoa(y),
	}, "")
}
func moveSet(set []string, m map[string]point, idx int) {
	curr := "x0y0"
	count := 0
	for _, op := range set {
		curr, count = move(op, m, curr, idx, count)
	}
}

func move(op string, m map[string]point, curr string, idx int, count int) (string, int) {
	opArr := strings.SplitN(op, "", 2)

	val, err := strconv.Atoi(opArr[1])
	if err != nil {
		fmt.Println(err)
	}

	key := curr

	for i := 0; i < val; i++ {
		var xval, yval int
		count1 := m[key].count1
		count2 := m[key].count2
		switch opArr[0] {
		case "R":
			xval = m[key].x + 1
			yval = m[key].y
			key = createKey(xval, m[key].y)
		case "L":
			xval = m[key].x - 1
			yval = m[key].y
			key = createKey(xval, m[key].y)
		case "U":
			xval = m[key].x
			yval = m[key].y + 1
			key = createKey(m[key].x, yval)
		case "D":
			xval = m[key].x
			yval = m[key].y - 1
			key = createKey(m[key].x, yval)
		}
		wireCount(idx, &count1, &count2, m, key)
		wire1, wire2 := wireState(idx, m, key)
		m[key] = point{x: xval, y: yval, wire1: wire1, wire2: wire2, count1: count1, count2: count2}
	}
	return key, val
}

func wireState(idx int, m map[string]point, key string) (bool, bool) {
	wire1 := idx == 1
	wire2 := idx == 2

	if val, ok := m[key]; ok {
		if idx == 1 {
			wire1 = true
			wire2 = val.wire2
		}
		if idx == 2 {
			wire1 = val.wire1
			wire2 = true
		}
	}
	return wire1, wire2
}

func wireCount(idx int, count1 *int, count2 *int, m map[string]point, key string) {
	if idx == 1 {
		*count1++
	}
	if idx == 2 {
		*count2++
	}
	if val, ok := m[key]; ok {
		if idx == 2 {
			*count1 = val.count1
		}
	}
}
