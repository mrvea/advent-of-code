package vacation11

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func SubMain(args ...string) {
	fmt.Println("vacation1.1 starting exectution...")
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	f, err := os.Open("2020/day-1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	collection := make(map[int]struct{}, 0)
	for s.Scan() {
		var num int
		fmt.Sscanf(s.Text(), "%d", &num)
		collection[num] = struct{}{}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	answer := findSummables(collection)
	fmt.Println("total: ", answer)
}

func findSummables(collection map[int]struct{}) int {

	currentNum := 0
	for {
		if len(collection) == 0 {
			break
		}
		// grab any key and remove it from collection;
	INNER:
		for key := range collection {
			currentNum = key
			delete(collection, key)
			break INNER
		}
		remainder := 2020 - currentNum

		if _, ok := collection[remainder]; ok {
			return currentNum * remainder
		}
	}
	return 0
}
