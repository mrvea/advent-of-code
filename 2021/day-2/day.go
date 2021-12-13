package day

import (
	p1 "advent-of-code/2021/day-2/p-1"
	p2 "advent-of-code/2021/day-2/p-2"
	"advent-of-code/global"
	"fmt"
)

var (
	subPackages = map[string]global.Runnable{
		p1.Path + "/p-1": p1.SubMain,
		p2.Path + "/p-2": p2.SubMain,
	}
)

func init() {
	fmt.Println("2021 called", subPackages)
	for name, action := range subPackages {
		global.Register(name, action)
	}
}
