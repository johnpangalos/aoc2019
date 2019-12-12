package main

import (
	"fmt"
	"strings"

	"github.com/johnny88/aoc2019/fileparse"
)

type link struct {
	parent, node string
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
			parent: pair[0],
			node:   pair[1],
		}
	}
	c := getCommonParent(a, "SAN", "YOU")
	fmt.Println(
		getDistanceToParent(a, c, a["SAN"]) + getDistanceToParent(a, c, a["YOU"]) - 2,
	)
}

func getParents(a map[string]link, key string) []string {
	val, ok := a[key]
	if !ok {
		return []string{}
	}
	return append(getParents(a, val.parent), val.parent)
}

func getCommonParent(a map[string]link, key1, key2 string) string {
	parents1 := getParents(a, key1)
	parents2 := getParents(a, key2)
	var common string
	for idx, p := range parents1 {
		if p != parents2[idx] {
			break
		}
		common = p
	}
	return common
}

func getDistanceToParent(a map[string]link, target string, l link) int {
	if l.node == target {
		return 0
	}
	if l.node == "COM" {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("Reached COM"))
	}
	return getDistanceToParent(a, target, a[l.parent]) + 1
}
