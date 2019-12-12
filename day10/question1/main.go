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
)

type coord struct {
	x, y int
}

type lineOfSight struct {
	ratios []float64
}

// type size struct {
// width, height int
// }

func main() {
	scanner, err := fileparse.NewScanner("day10/test.txt")
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
			c := coord{rowCount, idx}
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
			if _, ok := ls[k]; ok {
				ls[k] = lineOfSight{
					ratios: []float64{ratio(k, c)},
				}
			}
		}
	}
	fmt.Println(ls)
	// s := asteroidMapSize(m)

}

func ratio(p1, p2 coord) float64 {
	x := p1.x - p2.x
	y := p1.y - p2.y
	if y == 0 {
		return math.Inf(x)
	}
	return float64(x) / float64(y)
}

// func asteroidMapSize(m asteroidMap) size {
// return size{
// width:  len(m[0]),
// height: len(m),
// }
// }

// func (s *size) toString() string {
// return fmt.Sprintf("width: %d, height: %d", s.width, s.height)
// }
