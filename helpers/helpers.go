package helpers

import "fmt"

// XY coor struct to hold the xy of possition and value(owner in this instance)
type XY struct{ x, y int }

func (xy *XY) withOffset(offsets ...int) *XY {
	x, y := xy.x, xy.y
	for i, n := range offsets {
		switch i {
		case 0:
			x += n
		case 1:
			y += n
		default:
			fmt.Printf("unsupported number of offests: %d", i)
			break
		}
	}
	return NewXY(x, y)
}

func (xy *XY) String() string {
	return fmt.Sprintf("(%d, %d)", xy.x, xy.y)
}

// NewXY creates a new xy coor object
func NewXY(x, y int) *XY {
	return &XY{x, y}
}

// Abs converts a negative to positive
func Abs(v int) int {
	y := v >> 7
	return (v ^ y) - y
}

// Max finds larger of 2 ints
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ManhattanDistance finds right angle distance between to points.
func ManhattanDistance(a, b *XY) uint {
	return uint(Abs(a.x-b.x) + Abs(a.y-b.y))
}

// Pair finds if 2 runes have Upper to Lower case relation
func Pair(a, b rune) bool {
	if a == b {
		return false
	}
	const dist = 32
	if Abs(int(a)-int(b)) == dist {
		return true
	}
	return false
}

// Min finds the smallest of 2 intergers
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Odd finds if the number is odd
func Odd(n int) bool {
	return n&1 == 1
}
