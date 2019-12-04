package main

import (
	"fmt"
	"strconv"
	"strings"
)

const min = 146810
const max = 612564

type requirements struct {
	double, increasing bool
}

func main() {
	passwordMap := map[string]requirements{}

	for i := min; i <= max; i++ {
		password := strconv.Itoa(i)
		digitArr := strings.Split(password, "")

		req := requirements{
			double:     false,
			increasing: true,
		}
		digitMap := baseMap()

		for idx, digit := range digitArr {
			digit, err := strconv.Atoi(digit)
			if err != nil {
				fmt.Println(err)
				continue
			}
			digitMap[digit]++

			if idx == 0 {
				continue
			}

			prev, err := strconv.Atoi(digitArr[idx-1])
			if err != nil {
				fmt.Println(err)
				continue
			}

			if prev > digit {
				req.increasing = false
			}
		}

		for _, count := range digitMap {
			if count == 2 {
				req.double = true
			}
		}

		passwordMap[password] = req
	}

	count := 0
	for _, p := range passwordMap {
		if p.double && p.increasing {
			count++
		}
	}
	fmt.Println(count)
}

func baseMap() map[int]int {
	return map[int]int{
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
		7: 0,
		8: 0,
		9: 0,
	}

}
