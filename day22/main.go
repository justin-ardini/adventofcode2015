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
  mana int
}

type Spell struct {
  cost int
  dmg int
  heal int
  effect Effect
}

type Effect struct {
  turns int
  dmg int
  armor int
  mana int
}

func main() {
	path := "./day22/input.txt"
	lines, _ := util.ReadLines(path)
  boss := parseLines(lines)
  player := Entity{hp: 50, dmg: 0, armor: 0, mana: 500}
  spells := []Spell{
    Spell{cost: 53, dmg: 4},
    Spell{cost: 73, dmg: 2, heal: 2},
    Spell{cost: 113, effect: Effect{turns: 6, armor: 7}},
    Spell{cost: 173, effect: Effect{turns: 6, dmg: 3}},
    Spell{cost: 229, effect: Effect{turns: 5, mana: 101}},
  }
	fmt.Println("-- Part 1 --")
	fmt.Println(part1(player, boss, spells))
	fmt.Println("-- Part 2 --")
	fmt.Println(part2(player, boss, spells))
}

func parseLines(lines []string) Entity {
  parts := s.Split(lines[0], ": ")
  hp, _ := strconv.Atoi(parts[1])
  parts = s.Split(lines[1], ": ")
  dmg, _ := strconv.Atoi(parts[1])
  return Entity{hp: hp, dmg: dmg, armor: 0}
}

func part1(player, boss Entity, spells []Spell) int {
  effects := [5]Effect{}
  min := 10000
  for i, s := range spells {
    winner, mana := fightRec(player, boss, spells, i, effects, s.cost, false)
    if winner && mana < min {
      min = mana
    }
  }
	return min
}

func fightRec(player, boss Entity, spells []Spell, si int, effects [5]Effect, mana int, p2 bool) (bool, int) {
  if mana > 1500 {
    return false, -1
  }
  if player.mana < spells[si].cost {
    return false, -1
  }
  player, boss, effects = castSpell(spells, si, player, boss, effects)
  if boss.hp <= 0 {
    return true, mana
  }

  // Boss turn
  player, boss, effects = runEffects(player, boss, effects)
  if boss.hp <= 0 {
    return true, mana
  }
  player = damage(player, boss.dmg)
  if player.hp <= 0 {
    return false, -1
  }

  // Player turn
  if p2 {
    player = damage(player, 1)
  }
  if player.hp <= 0 {
    return false, -1
  }
  player, boss, effects = runEffects(player, boss, effects)
  if boss.hp <= 0 {
    return true, mana
  }

  win := false
  min := 10000
  for i, s := range spells {
    if effects[i].turns > 0 {
      continue
    }
    winner, mana := fightRec(player, boss, spells, i, effects, mana + s.cost, p2)
    if winner && mana < min {
      win = winner
      min = mana
    }
  }
  return win, min
}

func runEffects(player, boss Entity, effects [5]Effect) (Entity, Entity, [5]Effect) {
  player.armor = 0
  for k, effect := range effects {
    if effect.turns > 0 {
      if effect.dmg > 0 {
        boss = damage(boss, effect.dmg)
      }
      if effect.armor > 0 {
        player.armor = effect.armor
      }
      if effect.mana > 0 {
        player.mana += effect.mana
      }
      effects[k] = Effect{turns: effect.turns - 1, dmg: effect.dmg, armor: effect.armor, mana: effect.mana}
    }
  }
  return player, boss, effects
}

func castSpell(spells []Spell, si int, player, boss Entity, effects [5]Effect) (Entity, Entity, [5]Effect) {
  spell := spells[si]
  np := player
  np.mana -= spell.cost
  if spell.heal > 0 {
    np.hp += spell.heal
  }
  if spell.dmg > 0 {
    boss = damage(boss, spell.dmg)
  }
  effects[si] = spell.effect
  return np, boss, effects
}

func damage(e Entity, att int) Entity {
  n := e
  dmg := max(1, att - n.armor)
  n.hp -= dmg
  return n
}

func max(n1, n2 int) int {
  if n2 > n1 {
    return n2
  }
  return n1
}

func part2(player, boss Entity, spells []Spell) int {
  effects := [5]Effect{}
  min := 10000
  player = damage(player, 1)
  for i, s := range spells {
    winner, mana := fightRec(player, boss, spells, i, effects, s.cost, true)
    if winner && mana < min {
      min = mana
    }
  }
	return min
}
