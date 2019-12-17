package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/johnny88/aoc2019/fileparse"
)

const (
	ore  = "ORE"
	fuel = "FUEL"
)

type chemical struct {
	name   string
	amount float64
}

type chemicals []chemical
type equation struct {
	params chemicals
	result chemical
}
type equationMap map[string]equation
type amountMap map[string]float64

func main() {
	scanner, err := fileparse.NewScanner("day14/test4.txt")
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

		eqMap[key.name] = equation{params: cArr, result: key}
	}

	toDo := []chemical{eqMap[fuel].result}

	requirements := []amountMap{}
	count := 0
	for len(toDo) > 0 {
		required := make(amountMap)
		for _, c := range toDo {
			required = join(required, applyEquation(eqMap, c.name, c.amount))
		}
		toDo = []chemical{}
		fmt.Println(required)
		for k := range required {
			if k == ore {
				continue
			}
			for i := count - 1; i >= 1; i-- {
				if val, ok := requirements[i][k]; ok {
					toRemove := applyEquation(eqMap, k, val)
					fmt.Println(k, toRemove)
					required[k] += requirements[i][k]
					requirements[i][k] = 0

					for k, v := range toRemove {
						required[k] -= v
					}
				}
			}
			toDo = append(toDo, chemical{name: k, amount: required[k]})
		}
		requirements = append(requirements, required)
		count++
	}

	fmt.Println(requirements)
	sum := float64(0)
	for _, v := range requirements {
		for k, r := range v {
			if k == ore {
				sum += r
			}
		}
	}
	fmt.Println(int(sum))
}

func newChemical(input string) (chemical, error) {
	arr := strings.Split(strings.TrimSpace(input), " ")
	amount, err := strconv.ParseFloat(arr[0], 64)
	if err != nil {
		return chemical{}, err
	}
	return chemical{name: arr[1], amount: amount}, nil
}

func applyEquation(e equationMap, key string, amount float64) amountMap {
	amounts := make(amountMap)
	eq := e[key]
	for _, chem := range eq.params {
		mult := math.Ceil(amount / eq.result.amount)
		amounts[chem.name] = mult * chem.amount
	}
	return amounts
}

func join(m1, m2 amountMap) amountMap {
	for k, v := range m1 {
		m2[k] += v
	}
	return m2
}
