package vacation12

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func SubMain(args ...string) {
	fmt.Println("vacation1.2 starting exectution...")
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
		fmt.Printf("current: %d\nremainder: %d", currentNum, remainder)
		fmt.Println()
		subCollection := make(map[int]struct{}, 0)
		for key, value := range collection {
			if key < remainder {
				subCollection[key] = value
			}
		}
		secondNum := 0
		fmt.Println(subCollection)
	INNER2:
		for {
			if len(subCollection) == 0 {
				break INNER2
			}
		INNER3:
			for key := range subCollection {
				secondNum = key
				delete(subCollection, key)
				break INNER3
			}
			last := remainder - secondNum
			fmt.Printf("...inner => second: %d, remaider: %d\n", secondNum, remainder)
			if last < 0 {
				continue
			}
			if _, ok := subCollection[last]; ok {
				fmt.Printf("current: %d, second: %d, remaider: %d", currentNum, secondNum, last)
				return currentNum * last * secondNum
			}
		}

	}
	return 0
}
