package main

import (
	"advent-of-code/global"
	"fmt"
	"os"

	vacation11 "advent-of-code/2020/day-1/p-1"
	vacation12 "advent-of-code/2020/day-1/p-2"
	vacation101 "advent-of-code/2020/day-10/p-1"
	vacation102 "advent-of-code/2020/day-10/p-2"
	vacation111 "advent-of-code/2020/day-11/p-1"
	vacation112 "advent-of-code/2020/day-11/p-2"
	vacation21 "advent-of-code/2020/day-2/p-1"
	vacation22 "advent-of-code/2020/day-2/p-2"
	vacation31 "advent-of-code/2020/day-3/p-1"
	vacation32 "advent-of-code/2020/day-3/p-2"
	vacation41 "advent-of-code/2020/day-4/p-1"
	vacation42 "advent-of-code/2020/day-4/p-2"
	vacation51 "advent-of-code/2020/day-5/p-1"
	vacation52 "advent-of-code/2020/day-5/p-2"
	vacation61 "advent-of-code/2020/day-6/p-1"
	vacation62 "advent-of-code/2020/day-6/p-2"
	vacation71 "advent-of-code/2020/day-7/p-1"
	vacation72 "advent-of-code/2020/day-7/p-2"
	vacation81 "advent-of-code/2020/day-8/p-1"
	vacation82 "advent-of-code/2020/day-8/p-2"
	vacation91 "advent-of-code/2020/day-9/p-1"
	vacation92 "advent-of-code/2020/day-9/p-2"
	_ "advent-of-code/2021"
)

var (
	subPackages = map[string]global.Runnable{
		"2020/day-1/p-1":  vacation11.SubMain,
		"2020/day-1/p-2":  vacation12.SubMain,
		"2020/day-2/p-1":  vacation21.SubMain,
		"2020/day-2/p-2":  vacation22.SubMain,
		"2020/day-3/p-1":  vacation31.SubMain,
		"2020/day-3/p-2":  vacation32.SubMain,
		"2020/day-4/p-1":  vacation41.SubMain,
		"2020/day-4/p-2":  vacation42.SubMain,
		"2020/day-5/p-1":  vacation51.SubMain,
		"2020/day-5/p-2":  vacation52.SubMain,
		"2020/day-6/p-2":  vacation62.SubMain,
		"2020/day-6/p-1":  vacation61.SubMain,
		"2020/day-7/p-2":  vacation72.SubMain,
		"2020/day-7/p-1":  vacation71.SubMain,
		"2020/day-8/p-2":  vacation82.SubMain,
		"2020/day-8/p-1":  vacation81.SubMain,
		"2020/day-9/p-2":  vacation92.SubMain,
		"2020/day-9/p-1":  vacation91.SubMain,
		"2020/day-10/p-2": vacation102.SubMain,
		"2020/day-10/p-1": vacation101.SubMain,
		"2020/day-11/p-2": vacation112.SubMain,
		"2020/day-11/p-1": vacation111.SubMain,
	}
)

func init() {
	// save_christmas.Init()

	for name, action := range subPackages {
		global.Register(name, action)
	}
}

func main() {
	args := os.Args[1:]
	name := args[0]
	if name == "" {
		return
	}
	if fn, ok := global.GetAction(name); ok {
		fn(args[1:]...)
		return
	}
	fmt.Printf("cound not find %s in registered methods", name)
	// if ()
}
