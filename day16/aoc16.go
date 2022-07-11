package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	TypeSum     string = "000"
	TypeProduct string = "001"
	TypeMin     string = "010"
	TypeMax     string = "011"
	TypeLiteral string = "100"
	TypeGt      string = "101"
	TypeLt      string = "110"
	TypeEq      string = "111"
)

const (
	SubLength   uint64 = 11
	TotalLength        = 15
)

type RawData struct {
	d string
}

func (r *RawData) popVersion() int {
	version, _ := strconv.ParseInt(r.d[0:3], 2, 16)
	r.d = r.d[3:]
	return int(version)
}

func (r *RawData) popType() string {
	t := r.d[0:3]
	r.d = r.d[3:]
	return t
}

func (r *RawData) reduceLen(size int) {
	r.d = r.d[size:]
}

func (r *RawData) popLengthTypeID() int {
	ltID, _ := strconv.ParseUint(string(r.d[0]), 2, 8)
	r.d = r.d[1:]
	return int(ltID)
}

func (r *RawData) popTotalLength() int {
	length, _ := strconv.ParseUint(r.d[:TotalLength], 2, 32)
	r.d = r.d[TotalLength:]
	return int(length)
}

func (r *RawData) popNumberOfPackets() int {
	numOfSubPackets, _ := strconv.ParseInt(r.d[:SubLength], 2, 16)
	r.d = r.d[SubLength:]
	return int(numOfSubPackets)
}

type Packet struct {
	pktType string
	version int
	value   int
	data    []*Packet
}

func createPackets(ver int, rd *RawData) *Packet {

	p := &Packet{
		pktType: "",
		version: ver,
		value:   0,
		data:    []*Packet{},
	}
	switch p.pktType = rd.popType(); p.pktType {
	case TypeLiteral:
		len, value := newLiteral(rd.d)
		// fmt.Printf("\tLiteral value is: %d, len is: %d\n", value, len)
		p.value = value
		rd.reduceLen(len)
	default:
		switch ltID := rd.popLengthTypeID(); ltID {
		case 0: // length of packets
			length := rd.popTotalLength()
			tot := len(rd.d) - int(length)
			for tot < len(rd.d) {
				p.data = append(p.data, createPackets(rd.popVersion(), rd))
			}
		case 1: // number of sub-packets
			numOfSubPackets := rd.popNumberOfPackets()
			for i := 0; i < int(numOfSubPackets); i++ {
				p.data = append(p.data, createPackets(rd.popVersion(), rd))
			}
		}
	}

	return p
}

func newLiteral(rawData string) (int, int) {
	sum := ""
	literalLen := 0
	for i := 0; i < len(rawData); i += 5 {
		sum += rawData[i+1 : i+5]
		if rawData[i] == '0' {
			literalLen = i + 5
			break
		}
	}
	v, _ := strconv.ParseUint(sum, 2, 64)

	return literalLen, int(v)
}

func doMath(op string, values []int) int {
	result := 0
	for i, v := range values {
		if i == 0 {
			result = v
			continue
		}
		switch op {
		case "+":
			result += v
		case "*":
			result *= v
		case "min":
			if result > v {
				result = v
			}
		case "max":
			if result < v {
				result = v
			}
		case "lt":
			if result < v {
				return 1
			} else {
				return 0
			}
		case "gt":
			if result > v {
				return 1
			} else {
				return 0
			}
		case "eq":
			if result == v {
				return 1
			} else {
				return 0
			}
		}
	}

	return result
}

func (p *Packet) getValues() []int {
	values := make([]int, len(p.data), len(p.data))
	for i, v := range p.data {
		values[i] = v.doOperations()
	}
	return values
}

func (p *Packet) doOperations() int {
	value := 0
	switch p.pktType {
	case TypeLiteral:
		value = p.value
	case TypeSum:
		value = doMath("+", p.getValues())
	case TypeProduct:
		value = doMath("*", p.getValues())
	case TypeMin:
		value = doMath("min", p.getValues())
	case TypeMax:
		value = doMath("max", p.getValues())
	case TypeGt:
		value = doMath("gt", p.getValues())
	case TypeLt:
		value = doMath("lt", p.getValues())
	case TypeEq:
		value = doMath("eq", p.getValues())

	}

	return value
}

func parseData(day int) *RawData {
	file, _ := ioutil.ReadFile(fmt.Sprintf("day%d/input.txt", day))
	// file, _ := ioutil.ReadFile(fmt.Sprintf("day%d/testinput.txt", day))
	data := strings.Split(string(file), "\n")

	// This for sure could have been more efficient by converting to a byte array.
	v := ""
	for _, l := range data[0] {
		switch l {
		case '0':
			v += "0000"
		case '1':
			v += "0001"
		case '2':
			v += "0010"
		case '3':
			v += "0011"
		case '4':
			v += "0100"
		case '5':
			v += "0101"
		case '6':
			v += "0110"
		case '7':
			v += "0111"
		case '8':
			v += "1000"
		case '9':
			v += "1001"
		case 'A':
			v += "1010"
		case 'B':
			v += "1011"
		case 'C':
			v += "1100"
		case 'D':
			v += "1101"
		case 'E':
			v += "1110"
		case 'F':
			v += "1111"
		}
	}

	return &RawData{v}
}

func (p *Packet) Print() {
	fmt.Printf("Type: %s, Version: %d, Value: %d\n", p.pktType, p.version, p.value)
	for _, i := range p.data {
		i.Print()
	}
}

func (p *Packet) SumVer() int {
	sum := p.version
	for _, i := range p.data {
		sum += i.SumVer()
	}
	return sum
}

func part1(d *RawData) {
	p := createPackets(d.popVersion(), d)
	fmt.Println("The sum is:", p.SumVer())
}

func part2(d *RawData) {
	p := createPackets(d.popVersion(), d)
	// p.Print()
	fmt.Println("The operation result is:", p.doOperations())
}

func main() {
	data := parseData(16)
	// part1(data)
	part2(data)
}
