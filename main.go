package main

import (
	"advent-of-code/global"
	"os"

	"advent-of-code/2020/day-1/p-1"
	"advent-of-code/2020/day-1/p-2"
	"advent-of-code/2020/day-2/p-1"
	"advent-of-code/2020/day-2/p-2"
	"advent-of-code/2020/day-3/p-1"
	"advent-of-code/2020/day-3/p-2"
	"advent-of-code/2020/day-4/p-1"
	"advent-of-code/2020/day-4/p-2"
	"advent-of-code/2020/day-5/p-1"
	"advent-of-code/2020/day-5/p-2"
	"advent-of-code/2020/day-6/p-1"
	"advent-of-code/2020/day-6/p-2"
	"advent-of-code/2020/day-7/p-1"
	"advent-of-code/2020/day-7/p-2"
)

var subPackages = map[string]global.Runnable{
	"2020/day-1/p-1": vacation11.SubMain,
	"2020/day-1/p-2": vacation12.SubMain,
	"2020/day-2/p-1": vacation21.SubMain,
	"2020/day-2/p-2": vacation22.SubMain,
	"2020/day-3/p-1": vacation31.SubMain,
	"2020/day-3/p-2": vacation32.SubMain,
	"2020/day-4/p-1": vacation41.SubMain,
	"2020/day-4/p-2": vacation42.SubMain,
	"2020/day-5/p-1": vacation51.SubMain,
	"2020/day-5/p-2": vacation52.SubMain,
	"2020/day-6/p-2": vacation62.SubMain,
	"2020/day-6/p-1": vacation61.SubMain,
	"2020/day-7/p-2": vacation72.SubMain,
	"2020/day-7/p-1": vacation71.SubMain,
}

func init() {
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
	}
	// if ()
}
