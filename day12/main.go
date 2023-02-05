package main

import (
	"encoding/json"
	"fmt"
	"../util"
)

func main() {
	path := "./day12/input.txt"
	lines, _ := util.ReadLines(path)
  var f interface{}
  json.Unmarshal([]byte(lines[0]), &f)
  m := f.(map[string]interface{})
	fmt.Println("-- Part 1 --")
	fmt.Println(part1(m))
	fmt.Println("-- Part 2 --")
	fmt.Println(part2(m))
}

func part1(m map[string]interface{}) int {
  return mapSum(m)
}

func mapSum(m map[string]interface{}) int {
  total := 0
  for _, v := range m {
    switch vv := v.(type) {
    case string:
      continue
    case float64:
      total += int(vv)
    case []interface{}:
      total += arrSum(vv)
    case map[string]interface{}:
      total += mapSum(vv)
    }
  }
	return total
}

func arrSum(a []interface{}) int {
  total := 0
  for _, v := range a {
    switch vv := v.(type) {
    case string:
      continue
    case float64:
      total += int(vv)
    case []interface{}:
      total += arrSum(vv)
    case map[string]interface{}:
      total += mapSum(vv)
    }
  }
  return total
}

func part2(m map[string]interface{}) int {
  m = stripRed(m)
	return mapSum(m)
}

func stripRed(m map[string]interface{}) map[string]interface{} {
  for k, v := range m {
    switch vv := v.(type) {
    case string:
      if vv == "red" {
        return make(map[string]interface{})
      }
    case float64:
      continue
    case []interface{}:
      m[k] = arrStripRed(vv)
    case map[string]interface{}:
      m[k] = stripRed(vv)
    }
  }
	return m
}

func arrStripRed(a []interface{}) []interface{} {
  for k, v := range a {
    switch vv := v.(type) {
    case string:
      if vv == "red" {
        a[k] = ""
      }
    case float64:
      continue
    case []interface{}:
      a[k] = arrStripRed(vv)
    case map[string]interface{}:
      a[k] = stripRed(vv)
    }
  }
	return a
}
