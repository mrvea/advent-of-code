package vacation21

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func SubMain(args ...string) {
	fmt.Println("vacation2.1 starting exectution...")
	filePath := "2020/day-2/input.txt"
	if len(args) > 1 && args[0] == "test" {
		filePath = "2020/day-2/test_input.txt"
	}
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	total := 0
	for s.Scan() {
		min, max, char, pool := extract(s.Text())
		total += isValid(min, max, char, pool)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("total: ", total)
}

func extract(raw string) (int, int, byte, string) {
	var rawCounts, rawChar, pool string
	fmt.Sscanf(raw, "%s %s %s", &rawCounts, &rawChar, &pool)
	fmt.Printf("c: %s, char: %s, p:%s\n", rawCounts, rawChar, pool)
	counts, char := getCounts(rawCounts), getChar(rawChar)
	return counts[0], counts[1], char, pool
}

func getCounts(raw string) []int {
	parts := strings.Split(raw, "-")
	nums := make([]int, 2)
	for i, numStr := range parts {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatalf("could not covert number: %s", numStr)
		}
		nums[i] = num
	}
	return nums
}

func getChar(raw string) byte {
	return raw[0]
}

func isValid(min, max int, char byte, pool string) int {
	m := make(map[byte]int, 0)
	for _, b := range pool {
		m[byte(b)]++
	}
	if b, ok := m[char]; ok && b <= max && b >= min {
		return 1
	}
	return 0
}
