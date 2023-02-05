package main

import (
  "crypto/md5"
	"fmt"
	"strconv"
	"../util"
)

func main() {
	path := "./day04/input.txt"
	lines, _ := util.ReadLines(path)
	fmt.Println("-- Part 1 --")
	fmt.Println(part1(lines[0]))
	fmt.Println("-- Part 2 --")
	fmt.Println(part2(lines[0]))
}

func part1(key string) int {
  for i := 1;; i++ {
    v := key + strconv.Itoa(i)
    h := md5.Sum([]byte(v))
    if h[0] == 0 && h[1] == 0 && h[2] < 16 {
      return i
    }
  }
}

func part2(key string) int {
  for i := 1;; i++ {
    v := key + strconv.Itoa(i)
    h := md5.Sum([]byte(v))
    if h[0] == 0 && h[1] == 0 && h[2] == 0 {
      return i
    }
  }
}
