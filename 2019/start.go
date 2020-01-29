package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

// MatrixNode node with value that could have 2 or more connections.
type MatrixNode struct {
	Pos    *XY
	Value  *XY
	Top    *MatrixNode
	Bottom *MatrixNode
	Left   *MatrixNode
	Right  *MatrixNode
}

func (N *MatrixNode) nextTop() *MatrixNode {
	n := newMNode(N.Pos.withOffset(0, -1))
	n.toTopBottom(N)
	N.Left.Top.toRightLeft(n)
	return n
}

func (N *MatrixNode) nextLeft() *MatrixNode {
	n := newMNode(N.Pos.withOffset(-1))
	n.toRightLeft(N)
	N.Top.Left.toTopBottom(n)
	return n
}

func (N *MatrixNode) nextRight() *MatrixNode {
	n := newMNode(N.Pos.withOffset(1))
	N.toRightLeft(n)
	n.toTopBottom(N.Bottom.Right)
	return n
}

func (N *MatrixNode) nextBottom() *MatrixNode {
	n := newMNode(N.Pos.withOffset(0, 1))
	N.toTopBottom(n)
	n.toRightLeft(N.Right.Bottom)
	return n
}

func (N *MatrixNode) toTopBottom(node *MatrixNode) {
	N.Bottom = node
	node.Top = N
}
func (N *MatrixNode) toRightLeft(node *MatrixNode) {
	N.Right = node
	node.Left = N
}

func (N *MatrixNode) String() string {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("pos: %s, ", N.Pos.String()))
	v := "nil"
	if N.Value != nil {
		v = N.Value.String()
	}
	b.WriteString(fmt.Sprintf("Value: %s => ", v))
	if N.Top != nil {
		b.WriteString(fmt.Sprintf("top: %s", N.Top.Pos))
	}
	if N.Bottom != nil {
		b.WriteString(fmt.Sprintf("bottom: %s", N.Bottom.Pos))
	}
	if N.Left != nil {
		b.WriteString(fmt.Sprintf("left: %s", N.Left.Pos))
	}
	if N.Right != nil {
		b.WriteString(fmt.Sprintf("right: %s\n\n", N.Right.Pos))
	}
	return b.String()
}

func newMNode(p *XY, nodes ...*MatrixNode) *MatrixNode {
	mn := &MatrixNode{Pos: p}
	if len(nodes) > 0 {
		for i, n := range nodes {
			switch i {
			case 0:
				mn.Top = n
			case 1:
				mn.Bottom = n
			case 2:
				mn.Left = n
			case 3:
				mn.Right = n
			default:
				fmt.Printf("not supported number of nodes %d", i)
			}
		}
	}
	return mn
}

// RippleMatrix is a matrix which full consists of nodes, which could be connected up to 4 directions
// this matrix starts from center node and 'ripple' or radiate to every immediate node around it.
// there is head and tail of the matrix.
// iteration could be done cols to rows, or rows to cols or zigzag...(if need be)
type RippleMatrix struct {
	Head *MatrixNode
	Tail *MatrixNode
	Size int
}

// NextRipple adds and return all new node around the current matrix.
func (M *RippleMatrix) NextRipple() []*MatrixNode {
	if M.Size < 3 {
		M.Size = 3
		return []*MatrixNode{M.Head}
	}
	points := make([]*MatrixNode, 0, M.Size*2+(M.Size-2))
	_, topRow, leftCol := M.headCorner()
	_, bottomRow, rightCol := M.tailCorner()
	points = append(points, M.Head, M.Tail, leftCol, topRow, bottomRow, rightCol)
	for i := 2; i < M.Size; i++ {
		if i < M.Size-1 {
			leftCol = leftCol.nextBottom()
			rightCol = rightCol.nextTop()
			points = append(points, leftCol, rightCol)
		}
		topRow = topRow.nextRight()
		bottomRow = bottomRow.nextLeft()
		points = append(points, topRow, bottomRow)
	}
	M.Size += 2
	return points
}
func (M *RippleMatrix) headCorner() (head, right, bottom *MatrixNode) {
	head, right, bottom = M.leads(M.Head, -1)
	head.toRightLeft(right)
	head.toTopBottom(bottom)
	bottom.toRightLeft(M.Head)
	right.toTopBottom(M.Head)
	M.Head = head
	return head, right, bottom
}

func (M *RippleMatrix) tailCorner() (tail, left, top *MatrixNode) {
	tail, left, top = M.leads(M.Tail, 1)
	left.toRightLeft(tail)
	top.toTopBottom(tail)
	M.Tail.toRightLeft(top)
	M.Tail.toTopBottom(left)
	M.Tail = tail
	return tail, left, top
}

func (M *RippleMatrix) leads(from *MatrixNode, offset int) (start, x, y *MatrixNode) {
	start = newMNode(from.Pos.withOffset(offset, offset))
	x = newMNode(start.Pos.withOffset(-offset))
	y = newMNode(start.Pos.withOffset(0, -offset))
	return start, x, y
}

func (M *RippleMatrix) String() string {
	var b bytes.Buffer
	row := M.Head
	for row != nil {
		col := row
		for col != nil {
			str := ".."
			if col.Value != nil {
				str = fmt.Sprintf("%d%d", col.Value.x, col.Value.y)
			}
			b.WriteString(str)
			b.WriteRune(' ')
			col = col.Right
		}
		b.WriteRune('\n')
		row = row.Bottom
	}
	return b.String()
}

func newMatrix(midXY *XY) *RippleMatrix {
	n := newMNode(midXY)
	return &RippleMatrix{n, n, 1}
}

// XY coor struct to hold the xy of possition and value(owner in this instance)
type XY struct{ x, y int }

func (xy *XY) withOffset(offsets ...int) *XY {
	x, y := xy.x, xy.y
	for i, n := range offsets {
		switch i {
		case 0:
			x += n
		case 1:
			y += n
		default:
			fmt.Printf("unsupported number of offests: %d", i)
			break
		}
	}
	return newXY(x, y)
}

func (xy *XY) String() string {
	return fmt.Sprintf("(%d, %d)", xy.x, xy.y)
}

func newXY(x, y int) *XY {
	return &XY{x, y}
}

func main() {
	f, err := os.Open("../test_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	points := make([]*XY, 0)
	maxXY := newXY(0, 0)
	for s.Scan() {
		var x, y int
		fmt.Sscanf(s.Text(), "%d, %d", &x, &y)
		p := newXY(x, y)
		points = append(points, p)
		if maxXY.x < p.x {
			maxXY.x = p.x
		}
		if maxXY.y < p.y {
			maxXY.y = p.y
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	_ = runRipple(points, maxXY)

	// fmt.Println(mat)
}

func runRipple(points []*XY, maxXY *XY) *RippleMatrix {

	midXY := newXY(maxXY.x/2, maxXY.y/2)
	mat := newMatrix(midXY)
	max := max(maxXY.x, maxXY.y) + 2
	areas := map[*XY]int{}
	change := map[*XY]int{}
	stopped := map[*XY]bool{}
	for {
		for _, p := range mat.NextRipple() {
			p.Value = spaceOwner(points, p.Pos)
			if p.Value != nil {
				change[p.Value]++
			}
		}
		for _, p := range points {
			val, ok := change[p]
			fmt.Printf("changed: %s(%d)\n", p, val)
			if !ok || val == 0 {
				stopped[p] = true
				continue
			}
			change[p] = 0
			areas[p] += val
			if s, ok := stopped[p]; ok && s {
				delete(stopped, p)
			}
		}
		// fmt.Println(mat)
		if mat.Size > max {
			break
		}
		fmt.Println()
	}
	var maxArea int
	var maxPoint *XY
	for p := range stopped {
		if a := areas[p]; a > maxArea {
			maxArea = a
			maxPoint = p
		}
	}
	fmt.Printf("max bounded %s => Area: %d\n", maxPoint, maxArea)
	return mat
}

func spaceOwner(points []*XY, a *XY) *XY {
	var owner *XY
	minDist := ^uint(0)
	for _, p := range points {
		dist := manhattanDistance(a, p)
		if dist < minDist {
			minDist = dist
			owner = p
			continue
		}
		if dist == minDist {
			owner = nil
		}
	}

	return owner
}

func manhattenDistanceSum(points []*XY, a *XY) uint {
	var total uint
	for _, v := range points {
		total += manhattanDistance(a, v)
	}
	return total
}

func manhattanDistance(a, b *XY) uint {
	return uint(abs(a.x-b.x) + abs(a.y-b.y))
}

func abs(v int) int {
	y := v >> 7
	return (v ^ y) - y
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
