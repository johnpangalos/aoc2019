package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/johnny88/aoc2019/fileparse"
)

const (
	wall     = '#'
	empty    = '.'
	entrance = '@'
)

type coord struct {
	x, y int
}

type dungeonMap map[coord]rune
type posMap map[rune]coord

func main() {
	scanner, err := fileparse.NewScanner("day18/test1.txt")
	if err != nil {
		panic(err)
	}
	defer scanner.Close()

	dMap := make(dungeonMap)
	keyMap := make(posMap)
	doorMap := make(posMap)
	y := 0
	entr := coord{}
	for scanner.Scan() {
		t := scanner.Text()
		for x, r := range t {
			c := coord{x: x, y: y}
			dMap[c] = r
			if isLetter(r) && unicode.IsLower(r) {
				keyMap[r] = c
			} else if isLetter(r) {
				doorMap[r] = c
			}
			if r == entrance {
				entr = coord{x: x, y: y}
			}
		}
		y++
	}
	fmt.Println(keyMap.toString())
	fmt.Println(doorMap.toString())
	fmt.Println(entr.toString())
}

func isLetter(r rune) bool {
	return r != wall && r != empty && r != entrance
}

func (k posMap) toString() string {
	s := []string{}
	for r, c := range k {
		s = append(s, fmt.Sprintf("%c => %s", r, c.toString()))
	}
	return strings.Join(s, "\n")
}

func (c *coord) toString() string {
	return fmt.Sprintf("x: %d, y: %d", c.x, c.y)
}
