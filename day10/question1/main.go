package main

import (
	"fmt"
	"strings"

	"github.com/johnny88/aoc2019/fileparse"
)

const (
	asteroid = "#"
	empty    = "."
)

type coordinate struct {
	x, y int
}

type lineOfSight struct {
	x, y int
}

type size struct {
	width, height int
}

type asteroidMap [][]string

func main() {
	scanner, err := fileparse.NewScanner("day10/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer scanner.Close()

	var m asteroidMap
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		m = append(m, row)
	}
	s := asteroidMapSize(m)
	fmt.Println(m, s.toString())
}

func lineOfSightArray(point coordinate) []lineOfSight {
	return []lineOfSight{}
}

func asteroidMapSize(m asteroidMap) size {
	return size{
		width:  len(m[0]),
		height: len(m),
	}
}

func (s *size) toString() string {
	return fmt.Sprintf("width: %d, height: %d", s.width, s.height)
}
