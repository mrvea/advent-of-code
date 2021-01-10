package vacation61

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
	rowMax        = 127.0
	rowMin        = 0.0
	columnMax     = 7.0
	columnMin     = 0.0
)

/**

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
	ans := make(map[rune]struct{})
	for s.Scan() {
		raw := s.Text()
		if raw == "" {
			total += len(ans)
			ans = make(map[rune]struct{})
			continue
		}
		for _, char := range raw {
			ans[char] = struct{}{}
		}
	}
	total += len(ans)

	fmt.Println("total: ", total)
}
