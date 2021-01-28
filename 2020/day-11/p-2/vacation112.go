package vacation112

import (
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
		fmt.Println("text", s.Text())
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
		printSeats(seats)
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
	\ | /
	- 0 -
	/ | \
**/

func sit(i, j int, seats [][]rune, table map[string]bool) bool {
	// start := i - 1
	// end := j
	return !vS(i+1, j, seats) && !vN(i-1, j, seats) && !vW(i, j-1, seats) &&
		!vE(i, j+1, seats) && !vNE(i-1, j+1, seats) && !vNW(i-1, j-1, seats) &&
		!vSE(i+1, j+1, seats) && !vSW(i+1, j-1, seats)
}
func isFull(r rune) (next bool, value bool) {
	switch r {
	case '#':
		next, value = false, true
	case '.':
		next, value = true, false
	case 'L':
	}
	return
}

// down
func vS(i, j int, seats [][]rune) bool {
	// fmt.Printf("vS: i: %d, j:%d\n", i, j)
	for {
		if i >= len(seats) {
			return false
		}
		if next, value := isFull(seats[i][j]); !next {
			return value
		}
		i++
	}
}

// up
func vN(i, j int, seats [][]rune) bool {
	// fmt.Printf("vV: i: %d, j:%d\n", i, j)
	for {
		if i < 0 {
			return false
		}
		if next, value := isFull(seats[i][j]); !next {
			return value
		}
		i--
	}
}

// left
func vW(i, j int, seats [][]rune) bool {
	// fmt.Printf("vW: i: %d, j:%d\n", i, j)
	for {
		if j < 0 {
			return false
		}
		if next, value := isFull(seats[i][j]); !next {
			return value
		}
		j--
	}
}

// right
func vE(i, j int, seats [][]rune) bool {
	// fmt.Printf("vE: i: %d, j:%d\n", i, j)
	for {
		if j >= len(seats[0]) {
			return false
		}
		if next, value := isFull(seats[i][j]); !next {
			return value
		}
		j++
	}

}

// upper right
func vNE(i, j int, seats [][]rune) bool {
	for {
		if i < 0 || j >= len(seats[i]) {
			return false
		}
		if next, value := isFull(seats[i][j]); !next {
			return value
		}
		i--
		j++
	}
}

// upper left
func vNW(i, j int, seats [][]rune) bool {
	for {
		if i < 0 || j < 0 {
			return false
		}
		if next, value := isFull(seats[i][j]); !next {
			return value
		}
		i--
		j--
	}
}

// down right
func vSE(i, j int, seats [][]rune) bool {
	for {
		if i >= len(seats) || j >= len(seats[i]) {
			return false
		}
		if next, value := isFull(seats[i][j]); !next {
			return value
		}
		i++
		j++
	}
}

// down left
func vSW(i, j int, seats [][]rune) bool {

	for {
		if i >= len(seats) || j < 0 {
			return false
		}
		if next, value := isFull(seats[i][j]); !next {
			return value
		}
		i++
		j--
	}
}

/**
	0 0 0
	0 1 0
	0 0 0
**/
func up(i, j int, seats [][]rune, table map[string]bool) bool {
	total := 0
	max := 5
	list := []func() bool{
		func() bool { return vS(i+1, j, seats) },
		func() bool { return vN(i-1, j, seats) },
		func() bool { return vW(i, j-1, seats) },
		func() bool { return vE(i, j+1, seats) },
		func() bool { return vNE(i-1, j+1, seats) },
		func() bool { return vNW(i-1, j-1, seats) },
		func() bool { return vSE(i+1, j+1, seats) },
		func() bool { return vSW(i+1, j-1, seats) },
	}

	for _, fn := range list {
		if fn() {
			total++
		}
		if total >= max {
			return true
		}
	}
	// for p := helpers.Max(i-1, 0); p < helpers.Min(i+2, len(seats)); p++ {
	// 	if p < 0 {
	// 		continue
	// 	}
	// 	for r := helpers.Max(j-1, 0); r < helpers.Min(j+2, len(seats[0])); r++ {
	// 		if i == p && r == j {
	// 			continue
	// 		}

	// 		if seats[p][r] == '#' {
	// 			total++
	// 			if total >= max {
	// 				return true
	// 			}
	// 		}
	// 	}
	// }
	return false
}

func printSeats(table [][]rune) {
	for _, row := range table {
		fmt.Println(string(row))
	}
	fmt.Println()
}
