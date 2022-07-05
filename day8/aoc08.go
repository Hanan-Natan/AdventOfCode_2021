package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var wireToInt = map[int]int{
	2: 1,
	4: 4,
	3: 7,
	7: 8,
}

func parseData(day int) []string {
	file, _ := ioutil.ReadFile(fmt.Sprintf("day%d/input.txt", day))
	// file, _ := ioutil.ReadFile(fmt.Sprintf("day%d/testinput.txt", day))
	data := strings.Split(string(file), "\n")

	return data
}

func containsChars(str, chars string) bool {
	for x := 0; x < len(chars); x++ {
		if false == strings.Contains(str, string(chars[x])) {
			return false
		}
	}
	return true
}

func extractNotFound(str, chars string) byte {
	long, short := str, chars
	if len(long) < len(chars) {
		long, short = chars, str
	}

	for x := 0; x < len(short); x++ {
		if false == strings.Contains(long, string(short[x])) {
			return short[x]
		}
	}

	for x := 0; x < len(long); x++ {
		if false == strings.Contains(short, string(long[x])) {
			return long[x]
		}
	}
	return 0
}

func isNumbersFull(numbers map[int]string) bool {
	for i := 0; i < 10; i++ {
		_, ok := numbers[i]
		if false == ok {
			return false
		}
	}
	return true
}

func isElmIn(data map[int]string, elm string) bool {
	for _, i := range data {
		if elm == i {
			return true
		}
	}
	return false
}
func deduceNumbers(signal []string) map[int]string {

	numbers := map[int]string{}
	letters := map[int]byte{}
	// get the obvious 1, 4, 7, 8
	for i := 0; i < len(signal); i++ {
		v, ok := wireToInt[len(signal[i])]
		if ok {
			numbers[v] = signal[i]
		}
	}

	letters[0] = extractNotFound(numbers[1], numbers[7])

	nineComplement := numbers[4] + string(letters[0])
	for {
		if isNumbersFull(numbers) {
			break
		}
		for i := 0; i < len(signal); i++ {
			if len(signal[i]) == 6 {
				if containsChars(signal[i], nineComplement) {
					// map number 9
					numbers[9] = signal[i]
					letters[3] = extractNotFound(signal[i], nineComplement)
				} else if false == containsChars(signal[i], numbers[1]) {
					// map number 6
					numbers[6] = signal[i]
					letters[1] = extractNotFound(signal[i], numbers[1])
					if numbers[1][0] == letters[1] {
						letters[2] = numbers[1][1]
					} else {
						letters[2] = numbers[1][0]
					}
				} else if false == containsChars(signal[i], numbers[6]) {
					// map number 0
					numbers[0] = signal[i]
					letters[4] = extractNotFound(signal[i], numbers[7]+string(letters[3])+string(letters[5]))
				}
			} else if len(signal[i]) == 5 {
				if containsChars(signal[i], numbers[7]+string(letters[3])) {
					// map number 3
					numbers[3] = signal[i]
					letters[6] = extractNotFound(signal[i], numbers[7]+string(letters[3]))
				} else if false == containsChars(signal[i], string(letters[1])) {
					// map number 5
					numbers[5] = signal[i]
					letters[5] = extractNotFound(signal[i], numbers[3])
				} else if letters[2] != 0 && false == containsChars(signal[i], string(letters[2])) {
					numbers[2] = signal[i]
				}
			}
		}
	}

	return numbers
}

func countOutputDigits(entry string) int {
	data := strings.Split(entry, "|")
	_, output := strings.Fields(data[0]), strings.Fields(data[1])

	var outputDigits int
	for i := 0; i < len(output); i++ {
		_, ok := wireToInt[len(output[i])]
		if ok {
			outputDigits += 1
		}
	}
	return outputDigits
}

func getOutputSum(data []string) int {

	sum := make([]string, len(data), len(data))
	for lineIdx, line := range data {
		if len(line) > 0 {

			lineData := strings.Split(line, "|")
			signal, outputValues := strings.Fields(lineData[0]), strings.Fields(lineData[1])
			numbers := deduceNumbers(signal)

			for _, v := range outputValues {
				for x, y := range numbers {
					if len(v) == len(y) && containsChars(v, y) {
						sum[lineIdx] += strconv.Itoa(x)
					}
				}
			}
		}
	}

	var totalSum int
	for _, s := range sum {
		v, _ := strconv.Atoi(s)
		totalSum += v
	}
	return totalSum
}

func getOutputCount(data []string) int {
	var count int
	for i := 0; i < len(data); i++ {
		if len(data[i]) > 0 {
			count += countOutputDigits(data[i])
		}
	}

	return count
}

func part2(data []string) {

	sum := getOutputSum(data)

	fmt.Println("Total output is: ", sum)
}

func part1(data []string) {
	outputValues := getOutputCount(data)
	fmt.Println("Total output values are: ", outputValues)
}

/*
This challenge went horrible for me. I went into the rabbit hole of sticking with my first idea for the solution instead of trying to start over.
Lesson learnt!
*/

func main() {
	data := parseData(8)
	// part1(data)
	part2(data)
}
