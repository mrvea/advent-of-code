package vacation42

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

const (
	day           = "4"
	basePath      = "2020/day-" + day + "/"
	inputName     = "input.txt"
	testInputName = "test_" + inputName
)

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

	passport := ""
	total := 0
	for s.Scan() {
		row := s.Text()
		if row == "" {
			fmt.Printf("passport: %s\n", passport)
			total += isValid(passport)
			fmt.Print("\n\n\n")
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

func processPassport(raw string) map[string]string {
	slice := make(map[string]string, 0)
	parts := strings.Fields(raw)
	fmt.Println("parts: ", parts)
	for _, p := range parts {
		subs := strings.Split(p, ":")
		slice[subs[0]] = subs[1]
	}
	fmt.Println(slice)
	return slice
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

	fields := map[string]func(v string) bool{
		"byr": byr,
		"iyr": iyr,
		"eyr": eyr,
		"hgt": hgt,
		"hcl": hcl,
		"ecl": ecl,
		"pid": pid,
		// "cid",
	}
	for key := range fields {
		if !strings.Contains(raw, key+":") {
			fmt.Println("missing field => ", key)
			return 0
		}
	}
	fmt.Println("all fields are present....")
	parsed := processPassport(raw)

	for key, value := range parsed {
		if key == "cid" {
			continue
		}
		a, ok := fields[key]
		if !ok || !a(value) {
			fmt.Println("not ok field => ", key)
			return 0
		}
	}

	return 1
}

func maxMin(max, min, length int) func(v string) bool {
	return func(v string) bool {
		fmt.Printf("max: %d, min: %d, len: %d, str: %s\n", max, min, length, v)
		if length > 0 && len(v) != length {
			return false
		}

		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("...........invalid number......... ", v)
			return false
		}

		if num < min || num > max {
			return false
		}
		fmt.Println("valid")
		return true
	}
}

//four digits; at least 1920 and at most 2002
func byr(v string) bool {
	length, min, max := 4, 1920, 2002

	return maxMin(max, min, length)(v)
}

//four digits; at least 2010 and at most 2020
func iyr(v string) bool {
	length, min, max := 4, 2010, 2020

	return maxMin(max, min, length)(v)
}

//four digits; at least 2020 and at most 2030
func eyr(v string) bool {
	max, min, length := 2030, 2020, 4
	return maxMin(max, min, length)(v)
}

//a number followed by either cm or in
//If cm, the number must be at least 150 and at most 193
//If in, the number must be at least 59 and at most 76
func hgt(v string) bool {
	fmt.Printf("hgt: %s => ", v)
	rawNum := ""
	unit := ""
	for _, char := range v {
		if unicode.IsDigit(char) {
			rawNum += string(char)
			continue
		}
		unit += string(char)
	}
	switch unit {
	case "in":
		return maxMin(76, 59, 0)(rawNum)
	case "cm":
		return maxMin(193, 150, 0)(rawNum)
	default:
		fmt.Printf("unknown height unit: %s, exiting...\n", unit)
		return false
	}

	// return true
}

//a # followed by exactly six characters 0-9 or a-f.
func hcl(v string) bool {
	fmt.Printf("hcl: %s => ", v)
	if v[0] != '#' {
		return false
	}

	if len(v)-1 != 6 {
		return false
	}
	for _, char := range []byte(v)[1:] {
		matched, err := regexp.Match(`[a-f0-9]`, []byte{char})
		if err != nil || !matched {
			return false
		}
	}

	fmt.Println("valid")
	return true
}

//exactly one of: amb blu brn gry grn hzl oth
func ecl(v string) bool {
	fmt.Printf("ecl: %s => ", v)
	pool := map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}
	if _, ok := pool[v]; !ok {
		return false
	}

	fmt.Println("valid")
	return true
}

// a nine-digit number, including leading zeroes
func pid(v string) bool {
	fmt.Printf("pid: %s => ", v)
	if len(v) != 9 {
		return false
	}
	matched, err := regexp.Match("[0-9]", []byte(v))
	if err != nil || !matched {
		return false
	}

	fmt.Println("valid")
	return true
}
