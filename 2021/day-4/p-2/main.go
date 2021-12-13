package p1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	year = 2021
	day  = 4
	part = 2
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
	// prev := -1
	// count := 0
	var collection []int
	boardMaps := make([]map[int][][3]int, 0)
	boards := make([][5][5]int, 0)
	var b [5][5]int
	var m map[int][][3]int
	y := -1
	for s.Scan() {
		text := s.Text()
		if err := s.Err(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("text: ", text)
		if text == "" {
			m = make(map[int][][3]int, 0)
			b = [5][5]int{}
			boardMaps = append(boardMaps, m)
			boards = append(boards, b)
			y = 0
			continue
		}
		if collection == nil {
			parts := strings.Split(text, ",")
			collection = make([]int, len(parts))
			for i, numStr := range parts {
				collection[i], err = strconv.Atoi(numStr)
				if err != nil {
					log.Panic(err)
				}
			}
			continue
		}
		var x = [5]int{}
		fmt.Sscanf(text, "%d %d %d %d %d", &x[0], &x[1], &x[2], &x[3], &x[4])
		for i, n := range x {
			m[n] = append(m[n], [3]int{i, y})
		}
		y++
	}
	for i, m := range boardMaps {
		fmt.Println("map: ", i)
		fmt.Println(m)
	}
	// fmt.Println("map: ", boardMaps)
	// fmt.Println("boards", boards)
	winners, last, boardIndex := bingo(collection, boardMaps, boards)
	// answer := findSummables(collection)
	fmt.Println("winning numbers: ", winners, last)
	sum := 0

EX:
	for num, pos := range boardMaps[boardIndex] {
		fmt.Println("index pos: ", pos)
		if pos[0][2] == 1 {
			fmt.Println("marked: ", num)
			continue EX
		}
		sum += num
	}
	fmt.Println("sum: ", sum)
	fmt.Println("answer: ", sum*last)
}

func bingo(pool []int, index []map[int][][3]int, boards [][5][5]int) ([5]int, int, int) {
	// size := 5;
	// offset := 0;
	size := len(index)
	counts := make([][2][5]int, size)
	won := make(map[int]bool, 0)
	for _, num := range pool {
		fmt.Printf("number pulled: %d\n", num)
		for i := 0; i < size; i++ {
			if won[i] {
				continue
			}
			pos, ok := index[i][num]
			if !ok {
				continue
			}
			for numIndex, p := range pos {
				x, y := p[0], p[1]
				fmt.Printf("x: %d, y: %d\n", x, y)
				if boards[i][x][y] != 0 {
					continue
				}
				boards[i][x][y] = num
				index[i][num][numIndex][2] = 1
				counts[i][0][x]++
				counts[i][1][y]++

				if counts[i][0][x] == 5 {

					fmt.Println("board won: ", i)

					// if len(index) > 1 {
					// 	index = removeAt(index, i)
					// 	size = len(index)
					// 	continue
					// }
					won[i] = true
					if size != len(won) {
						continue
					}
					return boards[i][x], num, i
				}
				if counts[i][1][y] == 5 {
					fmt.Println("board won: ", i)
					// if len(index) > 1 {
					// 	index = removeAt(index, i)
					// 	size = len(index)
					// 	continue
					// }
					col := [5]int{}
					b := boards[i]
					for i, n := range b {
						col[i] = n[y]
					}

					// fmt.Println("board won: ", i)
					// if len(index) > 1 {
					// 	index = removeAt(index, i)
					// 	size = len(index)
					// 	continue
					// }
					won[i] = true
					if size != len(won) {
						continue
					}
					return col, num, i
				}
			}
		}

	}
	return [5]int{}, 0, 0
}

// func removeAt(pool []map[int][][3]int, i int) []map[int][][3]int {
// 	top := pool[:i]
// 	top = append(top, pool[i+1:]...)
// 	return top
// }
