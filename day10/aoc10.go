package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type Queue struct {
	items []byte
}

func (q *Queue) Push(value byte) {
	q.items = append(q.items, value)
}

func (q *Queue) Pop() byte {
	v := q.items[len(q.items)-1]
	q.items = q.items[:len(q.items)-1]
	return v
}

func (q *Queue) GetFirst() byte {
	return q.items[0]
}

func (q *Queue) GetLast() byte {
	return q.items[len(q.items)-1]
}

func (q *Queue) GetLength() int {
	return len(q.items)
}

func (q *Queue) IsEmpty() bool {
	return 0 == len(q.items)
}

type NavLine struct {
	q             Queue
	corrupted     bool
	corruptedChar byte
}

func (n *NavLine) GetCompletionScore() int {
	var score int
	for {
		if n.q.GetLength() == 0 {
			break
		}

		score *= 5
		switch n.q.Pop() {
		case '(':
			score += 1
			break
		case '[':
			score += 2
			break
		case '{':
			score += 3
			break
		case '<':
			score += 4
			break
		}
	}

	return score
}

const (
	round  = 3
	square = 57
	curly  = 1197
	angle  = 25137
)

func parseData(day int) []string {
	file, _ := ioutil.ReadFile(fmt.Sprintf("day%d/input.txt", day))
	// file, _ := ioutil.ReadFile(fmt.Sprintf("day%d/testinput.txt", day))
	data := strings.Split(string(file), "\n")
	return data[:len(data)-1]
}

func processLine(navline *NavLine, line string) {
	for i := 0; i < len(line); i++ {
		switch line[i] {
		case '{', '[', '<', '(':
			navline.q.Push(line[i])
			break
		case '}', ']', '>':
			if (navline.q.GetLast() + 2) == line[i] {
				navline.q.Pop()
			} else {
				navline.corrupted = true
				navline.corruptedChar = line[i]
				return
			}
			break
		case ')':
			if (navline.q.GetLast() + 1) == line[i] {
				navline.q.Pop()
			} else {
				navline.corrupted = true
				navline.corruptedChar = line[i]
				return
			}
			break
		}
	}
}

func getCorruptionScore(v byte) int {
	switch v {
	case ')':
		return round
	case ']':
		return square
	case '}':
		return curly
	case '>':
		return angle
	}

	return 0
}

/* We check for a closing element that doens't have a matching
opening element */
func getCorruptedLines(data []string) int {
	var corruptionScore int
	for _, line := range data {
		navline := NavLine{Queue{}, false, 0}
		navline.q.items = make([]byte, 0)
		processLine(&navline, line)
		if navline.corrupted {
			corruptionScore += getCorruptionScore(navline.corruptedChar)
		}
	}
	return corruptionScore
}

func part1(data []string) {
	fmt.Println(getCorruptedLines(data))
}

func part2(data []string) {
	scores := make([]int, 0)
	for _, line := range data {
		navline := NavLine{Queue{}, false, 0}
		navline.q.items = make([]byte, 0)
		processLine(&navline, line)
		if navline.corrupted == false {
			scores = append(scores, navline.GetCompletionScore())
		}
	}
	fmt.Println("GetCompletionScores", scores)
	sort.Ints(scores)
	fmt.Println("Middle score is", scores[(len(scores)/2)])
}

func main() {
	data := parseData(10)
	// part1(data)
	part2(data)
}
