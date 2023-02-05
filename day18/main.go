package main

import (
	"fmt"
	"../util"
)

type Grid [100][100]int

func main() {
	path := "./day18/input.txt"
	lines, _ := util.ReadLines(path)
  grid := parseLines(lines)
	fmt.Println("-- Part 1 --")
	fmt.Println(part1(grid))
	fmt.Println("-- Part 2 --")
	fmt.Println(part2(grid))
}
func parseLines(lines []string) Grid {
  grid := Grid{}
  for r, line := range lines {
    for c, ch := range line {
      if string(ch) == "#" {
        grid[r][c] = 1
      } else {
        grid[r][c] = 0
      }
    }
  }
  return grid
}

func part1(grid Grid) int {
  for i := 0; i < 100; i++ {
    grid = step(grid)
  }
	return brightness(grid)
}

func step(grid Grid) Grid {
  next := Grid{}
  for r, row := range grid {
    for c, v := range row {
      neighbors := 0
      for nr := max(0, r - 1); nr <= min(r + 1, len(grid) - 1); nr++ {
        for nc := max(0, c - 1); nc <= min(c + 1, len(row) - 1); nc++ {
          if nr == r && nc == c {
            continue
          }
          if grid[nr][nc] == 1 {
            neighbors++
          }
        }
      }
      if v == 1 {
        if neighbors == 2 || neighbors == 3 {
          next[r][c] = 1
        } else {
          next[r][c] = 0
        }
      } else {
        if neighbors == 3 {
          next[r][c] = 1
        } else {
          next[r][c] = 0
        }
      }
    }
  }
  return next
}

func part2(grid Grid) int {
  grid[0][0] = 1
  grid[0][99] = 1
  grid[99][0] = 1
  grid[99][99] = 1
  for i := 0; i < 100; i++ {
    grid = step2(grid)
  }
	return brightness(grid)
}

func step2(grid Grid) Grid {
  next := Grid{}
  for r, row := range grid {
    for c, v := range row {
      neighbors := 0
      for nr := max(0, r - 1); nr <= min(r + 1, len(grid) - 1); nr++ {
        for nc := max(0, c - 1); nc <= min(c + 1, len(row) - 1); nc++ {
          if nr == r && nc == c {
            continue
          }
          if grid[nr][nc] == 1 {
            neighbors++
          }
        }
      }
      if (r == 0 || r == 99) && (c == 0 || c == 99) {
        next[r][c] = 1
      } else if v == 1 {
        if neighbors == 2 || neighbors == 3 {
          next[r][c] = 1
        } else {
          next[r][c] = 0
        }
      } else {
        if neighbors == 3 {
          next[r][c] = 1
        } else {
          next[r][c] = 0
        }
      }
    }
  }
  return next
}

func max(n1, n2 int) int {
  if n2 > n1 {
    return n2
  }
  return n1
}

func min(n1, n2 int) int {
  if n2 < n1 {
    return n2
  }
  return n1
}

func brightness(grid Grid) int {
  brightness := 0
  for _, row := range grid {
    for _, v := range row {
      brightness += v
    }
  }
  return brightness
}
