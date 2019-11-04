package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type XY struct{ x, y int }

type Claim struct {
	ID          int
	TopLeft     XY
	BottomRight XY
	W, T        int
	free        bool
}

func (c *Claim) Overlap(b *Claim) bool {

	if c.TopLeft.x > b.BottomRight.x || b.TopLeft.x > c.BottomRight.x {
		return false
	}

	if c.TopLeft.y > b.BottomRight.y || b.TopLeft.y > c.BottomRight.y {
		return false
	}

	return true
}
func (c *Claim) OverlapCoor(b *Claim) (top XY, bottom XY) {
	top.x = Max(c.TopLeft.x, b.TopLeft.x)
	top.y = Max(c.TopLeft.y, b.TopLeft.y)
	bottom.x = Min(c.BottomRight.x, b.BottomRight.x)
	bottom.y = Min(c.BottomRight.y, b.BottomRight.y)
	return
}

func (c *Claim) String() string {
	return fmt.Sprintf("ID: %d\nTopLeft: (%d, %d)\nBottomRight: (%d, %d)\nw: %d\nt: %d\n",
		c.ID,
		c.TopLeft.x, c.TopLeft.y,
		c.BottomRight.x, c.BottomRight.y,
		c.W, c.T)
}

func newClaim(id, x, y, w, t int) *Claim {
	return &Claim{
		id,
		XY{x, y},
		XY{x + w, y + t},
		w,
		t,
		true,
	}
}

type UnorderdMatrix map[int]map[int]int

func (m UnorderdMatrix) SetOverlaps(p []*Claim, n *Claim) {
	for _, c := range p {
		if !n.Overlap(c) {
			continue
		}
		if n.free {
			n.free = false
		}
		if c.free {
			c.free = false
		}
		m.SetOverlap(n.OverlapCoor(c))
	}
}

func (m UnorderdMatrix) SetOverlap(top, bottom XY) {
	for i := top.y; i < bottom.y; i++ {
		row, ok := m[i]
		if !ok {
			row = map[int]int{}
		}
		for j := top.x; j < bottom.x; j++ {
			if _, ok := row[j]; ok {
				continue
			}
			row[j]++
		}
		m[i] = row
	}
}

func (m UnorderdMatrix) OverlapArea() int {
	var area int
	for _, row := range m {
		area += len(row)
	}
	return area
}

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	claims := make([]*Claim, 0)
	m := UnorderdMatrix{}
	for s.Scan() {
		var id, x, y, w, t int
		fmt.Sscanf(s.Text(), "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &t)
		c := newClaim(id, x, y, w, t)
		m.SetOverlaps(claims, c)
		claims = append(claims, c)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(m.OverlapArea())
	for _, c := range claims {
		if c.free {
			log.Fatalf("free clamin id %d", c.ID)
		}
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
