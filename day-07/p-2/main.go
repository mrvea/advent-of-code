package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

type node struct {
	Value rune
	Next  *node
}

type queue struct {
	size uint
	Head *node
	Tail *node
}

func newqueue() *queue {
	return &queue{}
}

func (l *queue) push(r rune) {
	if !l.empty() && r > l.Head.Value {
		n := l.pop()
		l.push(r)
		l.push(n)
		return
	}
	node := &node{r, l.Head}
	l.Head = node
	l.size++
}

func (l *queue) pop() rune {
	n := l.Head
	l.Head, n.Next = n.Next, nil
	l.size--
	return n.Value
}
func (l *queue) len() uint {
	return l.size
}
func (l *queue) top() rune {
	return l.Head.Value
}
func (l *queue) empty() bool {
	return l.Head == nil
}

func (l *queue) print() {
	n := l.Head
	for n != nil {
		fmt.Printf("%s(%d), ", string(n.Value), n.Value)
		n = n.Next
	}
	fmt.Println()
}

func main() {
	f, err := os.Open("test_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	m := make(map[rune]map[rune]bool)
	deps := make(map[rune]map[rune]bool)
	for s.Scan() {
		var first, second string
		fmt.Sscanf(s.Text(), "Step %s must be finished before step %s can begin.", &first, &second)
		rf, rs := rune(first[0]), rune(second[0])
		if _, ok := m[rf]; !ok {
			m[rf] = map[rune]bool{}
		}
		if _, ok := deps[rs]; !ok {
			deps[rs] = map[rune]bool{}
		}
		m[rf][rs] = true
		deps[rs][rf] = true
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	q := newqueue()
	released := map[rune]bool{}
	for s := range m {
		if _, ok := deps[s]; !ok {
			q.push(s)
		}
	}

	var b bytes.Buffer
	curR := q.pop()
	of := int('A')
	offset := 0
	workers := [2]int{}
	sec := 0
	for {
		b.WriteRune(curR)
		sec += (int(curR) - of) + 1 + offset
		fmt.Printf("curR: %s, sec: %d\n", string(curR), sec)
		released[curR] = true
		i := 0
		for key := range m[curR] {
			if i > 0 && len(workers)-1 == i {
				val := workers[i]
				if val == 0 || val < sec {
					fmt.Printf("worker: %d, sec: %d\n", i, sec)
					workers[i] = sec + (int(key) - of) + 1 + offset
					released[key] = true
					continue
				}
			}
			i++
			if isReleased(deps[key], released) {
				q.push(key)
			}
		}

		if q.empty() {
			break
		}
		curR = q.pop()
	}
	fmt.Println(b.String())
	fmt.Println(sec)
}

func isReleased(deps map[rune]bool, released map[rune]bool) bool {
	if len(deps) == 0 {
		return true
	}
	for key := range deps {

		if !released[key] {
			return false
		}
		delete(deps, key)
	}
	return true
}
