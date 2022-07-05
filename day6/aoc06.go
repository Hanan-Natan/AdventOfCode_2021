package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func parseData() []int {
	file, _ := ioutil.ReadFile("day6/input.txt")
	// file, _ := ioutil.ReadFile("day6/testdata.txt")
	lines := strings.Split(string(file), "\n")

	data := strings.Split(lines[0], ",")

	res := make([]int, len(data), len(data))
	for i := 0; i < len(data); i++ {
		v, _ := strconv.Atoi(data[i])
		res[i] = v
	}

	return res
}

// based on idea taken from: https://www.youtube.com/watch?v=AFixiY-5jTc
func process_part2(data []int, days int) int {
	// create a state
	state := make([]int, 9, 9)
	for s := 0; s < len(state); s++ {
		for i := 0; i < len(data); i++ {
			if s == data[i] {
				state[s] += 1
			}
		}
	}

	// process the state
	for d := 1; d <= days; d++ {
		newState := make([]int, 9, 9)
		for s := 1; s < len(state); s++ {
			newState[s-1] = state[s]
		}
		newState[6] += state[0]
		newState[8] += state[0]

		copy(state, newState)
	}
	fmt.Println(state)

	// calculate the sum
	var sum int
	for i := 0; i < len(state); i++ {
		sum += state[i]
	}
	return sum
}

func process(data []int, days int) []int {
	nextDay := []int{}

	iter := make([]int, len(data))
	copy(iter, data)

	for d := 1; d <= days; d++ {
		nextDay = []int{}
		for f := 0; f < len(iter); f++ {
			if iter[f] == 0 {
				nextDay = append(nextDay, 8, 6)
			} else {
				nextDay = append(nextDay, iter[f]-1)
			}
		}
		iter = make([]int, len(nextDay))
		copy(iter, nextDay)
	}

	return nextDay
}

func main() {
	data := parseData()
	// data = process(data, 80) // part1
	// fmt.Println(data)
	// fmt.Println("Number of fish are: ", len(data))
	sum := process_part2(data, 256)
	fmt.Println(sum)
	fmt.Println("Number of fish are: ", sum)
}
