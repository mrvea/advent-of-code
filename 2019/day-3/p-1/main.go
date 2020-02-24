package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/play/advent-of-code/helpers"
)

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
)

var origin = helpers.XY{0, 0}

type direction int

type vector struct {
	direction direction
	magnitude int
	from      helpers.XY
}

func (v *vector) String() string {
	return fmt.Sprintf("\t\ndirection: %d, magnitude: %d\n", v.direction, v.magnitude)
}
func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	rr := bufio.NewReader(f)
	var numStr strings.Builder

	wires := make([][]*vector, 0)
	currentWire := make([]*vector, 0)
	wires = append(wires, currentWire)
	var currentVector *vector
	first := true

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
			currentVector = &vector{direction: getDirection(r)}
			currentWire = append(currentWire, currentVector)
			first = false
			continue
		}

		if r == '\n' {
			fmt.Println("new wire...")
			wires = append(wires, currentWire)
			currentWire = make([]*vector, 0)
			r = ','
		}

		if r == ',' {
			fmt.Printf("maginitude: %s\n", numStr.String())
			currentVector.magnitude, err = strconv.Atoi(numStr.String())
			if err != nil {
				log.Fatalf("cound not convert '%s'", numStr.String())
			}

			numStr.Reset()
			first = true
			continue
		}
		fmt.Printf("rune: %q\n", r)
		numStr.WriteRune(r)
	}
	fmt.Println("distance:", getDistance(wires))
}

func getDistance(wires [][]*vector) int {
	return -1
}

func getDirection(r rune) direction {
	switch r {
	case 'R':
		return RIGHT
	case 'U':
		return UP
	case 'L':
		return LEFT
	case 'D':
		return DOWN
	default:
		log.Fatalln("Unexpected direction rune ", r)
	}
	return -1
}
