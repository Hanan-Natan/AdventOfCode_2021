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

func parseData(day int) []int {
	file, _ := ioutil.ReadFile(fmt.Sprintf("day%d/input.txt", day))
	// file, _ := ioutil.ReadFile(fmt.Sprintf("day%d/testinput.txt", day))
	lines := strings.Split(string(file), "\n")

	data := strings.Split(lines[0], ",")

	res := make([]int, len(data), len(data))
	for i := 0; i < len(data); i++ {
		v, _ := strconv.Atoi(data[i])
		res[i] = v
	}

	return res
}

func getMinCost(data []int, isInc bool) int {
	var sum, heighest, lowest int
	lowest = data[0]
	for i := 0; i < len(data); i++ {
		sum += data[i]
		if data[i] > heighest {
			heighest = data[i]
		} else if data[i] < lowest {
			lowest = data[i]
		}
	}

	fmt.Println("Sum is: ", sum)
	fmt.Println("Avg is: ", sum/len(data))
	fmt.Println("High/Low num is: ", heighest, lowest)

	costData := []int{}
	for i := lowest; i <= heighest; i++ {
		var totalCost int
		for _, v := range data {
			if isInc {
				// to calculate do ((n*(n+1))/2)
				totalCost += (abs(v-i) * (abs(v-i) + 1)) / 2
			} else {
				totalCost += abs(v - i)
			}
		}
		costData = append(costData, totalCost)
	}

	return getLowestNum(costData)
}

func getLowestNum(data []int) int {
	lowestCost := data[0]
	for _, cost := range data {
		if cost == 0 {
			return cost
		} else if lowestCost > cost {
			lowestCost = cost
		}
	}
	return lowestCost
}

func main() {
	data := parseData(7)
	cost := getMinCost(data, true)
	fmt.Println("The cost is: ", cost)
}
