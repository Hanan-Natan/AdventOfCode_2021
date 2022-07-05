package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type EnergyLevel struct {
	grid    [][]int
	height  int
	width   int
	flashed [][]int
	flash   [][]int
}

func (e *EnergyLevel) Step() {
	e.flashed = make([][]int, len(e.grid))
	for l := 0; l < len(e.grid); l++ {
		for c := 0; c < len(e.grid[l]); c++ {
			e.grid[l][c] += 1
			if e.grid[l][c] > 9 {
				e.flash[l] = append(e.flash[l], c)
			}
		}
	}
}

func (e *EnergyLevel) PrintGrid() {
	for l := 0; l < len(e.grid); l++ {
		fmt.Println(e.grid[l])
	}
}

func (e *EnergyLevel) IsFlashed(line, col int) bool {
	for _, c := range e.flashed[line] {
		if c == col {
			return true
		}
	}
	return false
}

func (e *EnergyLevel) ResetFlashed() {
	for l := 0; l < len(e.grid); l++ {
		for c := 0; c < len(e.grid[l]); c++ {
			if e.grid[l][c] > 9 {
				e.grid[l][c] = 0
			}
		}
	}
}

func (e *EnergyLevel) GetFlashedCount() int {
	var count int
	for i := range e.flashed {
		count += len(e.flashed[i])
	}

	return count
}

func (e *EnergyLevel) Flash() {
	// left, right, down, up, rightDown, rightUp, leftDown, leftUp
	row := []int{0, 0, 1, -1, 1, -1, 1, -1}
	col := []int{-1, 1, 0, 0, 1, 1, -1, -1}
	flashCount := 0

	for l := 0; l < len(e.grid); l++ {
		for c := 0; c < len(e.grid[l]); c++ {
			if e.grid[l][c] > 9 && e.IsFlashed(l, c) == false {
				flashCount += 1
				e.flashed[l] = append(e.flashed[l], c)
				for r := range row {
					if (l+row[r]) >= 0 && ((l + row[r]) < e.height) && (c+col[r]) >= 0 && ((c + col[r]) < e.width) {
						e.grid[row[r]+l][c+col[r]] += 1
					}
				}
			}
		}
	}

	if flashCount != 0 {
		e.Flash()
	}

	return
}

func parseData(day int) [][]int {
	file, _ := ioutil.ReadFile(fmt.Sprintf("day%d/input.txt", day))
	// file, _ := ioutil.ReadFile(fmt.Sprintf("day%d/testinput.txt", day))
	// file, _ := ioutil.ReadFile(fmt.Sprintf("day%d/test1.txt", day))
	data := strings.Split(string(file), "\n")

	lines := make([][]int, len(data)-1, len(data)-1)
	for lineIdx, d := range data {
		if len(d) > 0 {
			tmp := make([]int, len(d))
			for i, v := range d {
				tmp[i], _ = strconv.Atoi(string(v))
			}
			lines[lineIdx] = tmp
		}
	}

	return lines
}

func part1(data [][]int) {
	el := EnergyLevel{
		data,
		len(data),
		len(data[0]),
		make([][]int, len(data)),
		make([][]int, len(data)),
	}
	fmt.Println("Height", len(data), "Width", len(data[0]))

	var flashCount int
	for i := 0; i < 100; i++ {
		el.Step()
		el.Flash()
		flashCount += el.GetFlashedCount()
		el.ResetFlashed()
	}
	fmt.Println("Total flashed", flashCount)
}

func part2(data [][]int) {
	el := EnergyLevel{
		data,
		len(data),
		len(data[0]),
		make([][]int, len(data)),
		make([][]int, len(data)),
	}
	fmt.Println("Height", len(data), "Width", len(data[0]))

	var flashCount int
	for i := 0; i < 1000; i++ {
		el.Step()
		el.Flash()
		flashCount += el.GetFlashedCount()
		if el.GetFlashedCount() == el.height*el.width {
			fmt.Println("The first step all flash simultenously is: ", i+1)
			break
		}
		el.ResetFlashed()
	}
	fmt.Println("Total flashed", flashCount)
}

/*
I'm solving it too structurally. I'm a bit disappointed that i didn't do it using recursion.
*/
func main() {
	data := parseData(11)
	// part1(data)
	part2(data)
}
