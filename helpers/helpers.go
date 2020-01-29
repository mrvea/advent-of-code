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

// Gcd find the greatest common factor of 2 integers.
// a is always the largest number.
func Gcd(a, b int64) int64 {
	r := a % b
	if r == 0 {
		return b
	}
	return Gcd(b, r)
}

// PrimeTest tests a number if it is  a prime with givin random number
func PrimeTest(n, a int64) string {
	if Gcd(n, a) == 1 {
		return "prime"
	}
	if Pow(a, n-1)%n == 1 {
		return "possibly prime"
	}
	return "congigate"
}

// PowRecursive returns number x to the power of e with recursive calls;
func PowRecursive(x int64, e int64) int64 {
	if e == 0 {
		return 1
	}
	val := Pow(x, e/2)
	if val%2 == 1 {
		return val * val * x
	}
	return val * val
}

// Pow return number x to the power of e
func Pow(x, e int64) int64 {
	val := int64(1)
	for e > 0 {
		if e&1 == 1 {
			val *= x
		}
		x *= x
		e >>= 1
	}
	return val
}

func Fib(n int) int {
	f := 1
	p := 1
	if n < 3 {
		return f
	}
	for i := 3; i < n; i++ {
		temp := f
		f = p + f
		p = temp
	}
	return f
}
