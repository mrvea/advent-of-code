package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/mrvea/advent-of-code/helpers"
)

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
)
const (
	X = iota
	Y
)

var origin = helpers.XY{0, 0}

type direction int
type axisType int

type vector struct {
	axis      axisType
	direction direction
	from      helpers.XY
	to        helpers.XY
}

func (v *vector) setTo(maginitude int) {
	x, y, d := v.from.Y, v.from.Y, int(v.direction)
	if v.axis == Y {
		y += maginitude * d
	}
	if v.axis == X {
		x += maginitude * d
	}

	v.to = helpers.XY{v.from.X + x, v.from.Y + y}
}

func (v *vector) cross(val int) bool {
	if v.axis == Y {
		return v.inY(val)
	}
	return v.inX(val)
}

func (v *vector) inX(val int) bool {
	x1, x2 := v.from.X, v.to.X
	if x1 > x2 {
		x1, x1 = x2, x1
	}
	return between(x1, x2, val)
}
func (v *vector) inY(val int) bool {
	y1, y2 := v.from.Y, v.to.Y
	if y1 < y2 {
		y1, y2 = y2, y1
	}
	return between(y1, y2, val)
}

func (v *vector) String() string {
	return fmt.Sprintf("\t\ndirection: %d, axis: %d\n", v.direction, v.axis)
}

func between(a, b, v int) bool {
	return a <= v && b >= v
}

type Wire struct {
	x []*vector
	y []*vector
}

func newWire() *Wire {
	w := &Wire{}
	w.x = make([]*vector, 0)
	w.y = make([]*vector, 0)
	return w
}
func (w *Wire) setSegment(v *vector) {
	if v.axis == Y {
		w.y = append(w.y, v)
		return
	}
	w.x = append(w.x, v)
}
func (w *Wire) getPool(v *vector) []*vector {
	if v.axis == X {
		return w.x
	}
	return w.y
}
func (w *Wire) intersections(v *vector) []helpers.XY {
	// list := make([]helpers.)
	// for v := range w.getPool(v) {

	// }
}

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	rr := bufio.NewReader(f)
	var numStr strings.Builder

	wires := make([]*Wire, 0)

	currentWire := newWire()
	wires = append(wires, currentWire)
	currentVector := &vector{}
	currentWire.setSegment(currentVector)
	currentVector.from = helpers.XY{0, 0}
	first := true
	intersections := make([]*vector, 0)
	for {
		r, _, err := rr.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("could not read rune, %s", err)
		}

		if first {
			fmt.Printf("first rune: %q\n", r)
			currentVector.axis, currentVector.direction = getDirection(r)
			// currentWire.setSegment(currentVector)
			first = false
			continue
		}

		if r == '\n' {
			fmt.Println("new wire...")
			currentWire = newWire()
			wires = append(wires, currentWire)
			r = ','
		}

		if strings.ContainsRune(",\n", r) {
			fmt.Printf("maginitude: %s\n", numStr.String())
			magnitude, err := strconv.Atoi(numStr.String())
			if err != nil {
				log.Fatalf("cound not convert '%s'", numStr.String())
			}
			currentVector.setTo(magnitude)
			if len(wires) > 1 {
				if val := getIntersections(wires, currentVector); val != nil {
					intersections = append(intersections, val...)
				}
			}

			currentVector = &vector{from: currentVector.to}
			currentWire.setSegment(currentVector)
			numStr.Reset()
			first = true

			if r == '\n' {
				fmt.Println("new wire...")
				currentWire = newWire()
				wires = append(wires, currentWire)
				r = ','
			}
			continue
		}
		fmt.Printf("rune: %q\n", r)
		numStr.WriteRune(r)
	}
	fmt.Println("distance:", getDistance(wires))
}

func getIntersections(wires []*Wire, v *vector) []*vector {
	l := len(wires) - 1
	intersections := make([]*helpers.XY, 0)
	for i = 0; i < l; i++ {
		if list := wire[i].intersections(v); list != nil {
			intersections := append(intersections, list...)
		}
	}
	return nil
}

func getDistance(wires []*Wire) int {
	return -1
}

func getDirection(r rune) (axisType, direction) {

	switch r {
	case 'R':
		return X, 1
	case 'U':
		return Y, 1
	case 'L':
		return X, -1
	case 'D':
		return Y, -1
	default:
		log.Fatalln("Unexpected direction rune ", r)
	}
	return -1, -2
}
