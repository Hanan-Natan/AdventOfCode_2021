package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"strconv"
	"strings"
)

const (
	TypeLiteral string = "100"
	// TypeOperator        = "6"
)

const (
	SubLength   uint64 = 11
	TotalLength        = 15
)

type Header struct {
	Version uint64
}

type Literal struct {
	Pkt Header
	// Values []uint
	sum uint64
}

func newLiteral(ver uint64, rawData string) Literal {
	// ver, _ := strconv.ParseUint(rawData[0:3], 2, 16)
	// values := []uint{}
	sum := ""
	for i := 0; i < len(rawData); i += 5 {
		// v, _ := strconv.ParseUint(rawData[i+1:i+5], 2, 16)
		// values = append(values, uint(v))
		sum += rawData[i+1 : i+5]
		if rawData[i] == '0' {
			break
		}
	}
	v, _ := strconv.ParseUint(sum, 2, 64)

	return Literal{
		Pkt: Header{
			Version: ver,
		},
		// Values: values,
		sum: v,
	}
}

type Operator struct {
	Pkt     Header
	SubPkts Packets
}

type Packets struct {
	Literals  []Literal
	Operators []Operator
}

func (p Packets) sumVer() int {
	sumVer := 0
	for _, l := range p.Literals {
		sumVer += int(l.Pkt.Version)
	}
	for _, o := range p.Operators {
		sumVer += int(o.Pkt.Version)
		sumVer += o.SubPkts.sumVer()
	}

	return sumVer
}

func createPacket(data string) *Packets {
	literals := []Literal{}
	operators := []Operator{}

	ver, _ := strconv.ParseUint(data[0:3], 2, 16)
	switch tp := data[3:6]; tp {
	case TypeLiteral:
		fmt.Println("new literal:", tp, ver, data[6:])
		l := newLiteral(ver, data[6:])
		literals = append(literals, l)
	default:
		fmt.Println("New Operator ver:", tp, ver, data[6:])
		o := newOperator(ver, data[6:])
		operators = append(operators, o)
	}

	return &Packets{Literals: literals, Operators: operators}
}

func newOperator(ver uint64, s string) Operator {
	packets := Packets{}
	switch ltID := string(s[0]); ltID {
	case "0": // length of packets
		length, _ := strconv.ParseUint(s[1:TotalLength+1], 2, 32)
		fmt.Println("Operator length type 0, Totallength:", length, "rest is", s[1+TotalLength:1+TotalLength+length])
		packets = *createPacket(s[1+TotalLength : 1+TotalLength+length])
	case "1": // number of sub-packets
		length, _ := strconv.ParseUint(s[1:SubLength+1], 2, 16)
		fmt.Println("Operator sub-packet type 1, length:", length, "rest is", s[1+SubLength:])
		packets = *createPacket(s[1+SubLength:])
		fmt.Println(packets)
	}

	return Operator{
		Pkt: Header{
			Version: ver,
		},
		SubPkts: packets,
	}
}

func parseData(day int) *Packets {
	// file, _ := ioutil.ReadFile(fmt.Sprintf("day%d/input.txt", day))
	file, _ := ioutil.ReadFile(fmt.Sprintf("day%d/testinput.txt", day))
	// file, _ := ioutil.ReadFile(fmt.Sprintf("day%d/test1.txt", day))
	data := strings.Split(string(file), "\n")

	n := new(big.Int)
	n.SetString(data[0], 16)
	v := fmt.Sprintf("%b", n)

	packets := createPacket(v)

	return packets
}

func part1(data *Packets) {
	fmt.Println(data.sumVer())
}

func main() {
	data := parseData(16)
	part1(data)
	// fmt.Println(newLiteral(23, "101111111000101000"))
}
