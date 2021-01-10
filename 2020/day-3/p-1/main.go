package vacation31

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	day           = "3"
	basePath      = "2020/day-" + day + "/"
	inputName     = "input.txt"
	testInputName = "test_" + inputName
)

func SubMain(args ...string) {
	fmt.Println("vacation" + day + ".1 starting exectution...")
	filePath := basePath + inputName
	if len(args) > 0 && args[0] == "test" {
		fmt.Println("Running test file...")
		filePath = basePath + testInputName
	}
	fmt.Printf("filePath: %s\n", filePath)
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	grid := make([]string, 0)
	for s.Scan() {
		grid = append(grid, s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("total: ", getEncounteredTrees(grid))
}

func getEncounteredTrees(grid []string) int {
	d, r, max, t := 1, 3, len(grid), 0

	x, y := r, d
	for y < max {
		row := grid[y]
		char := row[x]
		if '#' == char {
			t++
		}
		y += d
		x = (x + r) % len(row)
	}
	return t
}
