package p2

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
	// count := 0
	collection := [3]int{}
	size := len(collection)
	sum := 0
	i := 0
	cur := 0
	count := 0

	for s.Scan() {
		var num int
		fmt.Sscanf(s.Text(), "%d", &num)
		fmt.Printf("cur: %d, i: %d , num: %d => \n\t", cur, i, num)
		if sum != 0 && i == cur {
			action := " Decreased"
			fmt.Printf("window: %d to %d", prev, sum)
			if prev == -1 {
				action = " N/A - no previous measurement"
				prev = num
				fmt.Println(action)
			}
			if prev < sum {
				action = " increased"
				count++
			}
			prev = sum
			fmt.Printf(" removing: %d =>", collection[i])
			sum -= collection[i]
			i = (i + 1) % size
			fmt.Println(action)
		}
		collection[cur] = num
		sum += num

		cur = (cur + 1) % size
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	// answer := findSummables(collection)
	fmt.Println("answer: ", count)
}
