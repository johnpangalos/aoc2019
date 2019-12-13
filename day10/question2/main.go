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
	distance, direction int
	ratio               float64
}

type quadrantRatios map[int][]float64
type quadrantRatiosMap map[coord]quadrantRatios

type lineOfSightMap map[coord]lineOfSight

func main() {
	scanner, err := fileparse.NewScanner("day10/test1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer scanner.Close()

	m := map[coord]string{}
	coords := []coord{}
	rowCount := 0
	c := coord{x: 3, y: 4}
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		for idx, val := range row {
			c := coord{idx, rowCount}
			coords = append(coords, c)
			m[c] = val
		}
		rowCount++
	}

	ls := make(lineOfSightMap)
	qrm := make(quadrantRatiosMap)

	for k, v := range m {
		if v == "." {
			continue
		}

		qr := make(quadrantRatios)
		for _, c := range coords {
			if m[c] == "." || k == c {
				continue
			}
			r := ratio(k, c)
			d := dir(k, c)
			vec := vector{
				ratio:     r,
				direction: d,
				distance:  dist(k, c),
			}

			ls.addVector(k, vec)
			qr.addRatio(r, d)
		}
		qrm[k] = qr
	}
	fmt.Println(qrm.toString())
}

func (q quadrantRatios) addRatio(r float64, d int) {
	if _, ok := q[d]; !ok {
		q[d] = []float64{}
	}

	rs := q[d]
	for _, ratio := range rs {
		if ratio == r {
			return
		}
	}

	rs = append(rs, r)
	q[d] = rs
}

func (qrm quadrantRatiosMap) toString() string {
	var s []string
	for k, v := range qrm {
		s = append(s, fmt.Sprintf("%d: %v", k, v))
	}
	return strings.Join(s, "\n")
}

func (ls lineOfSightMap) addVector(c coord, v vector) {
	if _, oc := ls[c]; !oc {
		ls[c] = lineOfSight{
			vectors: []vector{v},
		}
	}

	l := ls[c]
	l.vectors = append(l.vectors, v)
	ls[c] = l
}

func ratio(p1, p2 coord) float64 {
	x := p1.x - p2.x
	y := p1.y - p2.y
	if y == 0 {
		return math.Inf(x)
	}
	return float64(x) / float64(y)
}

func dist(p1, p2 coord) int {
	x := p1.x - p2.x
	y := p1.y - p2.y
	return x + y
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
		return right
	}
	if y == 0 && x > 0 {
		return left
	}
	if x > 0 && y < 0 {
		return downLeft
	}
	if x > 0 && y > 0 {
		return upLeft
	}
	if x < 0 && y < 0 {
		return downRight
	}
	if x < 0 && y > 0 {
		return upRight
	}
	return -1
}

func (c *coord) toString() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}
