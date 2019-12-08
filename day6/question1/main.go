package main

import (
	"fmt"
	"strings"

	"github.com/johnny88/aoc2019/fileparse"
)

type link struct {
	parent, node             string
	indirectLink, directLink int
}

func main() {
	scanner, err := fileparse.NewScanner("day6/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer scanner.Close()

	a := map[string]link{}
	for scanner.Scan() {
		pair := strings.Split(scanner.Scanner.Text(), ")")
		a[pair[1]] = link{
			parent:     pair[0],
			node:       pair[1],
			directLink: 1,
		}
	}
	for k, l := range a {
		numLinks := getIndirectLink(a, a[k].parent)
		l.indirectLink = numLinks
		a[k] = l
	}

	var sum int
	for _, l := range a {
		sum += l.directLink + l.indirectLink
	}
	fmt.Println(sum)
}

func getIndirectLink(a map[string]link, key string) int {
	val, ok := a[key]
	if !ok {
		return 0
	}
	return 1 + getIndirectLink(a, val.parent)
}
