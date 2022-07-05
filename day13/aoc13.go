package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Board struct {
	table        [][]int
	instructions [][]string
}

func (b *Board) PrintBoard() {
	for _, t := range b.table {
		fmt.Println(t)
	}
}

func (b *Board) FoldHorizentaly(lines int) {
	tmp := make([][]int, lines, lines)
	for t := range tmp {
		tmp[t] = b.table[t]
	}
	b.table = tmp
}

func (b *Board) FoldVertically(cIdx int) {
	cIdx += 1
	lineSize := len(b.table[0]) - cIdx
	fmt.Println(lineSize)
	for l := range b.table {
		tmp := make([]int, lineSize, lineSize)
		for c := 0; c < lineSize; c++ {
			tmp[c] = b.table[l][c+cIdx]
		}
		b.table[l] = tmp
	}
}

func (b *Board) CountDots() int {
	coutner := 0
	for l := range b.table {
		for _, v := range b.table[l] {
			if v > 0 {
				coutner += 1
			}
		}
	}

	return coutner
}

func (b *Board) Fold() {
	for _, instruction := range b.instructions {
		starting, _ := strconv.Atoi(instruction[1])
		if instruction[0] == string('y') {
			for i := 1; i < len(b.table)-starting; i++ {
				for c := range b.table[starting+i] {
					b.table[starting-i][c] |= b.table[starting+i][c]
				}
			}
			b.FoldHorizentaly(starting)
		} else {
			for l := range b.table {
				for c := 1; c <= starting; c++ {
					b.table[l][starting+c] |= b.table[l][starting-c]
				}
			}
			b.FoldVertically(starting)
		}
	}

}

func parseData(day int) *Board {
	file, _ := ioutil.ReadFile(fmt.Sprintf("day%d/input.txt", day))
	// file, _ := ioutil.ReadFile(fmt.Sprintf("day%d/larger.txt", day))
	// file, _ := ioutil.ReadFile(fmt.Sprintf("day%d/testinput.txt", day))
	data := strings.Split(string(file), "\n")
	data = data[:len(data)-1]

	input := make([][]int, 0)
	instructions := make([][]string, 0)
	width, height := 0, 0

	for _, l := range data {
		if strings.HasPrefix(l, "fold") {
			v := strings.Split(l, " ")
			instructions = append(instructions, strings.Split(v[len(v)-1], "="))
		} else if len(l) > 0 {
			line := strings.Split(l, ",")
			x, _ := strconv.Atoi(line[0])
			y, _ := strconv.Atoi(line[1])
			if x > width {
				width = x
			}
			if y > height {
				height = y
			}
			input = append(input, []int{x, y})
		}

	}

	b := &Board{
		make([][]int, height+1, height+1),
		make([][]string, len(instructions)),
	}

	for y := 0; y <= height; y++ {
		line := make([]int, width+1, width+1)
		for x := 0; x <= width; x++ {
			line[x] = 0
		}
		b.table[y] = line
	}

	for _, l := range input {
		b.table[l[1]][l[0]] = 1
	}

	for i, l := range instructions {
		b.instructions[i] = l
	}

	return b
}

func part1(info *Board) {
	info.Fold()
	fmt.Println("Total dots after first fold is:", info.CountDots())
	info.PrintBoard()
}

/*
Done!
*/
func main() {
	data := parseData(13)
	part1(data)
}
