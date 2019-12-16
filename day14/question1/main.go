package main

import (
	"strconv"
	"strings"

	"github.com/johnny88/aoc2019/fileparse"
)

type chemical struct {
	name   string
	amount int
}

type chemicals []chemical
type equation struct {
	params chemicals
	result chemical
}
type equationMap map[string][]equation

func main() {
	scanner, err := fileparse.NewScanner("day14/test2.txt")
	if err != nil {
		panic(err)
	}
	defer scanner.Close()

	eqMap := make(equationMap)
	for scanner.Scan() {
		eq := scanner.Text()
		eqArr := strings.Split(eq, "=>")

		key, err := newChemical(eqArr[1])
		if err != nil {
			panic(err)
		}

		cArr := chemicals{}
		for _, s := range strings.Split(eqArr[0], ",") {
			c, err := newChemical(s)
			if err != nil {
				panic(err)
			}
			cArr = append(cArr, c)
		}

		eqMap[key.name] = []equation{equation{params: cArr, result: key}}
	}
	m := map[string]int{}
	leftovers := map[string]int{}
	calcOre(eqMap, "FUEL", 1, m, leftovers)

	for hasChemicals(m) {
		currM := map[string]int{}
		for key := range m {
			calcOre(eqMap, key, m[key], currM, leftovers)
		}
		m = currM
	}
}

func hasChemicals(m map[string]int) bool {
	hasChemicals := false
	for key := range m {
		if key != "ORE" {
			hasChemicals = true
			break
		}
	}
	return hasChemicals
}

func newChemical(input string) (chemical, error) {
	arr := strings.Split(strings.TrimSpace(input), " ")
	amount, err := strconv.Atoi(arr[0])
	if err != nil {
		return chemical{}, err
	}
	return chemical{name: arr[1], amount: amount}, nil
}

// func (c *chemical) toString() string {
// return fmt.Sprintf("name: %s, amount: %d", c.name, c.amount)
// }

func calcOre(e equationMap, key string, amount int, m, leftovers map[string]int) {
	for _, eq := range e[key] {
		for _, p := range eq.params {
			m[p.name] += (amount - (amount % eq.result.amount)) / eq.result.amount * p.amount
			leftovers[key] += amount % eq.result.amount
		}
	}
}
