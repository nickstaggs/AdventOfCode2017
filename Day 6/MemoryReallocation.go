package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type boolInt struct {
	isSeen bool
	rc     int
}

func arrToString(mb []int) string {
	var buffer bytes.Buffer
	for i := 0; i < len(mb); i++ {
		buffer.WriteString(string(mb[i]))
	}
	return buffer.String()
}

func isInMap(m map[string]boolInt, key string, rc int) bool {
	if m[key].isSeen {
		return true
	}
	bi := boolInt{true, rc}

	m[key] = bi

	return false
}

func largestIntInd(s []int) int {
	li := 0
	for i := 0; i < len(s); i++ {
		if s[i] > s[li] {
			li = i
		}
	}

	return li
}

func redistribute(s []int) {
	li := largestIntInd(s)
	n := s[li]
	s[li] = 0

	var i int
	if li == len(s)-1 {
		i = 0
	} else {
		i = li + 1
	}

	for ; n > 0; n, i = n-1, i+1 {
		s[i]++
		if i == len(s)-1 {
			i = -1
		}
	}
}

func main() {
	f, err := os.Open("memorybanks.txt")
	if err != nil {
		fmt.Println("Panicking...")
		panic(1)
	}

	r := bufio.NewReader(f)

	line, isPrefix, err := r.ReadLine()

	if isPrefix || err != nil {
		if isPrefix {
			panic("is prefix")
		} else {
			panic(err)
		}
	}

	strLine := string(line)
	strNums := strings.Split(strLine, "\t")

	mb := make([]int, 0)

	for i := 0; i < len(strNums); i++ {
		n, strErr := strconv.Atoi(strNums[i])

		if strErr != nil {
			panic(strErr)
		}

		mb = append(mb, n)
	}

	m := make(map[string]boolInt)

	rc := 0

	for !isInMap(m, arrToString(mb), rc) {
		rc++
		redistribute(mb)
	}

	fmt.Println("The number of redistribution cycles completed was", rc)
	fmt.Println("The number of redistribution cycles completed between the two equal states was", rc-m[arrToString(mb)].rc)
}
