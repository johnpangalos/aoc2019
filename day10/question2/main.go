package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/johnny88/aoc2019/fileparse"
)

const (
	asteroid = "#"
	empty    = "."

	left      = "left"
	right     = "right"
	up        = "up"
	down      = "down"
	downLeft  = "downLeft"
	downRight = "downRight"
	upLeft    = "upLeft"
	upRight   = "upRight"
)

type coord struct {
	x, y int
}

type lineOfSight struct {
	vectors []vector
}

type vector struct {
	distance  int
	direction string
	ratio     float64
	coord     coord
}

type orderItem struct {
	ratio     float64
	direction string
}

type quadrantRatios map[string][]float64
type order []orderItem

type vectorArr []vector
type vectorMap map[float64]vectorArr

func main() {
	scanner, err := fileparse.NewScanner("day10/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer scanner.Close()

	clockwise := []string{
		up, upRight, right, downRight, down, downLeft, left, upLeft,
	}

	m := map[coord]string{}
	coords := []coord{}
	rowCount := 0

	// Use answer from part one
	laser := coord{x: 37, y: 25}

	numAsteroids := 0
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		for idx, val := range row {
			c := coord{idx, rowCount}
			if val == "#" && c != laser {
				numAsteroids++
			}
			coords = append(coords, c)
			m[c] = val
		}
		rowCount++
	}

	qr := make(quadrantRatios)
	vm := make(vectorMap)

	for _, c := range coords {
		if m[c] == "." || laser == c {
			continue
		}
		r := ratio(laser, c)
		d := dir(laser, c)
		vec := vector{
			ratio:     r,
			direction: d,
			distance:  dist(laser, c),
			coord:     c,
		}

		vm.addVector(r, vec)
		qr.addRatio(r, d)
	}

	for k, rs := range qr {
		switch k {
		case upRight, downRight, downLeft, upLeft:
			sort.Sort(sort.Reverse(sort.Float64Slice(rs)))
		}
	}

	destroyOrder := []coord{}
	for len(destroyOrder) < numAsteroids {
		for _, currDir := range clockwise {
			for _, r := range qr[currDir] {
				var vs []vector
				if len(vm[r]) == 0 {
					continue
				}

				for _, v := range vm[r] {
					if v.direction == currDir {
						vs = append(vs, v)
					}
				}

				min := 9999999999999
				var minVec vector
				for _, v := range vs {
					if v.distance < min {
						min = v.distance
						minVec = v
					}
				}

				destroyOrder = append(destroyOrder, minVec.coord)
				vm[r] = remove(vm[r], minVec)
			}
		}
	}
	fmt.Println(destroyOrder[199].x*100 + destroyOrder[199].y)
}

func (q quadrantRatios) addRatio(r float64, d string) {
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

func (vm vectorMap) addVector(c float64, v vector) {
	if _, oc := vm[c]; !oc {
		vm[c] = []vector{}
	}

	vs := vm[c]
	vs = append(vs, v)
	vm[c] = vs
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
	x := math.Abs(float64(p1.x) - float64(p2.x))
	y := math.Abs(float64(p1.y) - float64(p2.y))
	return int(x + y)
}

func dir(p1, p2 coord) string {
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
	return ""
}

func (c *coord) toString() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}

func (vm vectorMap) toString() string {
	var s []string
	for k, v := range vm {
		s = append(s, fmt.Sprintf("%f: %v", k, v.toString()))
	}
	return strings.Join(s, "\n")
}

func (vm vectorMap) length() int {
	counter := 0
	for _, v := range vm {
		counter += len(v)
	}
	return counter
}

func (q quadrantRatios) toString() string {
	var s []string
	for k, v := range q {
		s = append(s, fmt.Sprintf("%s: %f", k, v))
	}
	return strings.Join(s, "\n")
}

func (v *vector) toString() string {
	return fmt.Sprintf(
		"direction: %s, ratio: %f, distance: %d, coord: %d,%d",
		v.direction,
		v.ratio,
		v.distance,
		v.coord.x,
		v.coord.y,
	)
}

func (vs vectorArr) toString() string {
	s := []string{"["}
	for _, v := range vs {
		s = append(s, fmt.Sprintf("  %s", v.toString()))
	}
	s = append(s, "]")
	return strings.Join(s, "\n")

}

func remove(s []vector, v vector) []vector {
	i := 0
	for idx, vec := range s {
		if vec == v {
			i = idx
			break
		}
	}

	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func removeCoord(s []coord, c coord) []coord {
	i := 0
	for idx, co := range s {
		if c == co {
			i = idx
			break
		}
	}

	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
