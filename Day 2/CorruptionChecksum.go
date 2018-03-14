package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
  "strings"
)

func main() {

  f, err := os.Open("numbers.txt")
	if err != nil {
		fmt.Println("Panicking...")
		panic(1)
	}

  r := bufio.NewReader(f)

  var lines [][]int
  numLines := 0

  line, isPrefix, err := r.ReadLine()

  if isPrefix {
    panic("isPrefix")
  }

  for ; err == nil; line, isPrefix, err = r.ReadLine() {
    strLine := string(line)
    strNums := strings.Split(strLine, "\t")

    lines = append(lines, []int{})

    for i := 0; i < len(strNums); i++ {
      n, strErr := strconv.Atoi(strNums[i])

      if strErr != nil {
        panic(strErr)
      }

      lines[numLines] = append(lines[numLines], n)
    }

    numLines++
  }

  sum := 0

  for i := 0; i < len(lines); i++ {
    low := lines[i][0]
    high := lines[i][0]

    for j := 0; j < len(lines[i]); j++ {
      if lines[i][j] > high {
        high = lines[i][j]
      }
      if lines[i][j] < low {
        low = lines[i][j]
      }
    }

    sum += high - low
  }

  fmt.Println("pt 1 sum is: ", sum)
  sum = 0

  for i := 0; i < len(lines); i++ {
    tFlag := false
    for j := 0; j < len(lines[i]); j++ {
      for k := 0; k < len(lines[i]); k++ {

        if lines[i][j]%lines[i][k] == 0 && j != k {
          sum += lines[i][j]/lines[i][k]
          tFlag = true
          break
        }

        if tFlag {
          break
        }
      }
    }
  }

  fmt.Println("pt 2 sum is: ", sum)
}
