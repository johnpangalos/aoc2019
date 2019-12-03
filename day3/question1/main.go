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
	scanner, err := fileparse.NewScanner("day3/test1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer scanner.Close()

	first := scanner.CommaStringParse()
	// second := scanner.CommaStringParse()
	m := map[string]point{"x0y0": point{x: 0, y: 0, count: 0}}
	curr := "x0y0"
	for _, op := range first {
		opArr := strings.SplitN(op, "", 2)
		val, err := strconv.ParseInt(opArr[1], 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		var target point

		switch opArr[0] {
		case "R":
			target.x = m[curr].x - val
		case "L":
			target.x = m[curr].x - val
		case "U":
			target.y = m[curr].y + val
		case "D":
			target.y = m[curr].y - val
		}

		for x := m[curr].x; x < target.x; x++ {
			key := strings.Join([]string{"x", m[curr].x.(string) + int64(1), "y", m[curr].y}, "")
			count := int64(0)
			if val, ok := m[key]; ok {
				count = val.count + int64(1)
			}
			m[key] = point{x: m[curr].x, y: m[curr].y, count: count}
		}

		for y := m[curr].y; y < target.y; y++ {

		}

	}
}
