package vacation41

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	day           = "4"
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

	passport := ""
	total := 0
	for s.Scan() {
		row := s.Text()
		if row == "" {
			fmt.Printf("passport: %s\n", passport)
			total += isValid(passport)
			passport = ""
			continue
		}
		passport += " " + row
		// grid = append(grid, s.Text())
	}
	fmt.Printf("passport: %s\n", passport)
	total += isValid(passport)
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("total: ", total)
}

func processPassport(raw string) {

}

func isValid(raw string) int {
	//byr (Birth Year)
	//iyr (Issue Year)
	//eyr (Expiration Year)
	//hgt (Height)
	//hcl (Hair Color)
	//ecl (Eye Color)
	//pid (Passport ID)
	//cid (Country ID)
	fields := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
		// "cid",
	}
	for _, field := range fields {
		if !strings.Contains(raw, field+":") {
			return 0
		}
	}
	return 1
}
