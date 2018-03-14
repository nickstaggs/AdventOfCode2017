package main

import (
	"bufio"
	"io"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isCondition(s []string, m map[string]int) bool {
  key := s[4]
  val, err := strconv.Atoi(s[6])

  if err != nil {
    panic(err)
  }

  switch cond := s[5]; cond {
  case "==":
    return m[key] == val
  case "!=":
    return m[key] != val
  case ">=":
    return m[key] >= val
  case "<=":
    return m[key] <= val
  case ">":
    return m[key] > val
  case "<":
    return m[key] < val
  default:
    panic(cond)
  }
}

func main() {
  f, err := os.Open("registers.txt")
	if err != nil {
		panic(err)
	}

  r := bufio.NewReader(f)

  m := make(map[string]int)
  max1 := 0

  for l, isPrefix, err := r.ReadLine(); err != io.EOF; l, isPrefix, err = r.ReadLine() {

    if isPrefix || err != nil {
  		if isPrefix {
  			panic("is prefix")
  		} else {
  			panic(err)
  		}
  	}

    lArr := strings.Split(string(l), " ")

    if _, ok := m[lArr[0]]; !ok {
      m[lArr[0]] = 0
    }
    if _, ok := m[lArr[4]]; !ok {
      m[lArr[4]] = 0
    }

    if isCondition(lArr, m) {
      val, err := strconv.Atoi(lArr[2])

      if err != nil {
        panic(err)
      } else {
        if lArr[1] == "inc" {
          m[lArr[0]] += val
        } else {
          m[lArr[0]] -= val
        }
      }
    }
    if m[lArr[0]] > max1 {
      max1 = m[lArr[0]]
    }
  }

  max2 := -99999
  maxKey := ""
  for key, val := range m {
    if val > max2 {
      max2 = val
      maxKey = key
    }
  }

  fmt.Println("The register with the largest value is", maxKey, "with:", max2);
  fmt.Println("The largest value ever held in a register is:", max1)
}
