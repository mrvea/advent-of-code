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
	bags := make(map[string]map[string]int)
	sbagName := "shiny gold"
	for s.Scan() {
		fmt.Println(s.Text())
		bagName, contains := processDirections(s.Text())
		// fmt.Printf("bag: %s, contains: ", bagName)
		// fmt.Println(contains)
		if contains == nil || bagName == sbagName {
			continue
		}
		bags[bagName] = contains
	}
	// total += len(bags)
	m := map[string]struct{}{"shiny gold": struct{}{}}
	newM := map[string]struct{}{}
	fmt.Println("bags: ", bags)

	for {
		if len(m) == 0 {
			break
		}
		// var name string
		// var bag map[string]int
		for name, bag := range bags {
			fmt.Printf("checking in %s => ", name)
			fmt.Println(m)
		INNER:
			for v := range m {
				fmt.Printf("%s checking in %s\n", v, name)
				// fmt.Println(bag[v])
				if _, ok := bag[v]; ok {
					fmt.Printf("%s in %s\n", v, name)
					total++
					delete(bags, name)
					newM[name] = struct{}{}
					break INNER
				}
			}
		}
		// if bag == nil {
		// 	delete(bags, name);
		// 	continue;
		// }

		m = newM
		newM = make(map[string]struct{}, 0)
	}

	fmt.Println("total: ", total)
}

// func makeAssociation(bags map[string][]string) {

// }

func isShinyGoldInRaw(pool string) int {

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
