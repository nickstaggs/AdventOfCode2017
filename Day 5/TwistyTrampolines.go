package main

import (
	"bufio"
	"fmt"
	"os"
  "strconv"
)

func main() {

  f, err := os.Open("jumps.txt")
	if err != nil {
		fmt.Println("Panicking...")
		panic(1)
	}

  r := bufio.NewReader(f)
  s := make([]int, 0)

  for l, isPrefix, err := r.ReadLine(); err == nil; l, isPrefix, err = r.ReadLine() {

    if err != nil || isPrefix {
  		fmt.Println("Panicking...")
  		panic(err)
  	}

		n, err := strconv.Atoi(string(l))
    if err != nil {
  		fmt.Println("Panicking...")
  		panic(err)
  	}

    s = append(s, n)
  }

  len := len(s)
  jmp := 0
  steps := 0

  for i := 0; i >= 0 && i < len; i += jmp {
    jmp = s[i]

    // for solution to part 2
    if jmp > 2 {
      s[i]--
    } else {
      s[i]++
    }

    // for solution to part 1
    // s[i]++

    steps++
  }

  fmt.Println("It took ", steps, " steps.")
}
