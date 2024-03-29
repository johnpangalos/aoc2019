package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/johnny88/aoc2019/fileparse"
)

const (
	iter = 1000
)

type point struct {
	x, y, z int
}

type moon struct {
	position, velocity point
}

type moons []moon

func main() {
	scanner, err := fileparse.NewScanner("day12/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer scanner.Close()

	mArr := moons{}
	for scanner.Scan() {
		p, err := newPoint(scanner.Text())
		if err != nil {
			panic(err)
		}
		mArr = append(mArr, moon{
			position: p,
			velocity: point{x: 0, y: 0, z: 0},
		})
	}

	for i := 0; i < iter; i++ {
		mArr.applyGravity()
		mArr.applyVelocity()
	}
	fmt.Println(mArr.energy())
}

func (mArr moons) energy() int {
	sum := 0
	for _, m := range mArr {
		sum += m.energy()
	}
	return sum
}

func (m *moon) energy() int {
	return m.position.energy() * m.velocity.energy()
}

func (p *point) energy() int {
	return int(
		math.Abs(float64(p.x)) +
			math.Abs(float64(p.y)) +
			math.Abs(float64(p.z)),
	)
}

func (mArr moons) applyVelocity() {
	for idx, m := range mArr {
		p := point{
			x: m.position.x + m.velocity.x,
			y: m.position.y + m.velocity.y,
			z: m.position.z + m.velocity.z,
		}
		mArr[idx] = moon{
			position: p,
			velocity: m.velocity,
		}
	}
}

func (mArr moons) applyGravity() {
	for idx := range mArr {
		mArr[idx] = mArr.applyGravityToMoon(idx)
	}
}

func (mArr moons) applyGravityToMoon(idx int) moon {
	m := mArr[idx]
	for j, m2 := range mArr {
		if j == idx {
			continue
		}
		m.velocity = point{
			x: m.velocity.x + gravityDelta(m.position.x, m2.position.x),
			y: m.velocity.y + gravityDelta(m.position.y, m2.position.y),
			z: m.velocity.z + gravityDelta(m.position.z, m2.position.z),
		}
	}
	return m
}

func gravityDelta(a, b int) int {
	if b > a {
		return 1
	}
	if b == a {
		return 0
	}
	return -1
}

func newPoint(s string) (point, error) {
	s = sanitize(s, []string{">", "<", "=", " ", "x", "y", "z"})
	sArr := strings.Split(s, ",")
	intArr := []int{}
	for _, s := range sArr {
		v, err := strconv.Atoi(s)
		if err != nil {
			return point{}, err
		}
		intArr = append(intArr, v)
	}
	return point{
		x: intArr[0],
		y: intArr[1],
		z: intArr[2],
	}, nil
}

func (mArr moons) toString() string {
	var s string
	for idx, m := range mArr {
		if idx == 0 {
			s = m.toString()
			continue
		}
		s = strings.Join([]string{s, m.toString()}, "\n")
	}
	return s
}

func (m *moon) toString() string {
	return fmt.Sprintf(
		"pos=%s, vel=%s",
		m.position.toString(),
		m.velocity.toString(),
	)
}

func (p *point) toString() string {
	return fmt.Sprintf("<x=%d, y=%d, z=%d>", p.x, p.y, p.z)
}

func sanitize(s string, remove []string) string {
	for _, r := range remove {
		s = strings.ReplaceAll(s, r, "")
	}
	return s
}
