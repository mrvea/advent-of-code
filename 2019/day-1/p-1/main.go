package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	var total int
	for s.Scan() {
		var num int
		fmt.Sscanf(s.Text(), "%d", &num)
		total += getFuel(num)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("total: ", total)
}

func getFuel(num int) int {
	return int(math.Floor(float64(num)/3) - 2)
}
