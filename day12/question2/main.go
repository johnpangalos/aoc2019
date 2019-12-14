package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/johnny88/aoc2019/fileparse"
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
	mInitial := make(moons, len(mArr))
	copy(mInitial, mArr)

	count := 1
	for {
		mArr.applyGravity()
		mArr.applyVelocity()
		if mArr.equals(mInitial) {
			fmt.Println(count)
			break
		}
		count++
	}
}

func (mArr moons) equals(mInitial moons) bool {
	equal := true
	for idx, m := range mArr {
		if m != mInitial[idx] {
			equal = false
			break
		}
	}
	return equal
}

func (mArr moons) applyVelocity() {
	for idx, m := range mArr {
		// calc per demension iteratively, i.e. comment out and run again
		p := point{
			// x: m.position.x + m.velocity.x,
			x: m.position.x,
			// y: m.position.y + m.velocity.y,
			y: m.position.y,
			z: m.position.z + m.velocity.z,
			// z: m.position.z,
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
		// calc per demension iteratively, i.e. comment out and run again
		m.velocity = point{
			// x: m.velocity.x + gravityDelta(m.position.x, m2.position.x),
			x: m.velocity.x,
			// y: m.velocity.y + gravityDelta(m.position.y, m2.position.y),
			y: m.velocity.y,
			z: m.velocity.z + gravityDelta(m.position.z, m2.position.z),
			// z: m.velocity.z,
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
