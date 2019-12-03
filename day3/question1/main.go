package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/johnny88/aoc2019/fileparse"
)

type point struct {
	x, y, count int64
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
	m := map[string]point{"x0y0": point{x: 0, y: 0, count: 0}}
	fmt.Println(first, second)
	curr := "x0y0"
	for _, op := range first {
		curr = logic(op, m, curr)
	}

	curr = "x0y0"
	for _, op := range second {
		curr = logic(op, m, curr)
	}
	min := int64(10000000000)

	for _, val := range m {
		num := abs(val.x) + abs(val.y)
		if val.count > 1 {
			fmt.Println(val, num)
		}
		if val.count > 1 && (min > num) {
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
	return strings.Join([]string{"x", strconv.FormatInt(x, 10), "y", strconv.FormatInt(y, 10)}, "")
}

func logic(op string, m map[string]point, curr string) string {
	opArr := strings.SplitN(op, "", 2)

	val, err := strconv.ParseInt(opArr[1], 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	var key string

	switch opArr[0] {
	case "R":
		target := point{
			x: m[curr].x + val,
			y: m[curr].y,
		}

		for x := m[curr].x; x < target.x; x++ {
			xval := m[curr].x + int64(1)
			key = createKey(xval, m[curr].y)
			count := int64(1)
			if val, ok := m[key]; ok {
				count = val.count + int64(1)
			}
			m[key] = point{x: xval, y: m[curr].y, count: count}
			curr = key
		}
	case "L":
		target := point{
			x: m[curr].x - val,
			y: m[curr].y,
		}

		for x := m[curr].x; x > target.x; x-- {
			xval := m[curr].x - int64(1)
			key = createKey(xval, m[curr].y)
			count := int64(1)
			if val, ok := m[key]; ok {
				count = val.count + int64(1)
			}
			m[key] = point{x: xval, y: m[curr].y, count: count}
			curr = key
		}
	case "U":
		target := point{
			y: m[curr].y + val,
			x: m[curr].x,
		}

		for y := m[curr].y; y < target.y; y++ {
			yval := m[curr].y + int64(1)
			key = createKey(m[curr].x, yval)

			count := int64(1)
			if val, ok := m[key]; ok {
				count = val.count + int64(1)
			}
			m[key] = point{x: m[curr].x, y: yval, count: count}
			curr = key
		}
	case "D":
		target := point{
			y: m[curr].y - val,
			x: m[curr].x,
		}
		for y := m[curr].y; y > target.y; y-- {
			yval := m[curr].y - int64(1)
			key = createKey(m[curr].x, yval)

			count := int64(1)
			if val, ok := m[key]; ok {
				count = val.count + int64(1)
			}
			m[key] = point{x: m[curr].x, y: yval, count: count}
			curr = key
		}
	}
	return key
}
