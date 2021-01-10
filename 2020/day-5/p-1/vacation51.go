package vacation51

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

const (
	day           = "5"
	basePath      = "2020/day-" + day + "/"
	inputName     = "input.txt"
	testInputName = "test_" + inputName
	rowMax        = 127.0
	rowMin        = 0.0
	columnMax     = 7.0
	columnMin     = 0.0
)

/**
	F => front
	B => back
	L => left
	R => right
	first 7 are F or B
	128 rows, from 0 to 127
	front 0 to 63 or back 64 to 127
	example
	Start by considering the whole range, rows 0 through 127.
	F means to take the lower half, keeping rows 0 through 63.
	B means to take the upper half, keeping rows 32 through 63.
	F means to take the lower half, keeping rows 32 through 47.
	B means to take the upper half, keeping rows 40 through 47.
	B keeps rows 44 through 47.
	F keeps rows 44 through 45.
	The final F keeps the lower of the two, row 44.

	last 3 char are L or R
	8 columns => 0 to 7
	example:
	Start by considering the whole range, columns 0 through 7.
	R means to take the upper half, keeping columns 4 through 7.
	L means to take the lower half, keeping columns 4 through 5.
	The final R keeps the upper of the two, column 5
	example
	FBFBBFFRLR reveals that it is the seat at row 44, column 5

	multiply the row by 8, then add the column. In this example, the seat has ID 44 * 8 + 5 = 357.
**/
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

	total := 0
	var seats [128][8]string
	for s.Scan() {
		row, col, id := getID(s.Text())
		seats[row][col] = "X"
		total = id
	}
	for row, r := range seats {
		if row == 0 || row == len(seats)-1 {
			continue
		}
		for col, c := range r {
			if c != "X" {
				fmt.Print(row*8 + col)
			}
			fmt.Print(c)
		}
		fmt.Println()
	}

	fmt.Println("total: ", total)
}

func getID(raw string) (int, int, int) {
	rMax := rowMax
	rMin := rowMin
	row := -1.0
	cMax := columnMax
	cMin := columnMin
	col := -1.0
	for i := 0; i < 7; i++ {
		char := raw[i]
		switch char {
		case 'F':
			rMax = rMin + math.Floor((rMax-rMin)/2)
			row = rMax
		case 'B':
			rMin = rMin + math.Ceil((rMax-rMin)/2)
			row = rMin
		default:
			log.Fatalf("unknown row char %s", string(char))
		}
	}

	for i := 0; i < 3; i++ {
		char := raw[7+i]
		switch char {
		case 'R':
			cMin = cMin + math.Ceil((cMax-cMin)/2)
			col = cMin
		case 'L':
			cMax = cMin + math.Floor((cMax-cMin)/2)
			col = cMax
		default:
			log.Fatalf("unknown column char %s", string(char))
		}
	}
	fmt.Printf("row %d, column %d, seat ID %d\n\n", int(row), int(col), int(row*8+col))
	return int(row), int(col), int(row*8.0 + col)

}
