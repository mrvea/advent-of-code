package vacation71

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	day           = "7"
	basePath      = "2020/day-" + day + "/"
	inputName     = "input.txt"
	testInputName = "test_" + inputName
	rowMax        = 127.0
	rowMin        = 0.0
	columnMax     = 7.0
	columnMin     = 0.0
)

/**
	light red bags contain 1 bright white bag, 2 muted yellow bags.
	dark orange bags contain 3 bright white bags, 4 muted yellow bags.
	bright white bags contain 1 shiny gold bag.
	muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
	shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
	dark olive bags contain 3 faded blue bags, 4 dotted black bags.
	vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
	faded blue bags contain no other bags.
	dotted black bags contain no other bags.
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
		return
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	total := 0
	ans := make(map[rune]struct{})
	for s.Scan() {
		fmt.Println(s.Text())
		bagName, contains := processDirections(s.Text())

		fmt.Printf("bag: %s", bagName)
		fmt.Println(contains)
	}
	total += len(ans)

	fmt.Println("total: ", total)
}

func isShinyGoldIn(raw string) int {

	return 0
}

func processDirections(raw string) (string, map[string]int) {
	var key, containRaw string
	parts := strings.Split(raw, " bags contain ")
	key, containRaw = parts[0], parts[1]
	// var shade, color string
	// _, err := fmt.Sscanf(key, "%s %s bags", &shade, &color)
	// if err != nil {
	// 	log.Panic(err)
	// }
	fmt.Printf("key: %s, contains: %s\n", key, containRaw)
	if containRaw == "no other bags." {
		return key, nil
	}

	parts = strings.Split(containRaw, ", ")
	containsMap := make(map[string]int)
	for _, str := range parts {
		var shade, color string
		var count int
		fmt.Printf("part: %s\n", str)
		_, err := fmt.Sscanf(str, "%d %s %s bag", &count, &shade, &color)
		if err != nil {
			log.Panic(err)
		}
		containsMap[shade+" "+color] = count
	}

	return key, containsMap
}
