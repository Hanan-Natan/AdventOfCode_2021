package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	CellsPerRow = 5
	MarkedNum   = -1
)

type Board struct {
	rowCell    [CellsPerRow][CellsPerRow]int
	HasWon     bool
	BoardValue int
}

type Game struct {
	Boards         []*Board
	LastBoardToWin int
}

func (game *Game) CreateBoards(lines []string) {
	var row, boardCount int

	game.Boards = append(game.Boards, &Board{})
	for _, line := range lines {
		if len(line) == 0 {
			game.Boards = append(game.Boards, &Board{})
			boardCount += 1
			row = 0
			continue
		}

		for c, l := range strings.Fields(line) {
			num, err := strconv.Atoi(l)
			if nil == err {
				game.Boards[boardCount].rowCell[row][c] = num
			}
		}
		row += 1
	}
}

func (game *Game) GetLastWinningBoardValue() int {
	if game.LastBoardToWin != -1 {
		return game.Boards[game.LastBoardToWin].BoardValue
	}

	return 0
}

func (game *Game) PrintBoards() {
	for idx, board := range game.Boards {
		fmt.Println("Board Num: ", idx)
		for _, row := range board.rowCell {
			fmt.Println(row)
		}
	}
}

func (game *Game) MarkDrawNum(draw []int) {
	// var lastDraw int
	for d := 0; d < len(draw); d++ {
		for boardKey, board := range game.Boards {
			for rowKey, row := range board.rowCell {
				for i := 0; i < len(row); i++ {
					if row[i] == draw[d] {
						game.Boards[boardKey].rowCell[rowKey][i] = MarkedNum
					}
				}
			}
			game.MarkWinningBoard()
			// part1
			// if game.GetLastWinningBoardValue() > 0 {
			// fmt.Println("The results are: ", game.GetLastWinningBoardValue(), draw[d], game.GetLastWinningBoardValue()*draw[d])
			//      panic("Done, part1")
			// }
		}
		// part2
		if game.HasAllBoardsWon() {
			fmt.Println("Last winning board is: ", game.LastBoardToWin)
			fmt.Println(
				"The results are: ",
				game.GetLastWinningBoardValue(),
				draw[d],
				game.GetLastWinningBoardValue()*draw[d],
			)
		}
	}
}

func (game *Game) HasAllBoardsWon() bool {
	for _, board := range game.Boards {
		if board.HasWon == false {
			return false
		}
	}
	return true
}

func (b *Board) SumBoardValue() {
	var sum int
	for _, row := range b.rowCell {
		for _, cV := range row {
			if cV != MarkedNum {
				sum += cV
			}
		}
	}
	b.BoardValue = sum
}

func (game *Game) MarkWinningBoard() {
	for boardIdx, board := range game.Boards {
		var colMatch int
		if board.HasWon == false {
			for rowIdx, row := range board.rowCell {
				var rowMatch int
				for colIdx, colValue := range row {
					if colValue == MarkedNum {
						rowMatch += 1

						if len(row) == rowMatch {
							// row match
							board.HasWon = true
							board.SumBoardValue()
							game.LastBoardToWin = boardIdx
							break
							// return
						}
						// column match
						if rowIdx == 0 {
							colMatch += 1
							for _, r := range board.rowCell[1:] {
								if r[colIdx] == MarkedNum {
									colMatch += 1
								}
							}
							if colMatch == len(row) {
								board.HasWon = true
								board.SumBoardValue()
								game.LastBoardToWin = boardIdx
								// break
								// return
							} else {
								colMatch = 0
							}
						}
					}
				}
			}
		}
	}
}

func part1and2() {
	file, _ := ioutil.ReadFile("day4/input.txt")
	// file, _ := ioutil.ReadFile("day4/testdata.txt")
	lines := strings.Split(string(file), "\n")

	draw := strings.Split(lines[0], ",")
	drawNum := make([]int, len(draw), len(draw))
	// convert draw number to int
	for i := 0; i < len(draw); i++ {
		num, err := strconv.Atoi(draw[i])
		if nil == err {
			drawNum[i] = num
		}
	}

	game := new(Game)
	game.LastBoardToWin = -1
	game.CreateBoards(lines[2 : len(lines)-1])
	game.MarkDrawNum(drawNum)

	fmt.Println("Done!")
}

func main() {
	part1and2()
}
