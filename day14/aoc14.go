package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

type Poly2 struct {
	pairMap   map[string]string
	pairCount map[string]int
}

type Pair struct {
	index []int
}

type Poly struct {
	template string
	pairs    []string
	counter  map[string]int
}

func (p *Poly) Count() {
	for _, c := range p.template {
		p.counter[string(c)] += 1
	}
}

func (p *Poly) Calculate() int {
	fmt.Println(p.counter)
	min, max := p.counter[string(p.template[0])], 0
	for _, i := range p.counter {
		if i > max {
			max = i
		}
		if i < min {
			min = i
		}
	}

	fmt.Println(min, max)
	return max - min
}

func parseDataPart2(day int) *Poly2 {
	file, _ := ioutil.ReadFile(fmt.Sprintf("day%d/input.txt", day))
	// file, _ := ioutil.ReadFile(fmt.Sprintf("day%d/larger.txt", day))
	// file, _ := ioutil.ReadFile(fmt.Sprintf("day%d/testinput.txt", day))
	data := strings.Split(string(file), "\n")
	data = data[:len(data)-1]

	poly := &Poly2{
		make(map[string]string, 0),
		make(map[string]int, 0),
	}

	for _, l := range data[2:] {
		pair, point := "", ""
		fmt.Sscanf(l, "%s -> %s", &pair, &point)
		poly.pairMap[pair] = point
	}

	for i := 0; i <= len(data[0])-2; i++ {
		poly.pairCount[string(data[0][i:i+2])] += 1
	}

	fmt.Println(poly.pairCount)

	return poly
}

func parseData(day int) *Poly {
	file, _ := ioutil.ReadFile(fmt.Sprintf("day%d/input.txt", day))
	// file, _ := ioutil.ReadFile(fmt.Sprintf("day%d/larger.txt", day))
	// file, _ := ioutil.ReadFile(fmt.Sprintf("day%d/testinput.txt", day))
	data := strings.Split(string(file), "\n")
	data = data[:len(data)-1]

	poly := &Poly{
		data[0],
		make([]string, len(data[2:]), len(data[2:])),
		make(map[string]int, 0),
	}

	for i, l := range data[2:] {
		poly.pairs[i] = l

	}

	return poly
}

func (p *Poly2) GetResult() int {
	count := map[string]int{}
	for k, v := range p.pairCount {
		count[string(k[0])] += v
		count[string(k[1])] += v
	}
	count["N"] += 1
	count["B"] += 1

	for k := range count {
		count[k] /= 2
	}
	fmt.Println(count)

	min, max := math.MaxInt, 0
	for _, i := range count {
		if i > max {
			max = i
		}
		if i < min {
			min = i
		}
	}

	return max - min
}

func (p *Poly2) expandPair() {
	// expand each pair to its matching pairs
	pairs := make(map[string]int, 0)

	for k, n := range p.pairCount {
		match := p.pairMap[k]
		pairs[string(k[0])+match] += n
		pairs[match+string(k[1])] += n
	}
	p.pairCount = pairs
}

func part2(info *Poly2) {

	for i := 0; i < 40; i++ {
		info.expandPair()
	}
	fmt.Println(info.GetResult())
}

func part1(info *Poly) {
	for i := 0; i < 10; i++ {
		tmp := ""
		for l := 0; l < len(info.template)-1; l++ {
			for _, v := range info.pairs {
				if string(info.template[l:l+2]) == string(v[:2]) {
					tmp += string(info.template[l]) + string(v[len(v)-1])
				}
			}
		}
		info.template = tmp + string(info.template[len(info.template)-1])
	}
	info.Count()
	fmt.Println("Results is: ", info.Calculate())
}

/*
Part 1 was relativly straight forward.
Part 2 forced me to think about optimization. I came up with a solution that count pairs for each step so that we don't create a big string.
*/
func main() {
	// data := parseData(14)
	// part1(data)
	data := parseDataPart2(14)
	part2(data)
}
