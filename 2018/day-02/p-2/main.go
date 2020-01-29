package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	rr := bufio.NewReader(f)
	m := map[int]map[rune][]*bytes.Buffer{}
	words := map[int]*bytes.Buffer{0: &bytes.Buffer{}}
	matched := map[string]int{}

	i := 0
	j := 0
	fmt.Println("Started: ")
	for {
		r, _, err := rr.ReadRune()
		if err != nil && err != io.EOF {
			log.Fatalf("cound not read rune from file: %s", err)
		}
		if r == ' ' {
			continue
		}
		b := words[j]

		if r == '\n' || err == io.EOF {
			size := b.Len()
			for s, c := range matched {
				// fmt.Printf("current str: %s\nmatch str: %s\nsize: %d\ncount: %d\n\n", words[j].String(), s, len(s), c)
				if size-c == 1 {
					log.Fatalf("found the ans: %s\n", excludeUncommon(s, words[j].String()))
				}
			}
			if err == io.EOF {
				break
			}
			i = 0
			j++
			matched = map[string]int{}
			words[j] = &bytes.Buffer{}
			continue
		}

		b.WriteRune(r)
		runesAtIndex, ok := m[i]
		if !ok {
			runesAtIndex = map[rune][]*bytes.Buffer{}
		}

		for _, s := range runesAtIndex[r] {
			matched[s.String()]++
		}
		runesAtIndex[r] = append(runesAtIndex[r], b)
		words[j] = b
		m[i] = runesAtIndex
		i++
	}
}

func excludeUncommon(a, b string) string {
	size := len(a)
	var buf bytes.Buffer
	for i := 0; i < size; i++ {
		if a[i] == b[i] {
			buf.WriteByte(a[i])
		}
	}
	return buf.String()
}
