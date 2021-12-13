package p1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	year = 2021
	day  = 3
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
	collection := make([]string, 0)

	for s.Scan() {
		text := s.Text()
		if err := s.Err(); err != nil {
			log.Fatal(err)
		}
		collection = append(collection, text)

	}

	oxygenGeneratorRating := oxygenGeneratorRating(collection)
	co2ScrubberRating := co2ScrubberRating(collection)

	fmt.Println("strings: ", oxygenGeneratorRating, co2ScrubberRating)
	gamma, err := strconv.ParseInt(oxygenGeneratorRating, 2, 64)
	if err != nil {
		log.Panic(err)
	}
	epsilon, err := strconv.ParseInt(co2ScrubberRating, 2, 64)
	if err != nil {
		log.Panic(err)
	}
	// answer := findSummables(collection)
	fmt.Printf("oxygen generator rating: %d\nCO2 scrubber ratting: %d\n", gamma, epsilon)
	fmt.Println("answer: ", gamma*epsilon)
}

func oxygenGeneratorRating(pool []string, is ...int) string {

	fmt.Println("pool: ", pool)
	if len(pool) == 1 {
		return pool[0]
	}

	i := 0
	if len(is) > 0 {
		i = is[0]
	}
	up, down := groupBinary(pool, i)
	i++
	if len(up) >= len(down) {
		return oxygenGeneratorRating(up, i)
	}

	return oxygenGeneratorRating(down, i)

}
func co2ScrubberRating(pool []string, is ...int) string {
	fmt.Println("co2 pool: ", pool)
	if len(pool) == 1 {
		return pool[0]
	}

	i := 0
	if len(is) > 0 {
		i = is[0]
	}
	up, down := groupBinary(pool, i)
	i++
	if len(up) < len(down) {
		return co2ScrubberRating(up, i)
	}

	return co2ScrubberRating(down, i)

}

func groupBinary(pool []string, i int) ([]string, []string) {
	up := make([]string, 0)
	down := make([]string, 0)
	for _, str := range pool {
		if str[i] == '1' {
			up = append(up, str)
			continue
		}
		down = append(down, str)
	}
	return up, down
}

// func isOne(str byte) bool {

// 	if str != "" && str[0] == '1' {
// 		return true
// 	}
// 	return false
// }
