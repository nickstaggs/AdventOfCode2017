package main

import (
	"bufio"
	"fmt"
	"os"
  "strings"
  "sort"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func IsAnagram(a, b string) bool{
  a = SortString(a)
  b = SortString(b)

  if strings.Compare(a, b) == 0 {
    return true
  }

  return false
}

func main() {
  f, err := os.Open("passphrases.txt")
	if err != nil {
		fmt.Println("Panicking...")
		panic(1)
	}

  r := bufio.NewReader(f)

  l, isPrefix, err := r.ReadLine()

  if isPrefix {
    panic("isPrefix")
  }

  numValid := 0
  for ; err == nil; l, isPrefix, err = r.ReadLine() {
    passPhrase := string(l)
    passPhraseArr := strings.Split(passPhrase, " ")
    valid := true

    for i := 0; i < len(passPhraseArr) - 1 && valid; i++ {
      for j := i + 1; j < len(passPhraseArr); j++ {
        // comment out is IsAnagram for solution to part 1 of this day's advent of code
        if strings.Compare(passPhraseArr[i], passPhraseArr[j]) == 0 ||
                           IsAnagram(passPhraseArr[i], passPhraseArr[j]){
          valid = false
        }
      }
    }

    if valid {
      numValid++
    }
  }

  fmt.Println("The number of valid passwords is:", numValid)
}
