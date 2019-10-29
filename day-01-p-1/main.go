package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	var total int
	for s.Scan() {
		var num int
		fmt.Sscanf(s.Text(), "%d", &num)
		total += num
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(total)
}
