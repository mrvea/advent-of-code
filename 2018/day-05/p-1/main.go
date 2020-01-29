package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type Node struct {
	Value rune
	Next  *Node
}

func newNode(r rune, n *Node) *Node {
	return &Node{r, n}
}

type Stack struct {
	Len  uint
	Head *Node
}

func (s *Stack) Push(r rune) {
	s.Head = newNode(r, s.Head)
	s.Len++
}

func (s *Stack) Pop() rune {
	n := s.Head
	s.Head = n.Next
	n.Next = nil
	s.Len--
	return n.Value
}

func (s *Stack) Top() rune {
	return s.Head.Value
}

func (s *Stack) Empty() bool {
	return s.Head == nil
}
func newStack() *Stack {
	return &Stack{}
}

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	rr := bufio.NewReader(f)
	stack := newStack()
	for {
		r, _, err := rr.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("could not read rune, %s", err)
		}
		if !stack.Empty() && pair(stack.Top(), r) {
			stack.Pop()
			continue
		}
		stack.Push(r)
	}
	fmt.Printf("size: %d\n", stack.Len)
}

func pair(a, b rune) bool {
	if a == b {
		return false
	}
	const dist = 32
	if abs(int(a)-int(b)) == dist {
		return true
	}
	return false
}

func abs(v int) int {
	if v < 0 {
		v *= -1
	}
	return v
}
