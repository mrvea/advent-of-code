package p1

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	year = 2021
	day  = 2
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
	forward := 0
	depth := 0
	// count := 0
	// collection := make(map[int]struct{}, 0)
	nav := map[string]int{
		"up":   -1,
		"down": 1,
	}
	for s.Scan() {
		var num int
		var direction string
		// fields := strings.Fields(s.Text())
		// direction := fields[0]
		// amount, err := strconv.Atoi(fields[1])
		// if err != nil {
		// 	log.Panic(err)
		// }
		fmt.Sscanf(s.Text(), "%s %d", &direction, &num)
		fmt.Printf("%s to %d\n", direction, num)
		if direction == "forward" {
			forward += num
			continue
		}

		d := nav[direction]

		depth = depth + (num * d)

		// action := " decreased"

		// collection[num] = struct{}{}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	// answer := findSummables(collection)
	fmt.Println("answer: ", forward*depth)
}
