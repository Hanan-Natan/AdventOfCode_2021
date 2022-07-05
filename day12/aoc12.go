package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Road struct {
	start string
	end   string
}

func parseData(day int) []Road {
	file, _ := ioutil.ReadFile(fmt.Sprintf("day%d/input.txt", day))
	// file, _ := ioutil.ReadFile(fmt.Sprintf("day%d/larger.txt", day))
	// file, _ := ioutil.ReadFile(fmt.Sprintf("day%d/testinput.txt", day))
	data := strings.Split(string(file), "\n")
	data = data[:len(data)-1]

	roads := make([]Road, len(data), len(data))
	for i, l := range data {
		roads[i] = Road{strings.Split(l, "-")[0], strings.Split(l, "-")[1]}
	}
	return roads
}

func findEdges(roads []Road) ([][]string, [][]string) {
	start, end := make([][]string, 0), make([][]string, 0)
	for r := range roads {
		if roads[r].start == "start" {
			ro := []string{roads[r].start, roads[r].end}
			start = append(start, ro)
		} else if roads[r].end == "start" {
			ro := []string{roads[r].end, roads[r].start}
			start = append(start, ro)
		} else if roads[r].end == "end" {
			endr := []string{roads[r].start, roads[r].end}
			end = append(end, endr)
		}
	}
	return start, end
}

func stringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func isInStore(item []string, store [][]string) bool {
	for i := range store {
		if stringSlicesEqual(item, store[i]) {
			return true
		}
	}
	return false
}

func duplicateCount(src []string, item string) int {
	var count int
	for _, v := range src {
		if item == v {
			count += 1
		}
	}
	return count
}

func hasLowerCaseDup(str []string) bool {
	// check if has duplicates
	for i := 0; i < len(str); i++ {
		if str[i] > "Z" {
			if duplicateCount(str, str[i]) == 2 {
				return true
			}
		}
	}

	return false
}

func isValid(str []string, postfix string, partTwo bool) bool {
	if postfix <= "Z" {
		// upper case
		return true
	}

	count := duplicateCount(str, postfix)

	if count == 0 {
		return true
	} else if count == 1 && partTwo {
		if hasLowerCaseDup(str) {
			return false
		}
		return true
	}

	return false
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func cleanupStore(store [][]string) [][]string {
	newStore := make([][]string, 0)
	for _, v := range store {
		if "end" == v[len(v)-1] {
			newStore = append(newStore, v)
		}
	}
	return newStore
}

func countConnections(li []string, roads []Road, allowDup bool) int {
	store := make([][]string, 0)
	store = append(store, li)
	for i := 0; i < len(store); i++ {
		if "end" == store[i][len(store[i])-1] {
			continue
		}

		for _, v := range roads {
			if v.start == "start" || v.end == "start" {
				continue
			}

			switch store[i][len(store[i])-1] {
			case v.start:
				if isValid(store[i], v.end, allowDup) && false == isInStore(append(store[i], v.end), store) {
					newli := make([]string, len(store[i])+1)
					copy(newli, store[i])
					newli = append(newli, v.end)
					store = append(store, newli)
				}
				break
			case v.end:
				if isValid(store[i], v.start, allowDup) && false == isInStore(append(store[i], v.start), store) {
					newli := make([]string, len(store[i])+1)
					copy(newli, store[i])
					newli = append(newli, v.start)
					store = append(store, newli)
				}
				break
			}
		}
	}
	return len(cleanupStore(store))
}

func part2(roads []Road) {
	starts, _ := findEdges(roads)
	total := 0
	for _, v := range starts {
		total += countConnections(v, roads, true)
	}
	fmt.Println("Total connections:", total)
}

func part1(roads []Road) {
	starts, _ := findEdges(roads)
	total := 0
	for _, v := range starts {
		total += countConnections(v, roads, false)
	}
	fmt.Println("Total connections:", total)

}

/*
A crude solution.
I don't think there's any algorithem or smart idea in this code.
I can't wait to see what others came up with...
*/
func main() {
	data := parseData(12)
	// part1(data)
	part2(data)
}
