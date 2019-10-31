package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	rr := bufio.NewReader(f)
	var twos, threes int
	var t, th int
	m := map[rune]int{}
	for {
		r, _, err := rr.ReadRune()
		if err != nil && err != io.EOF {
			log.Fatalf("cound not read rune from file: %s", err)
		}

		if r == '\n' || err == io.EOF {
			if t > 0 {
				twos++
			}
			if th > 0 {
				threes++
			}
			if err == io.EOF {
				break
			}
			t, th = 0, 0
			m = map[rune]int{}
		}
		m[r]++
		switch m[r] {
		case 2:
			t++
		case 3:
			t--
			th++
		default:
			if m[r] > 3 {
				th--
			}
		}

	}
	fmt.Println(twos * threes)
}
