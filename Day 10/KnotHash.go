package main

import (
	"bufio"
	"fmt"
	"os"
  "strings"
  "strconv"
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

  sArr := strings.Split(string(line), ",")
  lArr := make([]int, len(sArr))

  for i, str := range sArr {
    lArr[i], err = strconv.Atoi(str)
    if err != nil {
      panic(err)
    }
  }

  arr := make([]int, size)
  for i := 0; i < size; i++ {
    arr[i] = i
  }

  curr := 0
  for skip, l := range lArr {
    end := curr+l-1
    fmt.Println(curr, end)
    for j := curr; j < end; j++ {
      arr[j%size], arr[end%size] = arr[end%size], arr[j%size]
      end--
    }

    curr = (curr + skip + l)%size
  }

  fmt.Println("The multiplication of the first two numbers equals", arr[0]*arr[1])
  fmt.Println(arr)
}
