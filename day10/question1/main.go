package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/johnny88/aoc2019/fileparse"
)

const (
	asteroid = "#"
	empty    = "."

	left      = 0
	right     = 1
	up        = 2
	down      = 3
	downLeft  = 4
	downRight = 5
	upLeft    = 6
	upRight   = 7
)

type coord struct {
	x, y int
}

type lineOfSight struct {
	vectors []vector
}

type vector struct {
	ratio     float64
	direction int
}

func main() {
	scanner, err := fileparse.NewScanner("day10/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer scanner.Close()

	m := map[coord]string{}
	coords := []coord{}
	rowCount := 0

	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		for idx, val := range row {
			c := coord{idx, rowCount}
			coords = append(coords, c)
			m[c] = val
		}
		rowCount++
	}

	ls := map[coord]lineOfSight{}

	for k, v := range m {
		if v == "." {
			continue
		}
		for _, c := range coords {
			if m[c] == "." || k == c {
				continue
			}
			vec := vector{
				ratio:     ratio(k, c),
				direction: dir(k, c),
			}
			if _, ok := ls[k]; !ok {
				ls[k] = lineOfSight{
					vectors: []vector{vec},
				}
			}

			l := ls[k]
			hasVector := false

			for _, v := range l.vectors {
				if v == vec {
					hasVector = true
					break
				}
			}
			if hasVector {
				continue
			}
			l.vectors = append(l.vectors, vec)
			ls[k] = l
		}
	}

	max := 0
	var c coord
	for k, v := range ls {
		if max < len(v.vectors) {
			max = len(v.vectors)
			c = k
		}
	}
	fmt.Println(c.x, c.y, max)
}

func ratio(p1, p2 coord) float64 {
	x := p1.x - p2.x
	y := p1.y - p2.y
	if y == 0 {
		return math.Inf(x)
	}
	return float64(x) / float64(y)
}

func dir(p1, p2 coord) int {
	x := p1.x - p2.x
	y := p1.y - p2.y
	if x == 0 && y < 0 {
		return down
	}
	if x == 0 && y > 0 {
		return up
	}
	if y == 0 && x < 0 {
		return left
	}
	if y == 0 && x > 0 {
		return right
	}
	if x > 0 && y < 0 {
		return downRight
	}
	if x > 0 && y > 0 {
		return upRight
	}
	if x < 0 && y < 0 {
		return downLeft
	}
	if x < 0 && y > 0 {
		return upLeft
	}
	return -1
}

func (c *coord) toString() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}
