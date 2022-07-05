package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "strconv"
)

func part2() {
  file, _ := ioutil.ReadFile("day2/input.txt")
  lines := strings.Split(string(file), "\n")

  var horizontalPos, aim, depth int

  for _, line := range lines {
    v := strings.Fields(line)

    if len(v) == 0 {
      break
    }

    num, err := strconv.Atoi(v[1])
    if nil == err {
      if v[0] == "forward" {
        horizontalPos += num
        depth += num * aim
      } else if v[0] == "down" {
        aim += num
      } else if v[0] == "up" {
        aim -= num
      }
    }
  }

  fmt.Println("Our position is: ", horizontalPos * depth)
}

func part1() {
  file, _ := ioutil.ReadFile("day2/input.txt")
  lines := strings.Split(string(file), "\n")

  var forward, depth int

  for _, line := range lines {
    v := strings.Fields(line)

    if len(v) == 0 {
      break
    }

    num, err := strconv.Atoi(v[1])
    if nil == err {
      if v[0] == "forward" {
        forward += num
      } else if v[0] == "down" {
        depth += num
      } else if v[0] == "up" {
        depth -= num
      }
    }
  }

  fmt.Println("Our position is: ", forward * depth)

  }

func main() {
  // part1()
  part2()
}
