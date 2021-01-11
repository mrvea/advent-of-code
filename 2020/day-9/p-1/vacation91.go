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

	preamble := make(map[int]struct{})
	skip := false
	for s.Scan() {
		fmt.Println(s.Text())
		num, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Panic(err)
		}

		if !skip && Invalid(num, preamble) {
			total = num
			break
		}
		if len(preamble) < preableNumber {

			continue
		}
		preamble[num] = struct{}{}
		// bagName, contains := processDirections(s.Text())

	}

	fmt.Println("total: ", total)
}

func Invalid(num int, preamble map[int]struct{}) bool {
	for i := 1; i < num; i++ {
		if _, ok := preamble[i]; ok {
			if _, ok = preamble[num-i]; ok {
				return false
			}

		}
	}
	return true
}
