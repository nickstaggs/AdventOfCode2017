package main

import (
	"bufio"
	"io"
	"fmt"
	"os"
)

type stack []string

func (s *stack) push(v string) {
    *s = append(*s, v)
}

func (s *stack) pop() string {
    l := len(*s)
    if l == 0 {
      panic("Trying to pop element of empty stack")
    }
    c := (*s)[l-1]
    *s = (*s)[:l-1]
    return c
}

func (s stack) peek() string {
  if (len(s) == 0) {
    return ""
  }
  return s[len(s)-1]
}

func (s stack) getHeight() int {
  return len(s)
}

func main() {
  f, err := os.Open("stream.txt")
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)
  s := make(stack, 0)
  sum := 0
  gSum := 0

  for b, err := r.ReadByte(); err != io.EOF; b, err = r.ReadByte() {
    if err != nil {
      panic(err)
    }
    c := string(b)

    switch c {
    case "!":
      _, err := r.ReadByte()
      if err != nil {
        panic(err)
      }
    case "{":
      if s.peek() != "<" {
        s.push(c)
      } else if s.peek() == "<" {
        gSum++
      }
    case "<":
      if s.peek() != "<" {
        s.push(c)
      } else if s.peek() == "<" {
        gSum++
      }
    case ">":
      if s.peek() == "<" {
        _ = s.pop()
      }
    case "}":
      if s.peek() == "{" {
        _ = s.pop()
        sum += s.getHeight() + 1
      } else if s.peek() == "<" {
        gSum++
      }
    default:
      if s.peek() == "<" {
        gSum++
      }
    }
  }

  fmt.Println("The total score is:", sum)
  fmt.Println("The total amount of garbage is:", gSum)
}
