package vacation62

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	day           = "6"
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

	total := 0
	ans := make(map[rune]int)
	people := 0
	for s.Scan() {
		raw := s.Text()
		if raw == "" {
			fmt.Printf("people: %d", people)
			fmt.Println(ans)
			for _, c := range ans {
				if c == people {
					total++
				}
			}
			people = 0
			ans = make(map[rune]int)
			continue
		}
		people++
		for _, char := range raw {
			ans[char]++
		}
	}
	for _, c := range ans {
		if c == people {
			total++
		}
	}

	fmt.Println("total: ", total)
}
