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
	scanner, err := fileparse.NewScanner("day14/input.txt")
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
	trillion := float64(1000000000000)

	oreCount := float64(0)
	count := float64(1000)
	for oreCount < trillion {
		if trillion-oreCount > 1000000000 {
			count += 1000
		} else {
			count++
		}
		oreCount = logic(eqMap, count)
	}
	// this is one more for some reason, gotta come back to this
	// for now making count - 1
	fmt.Printf("%f\n", count-1)
}

func logic(eqMap equationMap, amount float64) float64 {
	count := 0
	requirements := []amountMap{}
	toDo := []chemical{eqMap[fuel].result}
	for len(toDo) > 0 {
		required := make(amountMap)
		for _, c := range toDo {
			if c.name == fuel {
				required = join(required, applyEquation(eqMap, c.name, amount))
			} else {
				required = join(required, applyEquation(eqMap, c.name, c.amount))
			}
		}
		toDo = []chemical{}
		for k := range required {

			for i := count - 1; i >= 1; i-- {
				if val, ok := requirements[i][k]; ok {
					toRemove := applyEquation(eqMap, k, val)
					required[k] += requirements[i][k]
					requirements[i][k] = 0

					for k, v := range toRemove {
						required[k] -= v
					}
				}
			}
		}
		for k := range required {
			if k == ore {
				continue
			}
			toDo = append(toDo, chemical{name: k, amount: required[k]})
		}
		requirements = append(requirements, required)
		count++
	}

	sum := float64(0)
	for _, v := range requirements {
		for k, r := range v {
			if k == ore {
				sum += r
			}
		}
	}
	return sum
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
