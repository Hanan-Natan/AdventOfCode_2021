package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	bitsLength = 12
)

func getRating(data []int64, shift int, mostCommon bool) int64 {
	if len(data) == 1 {
		return data[0]
	} else {
		if shift < 0 {
			panic("shift under 0")
		}
		ones := make([]int64, 0, len(data)/2)
		zeros := make([]int64, 0, len(data)/2)

		for _, line := range data {
			if ((line >> shift) & 1) != 0 {
				ones = append(ones, line)
			} else {
				zeros = append(zeros, line)
			}
		}
		if mostCommon {
			if len(ones) > len(zeros) {
				return getRating(ones, shift-1, mostCommon)
			} else if len(zeros) > len(ones) {
				return getRating(zeros, shift-1, mostCommon)
			}
			return getRating(ones, shift-1, mostCommon)
		} else {
			if len(ones) > len(zeros) {
				return getRating(zeros, shift-1, mostCommon)
			} else if len(zeros) > len(ones) {
				return getRating(ones, shift-1, mostCommon)
			}
			return getRating(zeros, shift-1, mostCommon)
		}
	}
}

func part2() {
	// file, _ := ioutil.ReadFile("day3/testdata.txt")
	file, _ := ioutil.ReadFile("day3/input.txt")
	lines := strings.Split(string(file), "\n")

	// convert the data to int slice
	data := make([]int64, len(lines)-1, len(lines)-1)
	for i, line := range lines {
		num, err := strconv.ParseInt(strings.Replace(line, "\n", "", -1), 2, 16)
		if nil == err && num != 0 {
			data[i] = num
		}
	}
	oxygen := getRating(data, bitsLength-1, true)
	fmt.Println("The oxygen rating is: ", oxygen)
	co2 := getRating(data, bitsLength-1, false)
	fmt.Println("The CO2 rating is: ", co2)

	fmt.Println("The life support rating is: ", oxygen*co2)
}

func part1() {
	file, _ := ioutil.ReadFile("day3/input.txt")
	// file, _ := ioutil.ReadFile("day3/testdata.txt")
	lines := strings.Split(string(file), "\n")
	bitsLength := 12

	onesCount := make([]int, bitsLength, bitsLength)
	var gamma, epsilon uint

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		num, err := strconv.ParseInt(strings.Replace(line, "\n", "", -1), 2, 16)
		if nil == err {
			for i := 0; i <= bitsLength; i++ {
				if ((num >> i) & 1) != 0 {
					onesCount[i] += 1
				}
			}
		}
	}

	for i := len(onesCount) - 1; i >= 0; i-- {
		if onesCount[i] > len(lines)/2 {
			gamma += (1 << i)
		} else {
			epsilon += (1 << i)
		}
	}

	fmt.Println("The power consumption is: ", gamma*epsilon)
}

func main() {
	part2()
}
