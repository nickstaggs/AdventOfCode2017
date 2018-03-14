package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Element stuff
type Element struct {
	i    int
	next *Element
}

func main() {

	f, err := os.Open("number.txt")
	if err != nil {
		fmt.Println("Panicking...")
		panic(1)
	}

	r := bufio.NewReader(f)

	n, err := r.ReadByte()

	if err != nil {
		panic(1)
	}

	p, err := strconv.Atoi(string(n))

	head := Element{p, nil}
	last := &head
	sum := 0
	numElements := 1

	var e byte
	e, err = r.ReadByte()

	for ; err == nil; e, err = r.ReadByte() {
		p, err1 := strconv.Atoi(string(e))
		if err1 != nil {
			break
		}

		curr := Element{p, nil}
		last.next = &curr

		if last.i == curr.i {
			sum += curr.i
		}
		numElements++
		last = &curr
	}

	last.next = &head

	curr := head

	if curr.i == last.i {
		sum += curr.i
	}

	fmt.Println("sum is: ", sum)
	fmt.Println("number of elements is: ", numElements)

	sum = 0

	for i := 0; i < numElements/2; i++ {
		curr = *curr.next
	}

	curr2 := head

	for i := 0; i < numElements; i++ {
		if curr.i == curr2.i {
			sum += curr.i
		}

		curr = *curr.next
		curr2 = *curr2.next
	}

	curr = *curr.next
	curr2 = *curr2.next

	if curr.i == curr2.i {
		sum += curr.i
	}

	fmt.Println("sum is: ", sum)
}
