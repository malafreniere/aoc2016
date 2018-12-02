package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/golang/geo/r3"
)

func main() {
	inputs := "L1, R3, R1, L5, L2, L5, R4, L2, R2, R2, L2, R1, L5, R3, L4, L1, L2, R3, R5, L2, R5, L1, R2, L5, R4, R2, R2, L1, L1, R1, L3, L1, R1, L3, R5, R3, R3, L4, R4, L2, L4, R1, R1, L193, R2, L1, R54, R1, L1, R71, L4, R3, R191, R3, R2, L4, R3, R2, L2, L4, L5, R4, R1, L2, L2, L3, L2, L1, R4, R1, R5, R3, L5, R3, R4, L2, R3, L1, L3, L3, L5, L1, L3, L3, L1, R3, L3, L2, R1, L3, L1, R5, R4, R3, R2, R3, L1, L2, R4, L3, R1, L1, L1, R5, R2, R4, R5, L1, L1, R1, L2, L4, R3, L1, L3, R5, R4, R3, R3, L2, R2, L1, R4, R2, L3, L4, L2, R2, R2, L4, R3, R5, L2, R2, R4, R5, L2, L3, L2, R5, L4, L2, R3, L5, R2, L1, R1, R3, R3, L5, L2, L2, R5"

	pos := puzzle1(inputs)

	fmt.Printf("Distance: %v\n", math.Abs(pos.X)+math.Abs(pos.Y))

	pos = puzzle2(inputs)

	fmt.Printf("Distance: %v\n", math.Abs(pos.X)+math.Abs(pos.Y))
}

func puzzle1(inputs string) r3.Vector {
	dir, pos := r3.Vector{X: 0, Y: 1, Z: 0}, r3.Vector{X: 0, Y: 0, Z: 0}

	for _, s := range strings.Split(inputs, ", ") {
		dir = rotate90(dir, s[0])
		pos = pos.Add(dir.Mul(toFloat64(s[1:])))
	}

	return pos
}

func puzzle2(inputs string) r3.Vector {
	dir, pos := r3.Vector{X: 0, Y: 1, Z: 0}, r3.Vector{X: 0, Y: 0, Z: 0}

	visited := make(map[string]bool)
	for _, s := range strings.Split(inputs, ", ") {
		dir = rotate90(dir, s[0])

		for i := float64(0); i < toFloat64(s[1:]); i++ {
			pos = pos.Add(dir)

			_, exists := visited[pos.String()]

			if exists {
				return pos
			} else {
				visited[pos.String()] = true
			}
		}
	}

	panic(fmt.Errorf("not found"))
}

func toFloat64(v string) float64 {
	r, err := strconv.Atoi(v)

	if err != nil {
		panic(err)
	}

	return float64(r)
}

func rotate90(v r3.Vector, dir byte) r3.Vector {
	switch dir {
	case 'R':
		return r3.Vector{X: v.Y, Y: -v.X}
	case 'L':
		return r3.Vector{X: -v.Y, Y: v.X}
	}

	panic(fmt.Errorf("invalid rotation direction %q", dir))
}
