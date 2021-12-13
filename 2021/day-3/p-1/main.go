package p1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	year = 2021
	day  = 3
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
	// prev := -1
	// count := 0
	var collection []int

	for s.Scan() {
		text := s.Text()
		if err := s.Err(); err != nil {
			log.Fatal(err)
		}
		if collection == nil {
			collection = make([]int, len(text))
		}
		for i, r := range text {
			if r == '0' {
				collection[i]--
				continue
			}
			collection[i]++
		}
	}
	var g string
	var e string
	for _, d := range collection {
		if d < 0 {
			g += "0"
			e += "1"
			continue
		}
		g += "1"
		e += "0"
	}
	fmt.Printf("%c", 1)
	fmt.Println("strings: ", g, e)
	gamma, err := strconv.ParseInt(g, 2, 64)
	if err != nil {
		log.Panic(err)
	}
	epsilon, err := strconv.ParseInt(e, 2, 64)
	if err != nil {
		log.Panic(err)
	}
	// answer := findSummables(collection)
	fmt.Println("answer: ", gamma*epsilon)
}
