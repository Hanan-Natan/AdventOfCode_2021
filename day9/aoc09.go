package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

type Directions struct {
	left  int
	right int
	up    int
	down  int
}

func parseData(day int) [][]int {
	file, _ := ioutil.ReadFile(fmt.Sprintf("day%d/input.txt", day))
	// file, _ := ioutil.ReadFile(fmt.Sprintf("day%d/testinput.txt", day))
	data := strings.Split(string(file), "\n")

	lines := make([][]int, len(data)-1, len(data)-1)
	for dataIdx, d := range data {
		if len(d) > 0 {
			tmp := make([]int, len(d), len(d))
			for i, c := range d {
				tmp[i] = int(c - '0')
			}
			lines[dataIdx] = tmp
		}
	}

	return lines
}

func getAdjecentValues(lineIdx, colIdx int, data [][]int) Directions {
	dir := Directions{10, 10, 10, 10}

	if colIdx > 0 {
		dir.left = data[lineIdx][colIdx-1]
	}
	if colIdx < len(data[lineIdx])-1 {
		dir.right = data[lineIdx][colIdx+1]
	}
	if lineIdx > 0 {
		dir.up = data[lineIdx-1][colIdx]
	}
	if len(data) > lineIdx+1 {
		dir.down = data[lineIdx+1][colIdx]
	}

	return dir
}

func findLineLen(direction, colIdx int, data []int) int {
	if colIdx >= len(data) || colIdx < 0 {
		return 0
	} else if data[colIdx] == 9 {
		return 0
	}

	if direction == 1 {
		return 1 + findLineLen(direction, colIdx+1, data)
	} else {
		return 1 + findLineLen(direction, colIdx-1, data)
	}
}

func isSeen(point BasinPoint, seen map[BasinPoint]bool) bool {
	if _, ok := seen[point]; ok {
		return true
	}
	return false
}

type BasinPoint struct {
	line int
	col  int
}

type Basin struct {
	points []BasinPoint
	size   int
}

func (b *Basin) popFirst() BasinPoint {
	bp := b.points[0]
	b.points = b.points[1:]

	return bp
}

func (b *Basin) empty() bool {
	return 0 == len(b.points)
}

func (b *Basin) length() int {
	return len(b.points)
}

func (b *Basin) push(point BasinPoint) {
	b.points = append(b.points, point)
}

func findBasinSize(data [][]int) int {
	seen := make(map[BasinPoint]bool)
	basinSize := make([]int, 0)

	for lineIdx, line := range data {
		for colIdx, v := range line {
			if isSeen(BasinPoint{lineIdx, colIdx}, seen) == false && v != 9 {
				basin := Basin{
					make([]BasinPoint, 0),
					0,
				}

				basin.push(BasinPoint{lineIdx, colIdx})
				for {
					if basin.empty() {
						break
					}
					point := basin.popFirst()
					// fmt.Println("Checking point", point)
					// fmt.Println("Points", basin.points)
					if isSeen(point, seen) {
						continue
					}
					seen[point] = true
					basin.size += 1

					adjPoints := getAdjecentValues(point.line, point.col, data)

					if adjPoints.up < 9 {
						basin.push(BasinPoint{point.line - 1, point.col})
					}
					if adjPoints.down < 9 {
						basin.push(BasinPoint{point.line + 1, point.col})
					}
					if adjPoints.right < 9 {
						basin.push(BasinPoint{point.line, point.col + 1})
					}
					if adjPoints.left < 9 {
						basin.push(BasinPoint{point.line, point.col - 1})
					}
				}
				basinSize = append(basinSize, basin.size)
			}
		}
	}

	sort.Ints(basinSize)
	basinTotal := basinSize[len(basinSize)-1] * basinSize[len(basinSize)-2] * basinSize[len(basinSize)-3]

	return basinTotal
}

func findLowPoint(data [][]int) []int {
	lowPointValues := []int{}
	for lineIdx, line := range data {
		for colIdx, v := range line {
			dir := getAdjecentValues(lineIdx, colIdx, data)
			if v < dir.left && v < dir.right && v < dir.up && v < dir.down {
				lowPointValues = append(lowPointValues, v)
			}
		}
	}

	return lowPointValues
}

func sumLowPointValues(data []int) int {
	var sum int
	for _, v := range data {
		sum += v + 1
	}
	return sum
}

func part2(data [][]int) {
	res := findBasinSize(data)
	fmt.Println(res)
}

func part1(data [][]int) {
	lowPoints := findLowPoint(data)
	sum := sumLowPointValues(lowPoints)
	fmt.Println("Sum of low point values is:", sum)
}

/*
I had real struggle with part2.
I tried to solve it using the same mindset i solve part1 which is to find the lowest point of a basin and then go to each line and column that is cononected to that basin. That didn't work and despite trying many different approaches i failed.
I then decided to search for the solution on the net and I basically copy the solution from here:
https://github.com/jonathanpaulson/AdventOfCode/blob/master/2021/9.py
*/

func main() {
	data := parseData(9)
	// part1(data)
	part2(data)
}
