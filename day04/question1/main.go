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

		for idx, digit := range digitArr {
			if idx == 0 {
				continue
			}
			prev, err := strconv.Atoi(digitArr[idx-1])
			if err != nil {
				fmt.Println(err)
				continue
			}
			digit, err := strconv.Atoi(digit)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if prev == digit {
				req.double = true
			}

			if prev > digit {
				req.increasing = false
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
