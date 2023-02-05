package main

import (
	"fmt"
  "strconv"
  s "strings"
	"../util"
)

type Entity struct {
  hp int
  dmg int
  armor int
}

type Item struct {
  price int
  dmg int
  armor int
}

func main() {
	path := "./day21/input.txt"
	lines, _ := util.ReadLines(path)
  boss := parseLines(lines)
  player := Entity{hp: 100, dmg: 0, armor: 0}
  weapons := []Item{
    Item{price: 8, dmg: 4, armor: 0},
    Item{price: 10, dmg: 5, armor: 0},
    Item{price: 25, dmg: 6, armor: 0},
    Item{price: 40, dmg: 7, armor: 0},
    Item{price: 74, dmg: 8, armor: 0},
  }
  armor := []Item{
    Item{price: 0, dmg: 0, armor: 0},
    Item{price: 13, dmg: 0, armor: 1},
    Item{price: 31, dmg: 0, armor: 2},
    Item{price: 53, dmg: 0, armor: 3},
    Item{price: 75, dmg: 0, armor: 4},
    Item{price: 102, dmg: 0, armor: 5},
  }
  rings := []Item{
    Item{price: 0, dmg: 0, armor: 0},
    Item{price: 0, dmg: 0, armor: 0},
    Item{price: 25, dmg: 1, armor: 0},
    Item{price: 50, dmg: 2, armor: 0},
    Item{price: 100, dmg: 3, armor: 0},
    Item{price: 20, dmg: 0, armor: 1},
    Item{price: 40, dmg: 0, armor: 2},
    Item{price: 80, dmg: 0, armor: 3},
  }
	fmt.Println("-- Part 1 --")
	fmt.Println(part1(player, boss, weapons, armor, rings))
	fmt.Println("-- Part 2 --")
	fmt.Println(part2(player, boss, weapons, armor, rings))
}

func parseLines(lines []string) Entity {
  parts := s.Split(lines[0], ": ")
  hp, _ := strconv.Atoi(parts[1])
  parts = s.Split(lines[1], ": ")
  dmg, _ := strconv.Atoi(parts[1])
  parts = s.Split(lines[2], ": ")
  armor, _ := strconv.Atoi(parts[1])
  return Entity{hp: hp, dmg: dmg, armor: armor}
}

func part1(player, boss Entity, weapons, armor, rings []Item) int {
  min := 1000
  for _, w := range weapons {
    for _, a := range armor {
      for i := 0; i < len(rings); i++ {
        for j := i + 1; j < len(rings); j++ {
          r1 := rings[i]
          r2 := rings[j]
          price := w.price + a.price + r1.price + r2.price
          player = Entity{hp: 100, dmg: w.dmg + r1.dmg + r2.dmg, armor: a.armor + r1.armor + r2.armor}
          if price < min {
            if fight(player, boss) {
              min = price
            }
          }
        }
      }
    }
  }
	return min
}

// True if the player won
func fight(player, boss Entity) bool {
  for {
    boss = damage(boss, player.dmg)
    if boss.hp <= 0 {
      return true
    }
    player = damage(player, boss.dmg)
    if player.hp <= 0 {
      return false
    }
  }
}

func damage(e Entity, att int) Entity {
  dmg := max(1, att - e.armor)
  return Entity{hp: e.hp - dmg, dmg: e.dmg, armor: e.armor}
}

func max(n1, n2 int) int {
  if n2 > n1 {
    return n2
  }
  return n1
}

func part2(player, boss Entity, weapons, armor, rings []Item) int {
  max := 0
  for _, w := range weapons {
    for _, a := range armor {
      for i := 0; i < len(rings); i++ {
        for j := i + 1; j < len(rings); j++ {
          r1 := rings[i]
          r2 := rings[j]
          price := w.price + a.price + r1.price + r2.price
          player = Entity{hp: 100, dmg: w.dmg + r1.dmg + r2.dmg, armor: a.armor + r1.armor + r2.armor}
          if price > max {
            if !fight(player, boss) {
              max = price
            }
          }
        }
      }
    }
  }
	return max
}
