package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	m := map[int]int{}
	nums := make([]int, 0)
	var total int

	for s.Scan() {
		var num int
		fmt.Sscanf(s.Text(), "%d", &num)
		nums = append(nums, num)
		total += num
		m[total]++
		if m[total] > 1 {
			log.Fatalf("ans: %d\n", total)
		}
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	for {
		for _, num := range nums {
			total += num
			m[total]++
			if m[total] > 1 {
				log.Fatalf("ans: %d\n", total)
			}
		}
	}
}
