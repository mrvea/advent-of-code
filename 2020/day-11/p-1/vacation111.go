package vacation111

import (
	"advent-of-code/helpers"
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	day           = "11"
	basePath      = "2020/day-" + day + "/"
	inputName     = "input.txt"
	testInputName = "test_" + inputName
)

/**

**/
func SubMain(args ...string) {
	fmt.Println("vacation" + day + ".1 starting exectution...")
	filePath := basePath + inputName
	if len(args) > 0 && args[0] == "test" {
		fmt.Println("Running test file...")
		filePath = basePath + testInputName
		if len(args) > 1 {
			filePath = basePath + args[1]
		}
	}

	fmt.Printf("filePath: %s\n", filePath)
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	total := 0
	seats := make([][]rune, 0)
	for s.Scan() {
		fmt.Println(s.Text())
		seats = append(seats, []rune(s.Text()))
	}
	total = occupied(seats)
	// printSeats(seats)
	fmt.Println("total: ", total)
}

func occupied(seats [][]rune) int {

	total := 0
	table := make(map[string]bool)
	count := 0
	change := true
	for {
		if !change {
			break
		}
		change = false
		for i, row := range seats {
			// fmt.Printf("%v row\n", row)
			for j, char := range row {
				// fmt.Printf("\t%v char => ", string(char))
				switch char {
				case '.':
					// fmt.Println()
					continue
				case 'L':
					// fmt.Printf(" found am empty seat: (%d, %d)\n", i, j)
					key := fmt.Sprintf("%d,%d", i, j)
					// table[] = false
					if sit(i, j, seats, table) {
						// seats[i][j] = '#'
						table[key] = true
						change = true
					}
				case '#':
					// fmt.Printf("\tfound am occupied seat: (%d, %d)\n", i, j)
					key := fmt.Sprintf("%d,%d", i, j)
					// table[key] = true
					if up(i, j, seats, table) {
						// seats[i][j] = 'L'
						table[key] = false
						change = true
					}
				default:
					log.Panicf("unexpected char %s", string(char))
				}
			}
		}
		for i, row := range seats {
			for j := range row {
				key := fmt.Sprintf("%d,%d", i, j)
				if v, ok := table[key]; ok {
					if v {
						seats[i][j] = '#'
						continue
					}
					seats[i][j] = 'L'
				}
			}
		}
		// printSeats(seats)
		count++
	}

	for _, row := range seats {
		for _, char := range row {
			if char == '#' {
				total++
			}
		}
	}
	return total
}

/**
	0 0 0
	0 0 0
	0 0 0
**/

func sit(i, j int, seats [][]rune, table map[string]bool) bool {
	// start := i - 1
	// end := j
	for p := helpers.Max(i-1, 0); p < helpers.Min(i+2, len(seats)); p++ {
		if p < 0 {
			continue
		}
		for r := helpers.Max(j-1, 0); r < helpers.Min(j+2, len(seats[0])); r++ {
			if i == p && r == j {
				continue
			}

			if seats[p][r] == '#' {
				return false
			}
		}
	}
	return true
}

/**
	0 0 0
	0 1 0
	0 0 0
**/
func up(i, j int, seats [][]rune, table map[string]bool) bool {
	total := 0
	max := 4
	for p := helpers.Max(i-1, 0); p < helpers.Min(i+2, len(seats)); p++ {
		if p < 0 {
			continue
		}
		for r := helpers.Max(j-1, 0); r < helpers.Min(j+2, len(seats[0])); r++ {
			if i == p && r == j {
				continue
			}

			if seats[p][r] == '#' {
				total++
				if total >= max {
					return true
				}
			}
		}
	}
	return false
}

func printSeats(table [][]rune) {
	for _, row := range table {
		fmt.Println(string(row))
	}
	fmt.Println()
}
