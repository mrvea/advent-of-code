package p1

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	year = 2021
	day  = 1
	part = 1
)

var (
	Path      = fmt.Sprintf("%d/day-%d", year, day)
	inputFile = fmt.Sprintf("%s/input.txt", Path)
	logPrefix = fmt.Sprintf("saving christmas year: %d, day: %d, part: %d", year, day, part)
)

func SubMain(args ...string) {
	fmt.Println(logPrefix + " exectution...")
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	prev := -1
	count := 0
	// collection := make(map[int]struct{}, 0)
	for s.Scan() {
		var num int
		fmt.Sscanf(s.Text(), "%d", &num)
		fmt.Printf("%d to %d", prev, num)
		action := " decreased"
		if prev == -1 {
			action = " N/A - no previous measurement"
			prev = num
		}
		if prev < num {
			action = " increased"
			count++
		}
		fmt.Println(action)
		prev = num
		// collection[num] = struct{}{}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	// answer := findSummables(collection)
	fmt.Println("answer: ", count)
}
