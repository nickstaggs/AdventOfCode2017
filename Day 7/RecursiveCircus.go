package main

import (
	"bufio"
	"io"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Program struct {
  name string
  size int
  holding []Program
}

func getSize(n string) int {
  nArr := strings.Split(n, "(")
  nArr = strings.Split(nArr[1], ")")
  nInt, err := strconv.Atoi(nArr[0])

  if err != nil {
    panic(err)
  }
  return nInt
}

func getHolding(h string) []Program {
  hArr := strings.Split(h, ", ")
  ps := make([]Program, 0)

  for i := 0; i < len(hArr); i++ {
    ps = append(ps, Program{hArr[i], 0, nil})
  }

  return ps
}

func (p Program)in(ps []Program) bool {
  for i := 0; i < len(ps); i++ {
    if p.name == ps[i].name {
      ps[i] = p
      return true
    } else if ps[i].holding != nil && p.in(ps[i].holding){
      return true
    }
  }

  return false
}

func (p Program)checkBalance() int {
  ns := make([]int, 0)
  for i := 0; i < len(p.holding); i++ {
    ns = append(ns, p.holding[i].checkBalance())
  }

  if len(ns) != 0 {
    diff, index := unbalancedIndexDifference(ns)

    if diff != 0 {
      panic(p.holding[index].size + diff)
    }
  }

  return sum(ns) + p.size
}

func sum(ns []int) int {
  sum := 0
  for i := 0; i < len(ns); i++ {
      sum += ns[i]
  }
  return sum
}

func unbalancedIndexDifference(ns []int) (int, int) {
  counter := make([][]int, 2)
  for i := range counter {
    counter[i] = make([]int, 2)
  }

  counter[0][0] = ns[0]
  counter[0][1] = 1
  for i := 1; i < len(ns); i++ {
    if ns[i] == counter[0][0] {
      counter[0][1]++
    } else {
      counter[1][0] = ns[i]
      counter[1][1]++
    }
  }

  var correct int

  if counter[0][1] > 1 {
    correct = counter[0][0]
  } else {
    correct = counter[1][0]
  }

  var index int
  for i := 0; i < len(ns); i++ {
    if ns[i] != correct {
      index = i
    }
  }

  return correct-ns[index], index
}

func main() {
  f, err := os.Open("towers.txt")
	if err != nil {
		panic(err)
	}

  r := bufio.NewReader(f)

  ps := make([]Program, 0)
  for line, isPrefix, err := r.ReadLine(); err != io.EOF; line, isPrefix, err = r.ReadLine() {

  	if isPrefix || err != nil {
  		if isPrefix {
  			panic("is prefix")
  		} else {
  			panic(err)
  		}
  	}

    lArr := strings.Split(string(line), " -> ")

    base := lArr[0]

    bArr := strings.Split(base, " ")

    p := Program{bArr[0], getSize(bArr[1]), nil}

    if len(lArr) > 1 {
      p.holding = getHolding(lArr[1])
    }

    ps = append(ps, p)
  }

  i := 0
  for i < len(ps) {
    for j := 0; j < len(ps); j++ {
      if ps[j].holding != nil && ps[i].in(ps[j].holding) {
        ps = append(ps[:i], ps[i+1:]...)
        break
      }
      if j == len(ps) - 1 {
        i++
      }
    }
  }

  fmt.Println("The bottom program is", ps[0].name)

  defer func() {
        if r := recover(); r != nil {
            fmt.Println("Unbalanced tree, weight should be", r)
        }
  }()

  ps[0].checkBalance()
}
