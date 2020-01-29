package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

const (
	format = "2006-01-02 15:04"
)

type Node struct {
	Value *LogItem
	Next  *Node
	Prev  *Node
}

func newNode(li *LogItem, nodes ...*Node) *Node {
	n := &Node{Value: li}
	for i, t := range nodes {
		switch i {
		case 0:
			n.Next = t
		case 1:
			n.Prev = t
		default:
			log.Fatalln("wront number of nodes for newNode func, max is 2")
		}
	}
	return n
}

type LogItem struct {
	T         time.Time
	Type      string
	Substring string
}

func (li *LogItem) Less(b *LogItem) bool {
	return li.T.Before(b.T)
}

func (li *LogItem) String() string {

	return fmt.Sprintf("[%s] %s %s", li.T.Format(format), li.Type, li.Substring)
}

func newLogItem(t time.Time, typeStr, Substr string) *LogItem {

	return &LogItem{
		t,
		typeStr,
		Substr,
	}
}

type Log struct {
	Head *Node
	Tail *Node
	Len  uint
}

func (l *Log) Push(li *LogItem) {
	n := newNode(li)
	l.Len++
	if l.Empty() {
		l.Head = n
		l.Tail = n
		return
	}
	if li.Less(l.Bottom()) {
		t := l.PopBottom()
		l.Push(li)
		l.Push(t)
		return
	}
	n.Prev, l.Tail.Next, l.Tail = l.Tail, n, n
}

func (l *Log) Top() *LogItem {
	if l.Empty() {
		return nil
	}
	return l.Head.Value
}
func (l *Log) Bottom() *LogItem {
	if l.Tail == nil {
		return nil
	}
	return l.Tail.Value
}

func (l *Log) Pop() *LogItem {
	n := l.Head
	l.Len--
	l.Head, l.Head.Prev, n.Next = n.Next, nil, nil
	if l.Empty() {
		l.Tail = nil
	}
	return n.Value
}
func (l *Log) PopBottom() *LogItem {
	n := l.Tail
	l.Len--
	l.Tail, l.Tail.Next, n.Prev = n.Prev, nil, nil
	if l.Tail == nil {
		l.Head = nil
	}
	return n.Value
}

func (l *Log) Empty() bool {
	return l.Head == nil
}
func (l *Log) String() string {
	var b bytes.Buffer
	h := l.Head
	if h == nil {
		return "empty"
	}
	for h != nil {
		fmt.Println(h.Value)
		b.WriteString(fmt.Sprintln(h.Value))
		h = h.Next
	}
	return b.String()
}

func newLog() *Log {
	return &Log{}
}

type Sleep struct {
	From time.Time
	To   time.Time
}

func (s Sleep) Duration() time.Duration {
	return s.To.Sub(s.From)
}

func newSleep(from time.Time) Sleep {
	return Sleep{From: from}
}

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	l := newLog()
	for s.Scan() {
		var dateStr, typeStr, substr string
		err = Sfieldf(s.Text(), &dateStr, &typeStr, &substr)
		if err != nil {
			log.Fatalln(err)
		}
		l.Push(newLogItem(DateExec(dateStr), typeStr, substr))
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	asleeps := map[string][]Sleep{}
	// fmt.Println(l)
	var guardID string
	var asleep Sleep
	for !l.Empty() {
		li := l.Pop()
		switch li.Type {
		case "Guard":
			guardID = li.Substring
		case "falls":
			asleep = newSleep(li.T)
		case "wakes":
			asleep.To = li.T
			sleeps, ok := asleeps[guardID]
			if !ok {
				sleeps = make([]Sleep, 0)
			}
			sleeps = append(sleeps, asleep)
			asleeps[guardID] = sleeps
		default:
			log.Fatalf("unsupported log item time %s", li.T)
		}
	}

	maxMinute := -1
	maxCount := 0
	for id, g := range asleeps {
		r := make([]int, 60)
		for _, s := range g {
			for i := s.From.Minute(); i < s.To.Minute(); i++ {
				r[i]++
				if r[i] > maxCount {
					guardID = id
					maxCount = r[i]
					maxMinute = i
				}
			}
		}
	}
	fmt.Printf("max minute %d\n", maxMinute)

	id, err := strconv.Atoi(guardID)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("ans: %d\n", id*maxMinute)
}

//Sfieldf will separate string with specific Splitter to push them into string variables
func Sfieldf(str string, vars ...*string) error {
	values := strings.FieldsFunc(str, Splitter())
	if len(vars) > len(values) {
		return fmt.Errorf("more vars than values")
	}
	for i, v := range vars {
		*v = values[i]
	}
	return nil
}

//Splitter return a anonymous function which will be used in fieldFunc to separate string separate words or word sequence surrounded by square brackets
func Splitter() func(r rune) bool {
	var keep bool
	return func(r rune) bool {
		if r == '[' || r == ']' {
			keep = !keep
			if keep {
				return keep
			}
		}
		return !keep && !unicode.IsNumber(r) && !unicode.IsLetter(r)
	}
}

//DateExec converts date to one format; and it has to execute or it will fail
func DateExec(dateStr string) time.Time {
	t, err := time.Parse(format, dateStr)
	if err != nil {
		log.Fatalln(err)
	}
	return t
}
