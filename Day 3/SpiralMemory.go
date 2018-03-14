package main

import (
	"fmt"
	"math"
)

type Coordinate struct {
	x, y int
}

type Element struct {
	val  int
	c    Coordinate
	adj  [8]Coordinate
	prev *Element
}

func (c1 Coordinate) IsEqual(c2 Coordinate) bool {
	if c1.x == c2.x && c1.y == c2.y {
		return true
	}
	return false
}

func (e1 *Element) IsInAdj(e2 Element) bool {
	for _, c2 := range e2.adj {
		if e1.c.IsEqual(c2) {
			return true
		}
	}
	return false
}

func (e *Element) GetValue() {
	last := e.prev
	for !last.c.IsEqual(Coordinate{0, 0}) {
		if e.IsInAdj(*last) {
			e.val += last.val
		}
		last = last.prev
	}

	if e.IsInAdj(*last) {
		e.val += last.val
	}
}

func (e *Element) CreateAdj() {
	ai := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i != 0 || j != 0 {
				e.adj[ai] = Coordinate{e.c.x + i, e.c.y + j}
				ai++
			}
		}
	}
}

func (e1 Element) GetNext() Element {
	var e2 Element
	if float64(-e1.c.y) >= math.Abs(float64(e1.c.x)) {
		e2.c.x = e1.c.x + 1
		e2.c.y = e1.c.y

	} else if math.Abs(float64(e1.c.y)) <= float64(-e1.c.x) {
		e2.c.y = e1.c.y - 1
		e2.c.x = e1.c.x

	} else if float64(e1.c.y) >= float64(e1.c.x) {
		e2.c.x = e1.c.x - 1
		e2.c.y = e1.c.y

	} else {
		e2.c.y = e1.c.y + 1
		e2.c.x = e1.c.x
	}

	e2.CreateAdj()
	e2.prev = &e1

	return e2
}

func main() {
	n := 265149
	n2 := n
	square := n
	i := 1

	for ; i*i < n; i += 2 {
	}

	n = i*i - n

	sideSquare := n/i + 1

	moves := math.Sqrt(float64((i/2*sideSquare-n)*(i/2*sideSquare-n))) + float64(i/2)

	fmt.Println(moves, "steps are required to carry the data from", square, "all the way to the access port")

	var e Element
	e.c = Coordinate{0, 0}
	e.val = 1
	e.CreateAdj()

	for e.val < n2 {
		e = e.GetNext()
		e.GetValue()
	}

	fmt.Println("The first number greater than ", n2, " is ", e.val)
}
