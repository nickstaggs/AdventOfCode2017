package main

import (
	"bufio"
	"fmt"
	"os"
)

const size int = 256

func main() {
	f, err := os.Open("lengths.txt")
	if err != nil {
		panic(err)
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

	byteArr := []byte(line)
	addedBytes := []byte{17, 31, 73, 47, 23}
	byteArr = append(byteArr, addedBytes...)

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i
	}

	curr := 0
	skip := 0
	for i := 0; i < 64; i++ {
		for _, l := range byteArr {
			end := curr + int(l) - 1
			for j := curr; j < end; j++ {
				arr[j%size], arr[end%size] = arr[end%size], arr[j%size]
				end--
			}

			curr = (curr + skip + int(l)) % size
			skip++
		}
	}

	var denseHash []int
	var denseHashHex []string
	for i := 0; i < 16; i++ {
		num := arr[i*16]
		for j := 1; j < 16; j++ {
			num = num ^ arr[i*16+j]
		}
		denseHash = append(denseHash, num)
		denseHashHex = append(denseHashHex, fmt.Sprintf("%x", num))
	}

	// 0s wont be printed if leading aka int 9 will yield the hex "9" not "09"
	fmt.Println(denseHash)
	fmt.Println(denseHashHex)
}
