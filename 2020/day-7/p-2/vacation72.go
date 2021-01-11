package vacation72

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
)

/**

**/
func SubMain(args ...string) {
	fmt.Println("vacation" + day + ".2 starting exectution...")
	filePath := basePath + inputName
	if len(args) > 0 && args[0] == "test" {
		fmt.Println("Running test file...")
		if len(args) > 1 {

			filePath = basePath + args[1]
		} else {

			filePath = basePath + testInputName
		}
	}
	fmt.Printf("filePath: %s\n", filePath)
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	total := 0
	bags := make(map[string]map[string]int)
	sbagName := "shiny gold"
	for s.Scan() {
		fmt.Println(s.Text())
		bagName, contains := processDirections(s.Text())
		// fmt.Printf("bag: %s, contains: ", bagName)
		// fmt.Println(contains)
		// if contains == nil || bagName == sbagName {
		// 	continue
		// }
		bags[bagName] = contains
	}
	total = containedBags(sbagName, bags, make(map[string]int))
	fmt.Println("total: ", total)
}

func containedBags(bagName string, pool map[string]map[string]int, cache map[string]int) int {
	// fmt.Println(bagName, " => ")
	bag, ok := pool[bagName]
	if !ok {
		log.Panicf("bag named %s not found", bagName)
	}
	delete(pool, bagName)
	if bag == nil {
		return 0
	}
	total := 0
	for name, value := range bag {
		if c, ok := cache[name]; ok {
			total += value + c*value
			continue
		}
		count := containedBags(name, pool, cache)
		cache[name] = count
		// if count != 0 {
		total += value + count*value
		// }
	}
	return total
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
	// fmt.Printf("key: %s, contains: %s\n", key, containRaw)
	if containRaw == "no other bags." {
		return key, nil
	}

	parts = strings.Split(containRaw, ", ")
	containsMap := make(map[string]int)
	for _, str := range parts {
		var shade, color string
		var count int
		// fmt.Printf("part: %s\n", str)
		_, err := fmt.Sscanf(str, "%d %s %s bag", &count, &shade, &color)
		if err != nil {
			log.Panic(err)
		}
		containsMap[shade+" "+color] = count
	}

	return key, containsMap
}
