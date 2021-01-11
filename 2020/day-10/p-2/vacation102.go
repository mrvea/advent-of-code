package vacation102

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

const (
	day           = "10"
	basePath      = "2020/day-" + day + "/"
	inputName     = "input.txt"
	testInputName = "test_" + inputName
)

/**

**/
func SubMain(args ...string) {
	fmt.Println("vacation" + day + ".2 starting exectution...")
	filePath := basePath + inputName
	preableNumber := 25
	if len(args) > 0 && args[0] == "test" {
		fmt.Println("Running test file...")
		filePath = basePath + testInputName
		preableNumber = 5
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
	// insts := make([]*Instruction, 0)

	preamble := make([]int, 0)

	for s.Scan() {
		fmt.Println(s.Text())
		num, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Panic(err)
		}

		if len(preamble) >= preableNumber && Invalid(num, preamble, preableNumber) {
			total = num
			conPool := findContiguousSet(num, preamble)
			if conPool == nil {
				log.Panic("contiguous pool is empty")
			}
			min, max := findMinMax(conPool)
			total = min + max
			break
		}
		preamble = append(preamble, num)

	}

	fmt.Println("total: ", total)
}

func Invalid(num int, preamble []int, span int) bool {
	l := len(preamble)
	end := l - span
	for i := l - 1; i > end; i-- {
		s := num - preamble[i]
		if in(preamble[end:i], s) {
			return false
		}
	}

	return true
}

func in(pool []int, needle int) bool {
	for _, v := range pool {
		if v == needle {
			return true
		}
	}
	return false
}

func findContiguousSet(to int, pool []int) []int {
	// fmt.Println("pool: ", pool)
OUTER:
	for i, v := range pool[0 : len(pool)-2] {
		total := v
		// fmt.Printf("to: %d => start: %d", to, total)
		for j, v2 := range pool[i+1:] {
			total += v2
			// fmt.Printf(", next: %d => %d", v2, total)
			if total == to {
				// fmt.Printf("\ni: %d, j: %d =>", i, i+j+2)
				// fmt.Println(pool[i : i+j+2])
				return pool[i : i+j+2]
			}
			if total > to {
				// fmt.Println()
				continue OUTER
			}
		}
		// fmt.Println()
	}
	fmt.Println("not found a contiguous set...")
	return nil
}

func findMinMax(pool []int) (min int, max int) {
	if len(pool) == 0 {
		fmt.Println("min max pool is empty")
		return
	}
	sort.Ints(pool)
	min, max = pool[0], pool[len(pool)-1]
	return
}
