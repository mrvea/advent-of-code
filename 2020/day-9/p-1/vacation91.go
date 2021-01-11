package vacation91

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	day           = "9"
	basePath      = "2020/day-" + day + "/"
	inputName     = "input.txt"
	testInputName = "test_" + inputName
)

/**

**/
func SubMain(args ...string) {
	fmt.Println("vacation" + day + ".1 starting exectution...")
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
			break
		}
		preamble = append(preamble, num)
		// if len(preamble) < preableNumber {
		// 	continue
		// }

		// bagName, contains := processDirections(s.Text())

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
