package vacation101

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

const (
	day           = "10"
	basePath      = "2020/day-" + day + "/"
	inputName     = "input.txt"
	testInputName = "test_" + inputName
)

/**
	[16,10,15,5,1,11,7,19,6,12,4]
	[1, 4, 5, 6, 7, 10, 11, 12, 15, 16, 19]
	The charging outlet has an effective rating of 0 jolts,
		so the only adapters that could connect to it directly would need to have a joltage rating of 1, 2, or 3 jolts. Of these, only one you have is an adapter rated 1 jolt (difference of 1).
		0 => [1, 2, 3] => 1(1)
	From your 1-jolt rated adapter, the only choice is your 4-jolt rated adapter (difference of 3).
		1 => [2, 3, 4] => 4(3)
	From the 4-jolt rated adapter, the adapters rated 5, 6, or 7 are valid choices. However, in order to not skip any adapters, you have to pick the adapter rated 5 jolts (difference of 1).
		4 => [5, 6, 7] => 5(1)
	Similarly, the next choices would need to be the adapter rated 6 and then the adapter rated 7 (with difference of 1 and 1).
		5 => [6, 7, 8] => 6(1) => 7(1)
	The only adapter that works with the 7-jolt rated adapter is the one rated 10 jolts (difference of 3).
		7 => [8, 9, 10] => 10(3)
	From 10, the choices are 11 or 12; choose 11 (difference of 1) and then 12 (difference of 1).
		10 => [11, 12, 13] =>11(1) => 12(1)
	After 12, only valid adapter has a rating of 15 (difference of 3), then 16 (difference of 1), then 19 (difference of 3).
		12 => [13, 14, 15] => 15(3)
		15 => [16, 17, 18] => 16(1)
		16 => [17, 18, 19] => 19(3)
	Finally, your device's built-in adapter is always 3 higher than the highest adapter, so its rating is 22 jolts (always a difference of 3).

**/
func SubMain(args ...string) {
	fmt.Println("vacation" + day + ".1 starting exectution...")
	filePath := basePath + inputName
	if len(args) > 0 && args[0] == "test" {
		fmt.Println("Running test file...")
		filePath = basePath + testInputName
		if len(args) > 1 {
			filePath = basePath + args[1]
		}
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
	adapters := make([]int, 0)
	// highest := 0
	for s.Scan() {
		fmt.Println(s.Text())
		j, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Panic("value could not be converted to a number", err)
		}
		// if highest < j {
		// 	highest = j
		// }
		adapters = append(adapters, j)
	}

	total = getChainMul(adapters)
	fmt.Println("total: ", total)
}

func getChainMul(a []int) int {
	var base, ones, threes int
	sort.Ints(a)
	fmt.Println("sorted: ", a)
	for _, v := range a {
		// fmt.Printf("\tbase: %d, value: %d, dif: %d", base, v, v-base)
		switch v - base {
		case 1:
			ones++
		case 2:
		case 3:
			threes++
		default:
			log.Panicf("unexpected range: %d", v-base)
		}
		// fmt.Printf(" 1-jolt: %d, 3-jolt: %d\n", ones, threes)
		base = v

	}
	threes++
	fmt.Printf("1-jolt: %d, 3-jolt: %d\n", ones, threes)
	return ones * threes
}
