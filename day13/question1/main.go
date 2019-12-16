package main

import (
	"fmt"

	"github.com/johnny88/aoc2019/day09/intcode"
	"github.com/johnny88/aoc2019/fileparse"
)

const (
	empty            = 0
	wall             = 1
	block            = 2
	horizontalPaddle = 3
	ball             = 4
)

type gameObject struct {
	x, y, objectType int
}

func main() {
	scanner, err := fileparse.NewScanner("day13/input.txt")
	if err != nil {
		panic(err)
	}
	defer scanner.Close()

	vals := scanner.CommaStringParseInt()
	o, _, err := intcode.RunIntcode(0, vals)
	if err != nil {
		panic(err)
	}

	currObj := gameObject{}
	gameObjects := []gameObject{}
	for idx, v := range o {
		if (idx)%3 == 0 {
			currObj.x = v
		}
		if (idx)%3 == 1 {
			currObj.y = v
		}
		if (idx)%3 == 2 {
			currObj.objectType = v
			gameObjects = append(gameObjects, currObj)
			currObj = gameObject{}
		}
	}

	counter := 0
	for _, obj := range gameObjects {
		if obj.objectType == block {
			counter++
		}
	}
	fmt.Println(counter)
}
