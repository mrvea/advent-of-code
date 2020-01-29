package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const exit = 99

const (
	add = iota + 1
	mul
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// s := bufio.NewScanner(f)
	rr := bufio.NewReader(f)
	var numStr strings.Builder
	// for s.Scan() {
	// 	programStr = s.Text()
	// }
	program := make(map[int]int, 0)
	var index int
	for {
		r, _, err := rr.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("could not read rune, %s", err)
		}
		if r == ',' {
			program[index], err = strconv.Atoi(numStr.String())
			if err != nil {
				log.Fatalf("cound not very %s", numStr.String())
			}
			numStr.Reset()
			index++
			continue
		}
		numStr.WriteRune(r)
	}
	fmt.Println("program: ", program)
	// program[1], program[2] = 12, 2

MainLoop:
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			program[1], program[2] = i, j
			output := execute(program)
			if 19690720 == output[0] {
				fmt.Printf("result: %d\n", 100*i+j)
				break MainLoop
			}
		}
	}
	// fmt.Println("moded program: ", execute(program)[0])
}

func execute(program map[int]int) map[int]int {
	var alt = make(map[int]int, len(program))
	for i, v := range program {
		alt[i] = v
	}
	var index int
	const length = 4

	for {
		state := alt[index]
		if state == exit {
			return alt
		}

		if len(program) < index+3 {

			log.Fatal("Length exceeds the expected", alt[0])
		}

		l1, l2, l3 := alt[index+1], alt[index+2], alt[index+3]
		// fmt.Printf("state: %d, l1: %d, l2: %d, l3: %d\n", state, l1, l2, l3)
		v1, ok := alt[l1]
		if !ok {
			log.Fatal("no value at location ", l1)
		}

		v2, ok := alt[l2]
		if !ok {
			log.Fatal("no value at location ", l2)
		}

		alt[l3] = getValue(state, v1, v2)

		prettyPrint(state, l1, l2, l3, v1, v2, alt[l3])
		index += length
		if len(program) == index {
			return alt
		}
	}
}

func getValue(state, v1, v2 int) int {
	if state == add {
		return v1 + v2
	}
	return v1 * v2
}

func prettyPrint(state, l1, l2, l3, v1, v2, result int) {
	var stateString string
	switch state {
	case add:
		stateString = "+"
	case mul:
		stateString = "*"
	case exit:
		stateString = "exiting"
	default:
		stateString = "unknown"
	}
	fmt.Printf(`
	%d %s %d = %d
	l1 => %d, l2 => %d, l3 => %d

	`,
		v1,
		stateString,
		v2,
		result,
		l1, l2, l3,
	)
}
