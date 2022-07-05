package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func printMarks(width int, data map[int]int) {
  for i:=0; i <= (width+1) * 10; i++ {
    v, ok := data[i]
    if ok {
      fmt.Printf("%d", v)
    } else {
      fmt.Printf(".")
    }

    if (i > 0 && i%(width+1) == 0) {
      fmt.Printf("\n")
    }
  }
}

func markLines(width int, data [][]int) map[int]int {
	linesPath := map[int]int{}
	for _, line := range data {
		x1, y1, x2, y2 := line[0], line[1], line[2], line[3]

		height := min(y1, y2)
		crossHeight := (width+1) * height 

		if y1 == y2 {
			horizontalLength := abs(x1 - x2)
			for i := 0; i <= horizontalLength; i++ {
				linesPath[crossHeight+(min(x1, x2)+i)] += 1
			}
		} else if x1 == x2 {
			verticalLength := abs(y1 - y2)
			for i := 0; i <= verticalLength; i++ {
				linesPath[(crossHeight+min(x1, x2))+((width+1)*i)] += 1
			}
		} else {
			// diagonal line
			diagonalLength := abs(x2 - x1)
      diagonalDirection := min(x1, x2)
			if (y1 > y2 && x1 < x2) || (y1 < y2 && x1 > x2) {
				diagonalDirection = max(x1, x2) // top right to bottom left 
			}

			for i := 0; i <= diagonalLength; i++ {
				if diagonalDirection == max(x1,x2)  { // top right to bottom left 
					// fmt.Println("Marking top right to bottom left : ", line, diagonalLength, (crossHeight + ((width+1) * i) + diagonalDirection - i))
					linesPath[(crossHeight + ((width+1) * i) + diagonalDirection -i )] += 1
				} else { // top left to bottom right
					// fmt.Println("Marking top left to botton right : ", line, diagonalLength, (crossHeight + ((width+1) * i) + diagonalDirection + i))
					linesPath[(crossHeight + ((width+1) * i) + diagonalDirection + i)] += 1
				}
			}
		}
	}

  printMarks(width, linesPath)
	return linesPath
}

func convertNumbers(data []string, straightLines, diagonalLines bool) [][]int {
	res := [][]int{}
	for _, line := range data[:len(data)-1] {
		lineData := strings.Split(line, "->")

		xy := strings.Split(strings.ReplaceAll(strings.Join(lineData, ","), " ", ""), ",")
		tmp1, _ := strconv.Atoi(xy[0])
		tmp2, _ := strconv.Atoi(xy[1])
		tmp3, _ := strconv.Atoi(xy[2])
		tmp4, _ := strconv.Atoi(xy[3])

		x1 := tmp1
		y1 := tmp2
		x2 := tmp3
		y2 := tmp4

		if straightLines && ((x1 == x2) || (y1 == y2)) {
			res = append(res, []int{x1, y1, x2, y2})
			fmt.Println("straight line: ", x1, y1, x2, y2)
		} else if diagonalLines && (abs(x2-x1) == abs(y2-y1)) {
			res = append(res, []int{x1, y1, x2, y2})
			fmt.Println("diagonal line: ", x1, y1, x2, y2)
		}
	}
	return res
}

func getWidth(data [][]int) int {
	var width int
	for _, line := range data {
		if max(line[0], line[2]) > width {
			width = max(line[0], line[2])
		}
	}

	return width 
}

func countCrossings(lines map[int]int) int {
	var crossings int
	for _, cross := range lines {
		if cross > 1 {
			crossings += 1
		}
	}
	return crossings
}

func part2(data []string) {
	numbers := convertNumbers(data, true, true)
	width := getWidth(numbers)
	fmt.Println("The width is: ", width)
	lines := markLines(width, numbers)
	crossings := countCrossings(lines)
	fmt.Println("Tne number of crossings are: ", crossings)
}

func part1(data []string) {
	numbers := convertNumbers(data, true, false)
	width := getWidth(numbers)
	fmt.Println("The width is: ", width)
	lines := markLines(width, numbers)
	crossings := countCrossings(lines)
	fmt.Println("Tne number of crossings are: ", crossings)
}

func main() {
	// file, _ := ioutil.ReadFile("day5/input.txt")
	file, _ := ioutil.ReadFile("day5/testdata.txt")
	lines := strings.Split(string(file), "\n")
	// part1(lines)
	part2(lines)
}
