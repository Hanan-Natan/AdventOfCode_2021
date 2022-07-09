package main

import (
	"fmt"
	"io/ioutil"
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
	header Header
	// Values []uint
	sum uint64
}

type RawData struct {
	d string
}

func (r *RawData) popVersion() uint64 {
	version, _ := strconv.ParseUint(r.d[0:3], 2, 16)
	r.d = r.d[3:]
	return version
}

func (r *RawData) popType() string {
	t := r.d[0:3]
	r.d = r.d[3:]
	return t
}

func (r *RawData) reduceLen(size int) {
	r.d = r.d[size:]
}

func (r *RawData) popLengthTypeID() uint64 {
	ltID, _ := strconv.ParseUint(string(r.d[0]), 2, 8)
	r.d = r.d[1:]
	return ltID
}

func (r *RawData) popTotalLength() uint64 {
	// fmt.Println("Panic?", r.d)
	length, _ := strconv.ParseUint(r.d[:TotalLength], 2, 32)
	r.d = r.d[TotalLength:]
	return length
}

func (r *RawData) popNumberOfPackets() uint64 {
	numOfSubPackets, _ := strconv.ParseUint(r.d[:SubLength], 2, 16)
	r.d = r.d[SubLength:]
	return numOfSubPackets
}

func newLiteral(ver uint64, rawData string) (int, Literal) {
	// values := []uint{}
	sum := ""
	literalLen := 0
	for i := 0; i < len(rawData); i += 5 {
		sum += rawData[i+1 : i+5]
		if rawData[i] == '0' {
			literalLen = i + 5
			// literalLen = i  <-- THE fing BUG: This is part of the bag that took me so much time to uncover!!!
			break
		}
	}
	v, _ := strconv.ParseUint(sum, 2, 64)

	// literalLen = literalLen*5 + 5 <-- MY BUG: This took away 12 houres of my life!!! It works with the test cases but not with the real input data.
	return literalLen, Literal{
		header: Header{ver},
		sum:    v,
	}
}

type Operator struct {
	header  Header
	SubPkts Packets
}

type Packets struct {
	pktType   string
	literals  []Literal
	operators []Operator
}

func (p *Packets) sumVer() int {
	sumVer := 0
	for _, l := range p.literals {
		sumVer += int(l.header.Version)
	}
	for _, o := range p.operators {
		sumVer += int(o.header.Version)
		sumVer += o.SubPkts.sumVer()
	}

	return sumVer
}

func (p *Packets) createPackets(data *RawData) {

	ver := data.popVersion()
	switch tp := data.popType(); tp {
	case TypeLiteral:
		p.pktType = "Lit"
		fmt.Printf("Ver: %d, %s Packet, Data: %s\n", ver, "Literal", data.d)
		len, l := newLiteral(ver, data.d)
		fmt.Printf("\tLiteral value is: %d, len is: %d\n", l.sum, len)
		p.literals = append(p.literals, l)
		data.reduceLen(len)
	default:
		p.pktType = "Op"
		// fmt.Printf("Ver: %d, %s Packet, Data: %s\n", ver, "Operator", data.d)
		fmt.Printf("Ver: %d, %s Packet\n", ver, "Operator")
		p.operators = append(p.operators, newOperator(ver, data))
		// p.Operators = append(p.Operators, o)
	}

}

func newOperator(ver uint64, data *RawData) Operator {
	o := Operator{
		header:  Header{ver},
		SubPkts: Packets{},
	}
	switch ltID := data.popLengthTypeID(); ltID {
	case 0: // length of packets
		length := data.popTotalLength()
		// fmt.Println("Operator length type, Totallength:", length, "rest is", data.d)
		fmt.Println("\tOperator length type, Totallength:", length)
		if int(length) > len(data.d) {
			fmt.Println(length, data.d)
			panic("Wrong thing here")
		}
		tot := len(data.d) - int(length)
		for tot < len(data.d) {
			o.SubPkts.createPackets(data)
		}
		if tot > len(data.d) {
			fmt.Println(tot, len(data.d))
			panic("Tot is more than data")
		}
	case 1: // number of sub-packets
		numOfSubPackets := data.popNumberOfPackets()
		// fmt.Println("Operator sub-packet type, length:", numOfSubPackets, "rest is", data.d)
		fmt.Println("\tOperator sub-packet type, length:", numOfSubPackets)
		// o.SubPkts.data = RawData{data[1+SubLength:]}
		for i := 0; i < int(numOfSubPackets); i++ {
			o.SubPkts.createPackets(data)
		}
	}

	return o
}

func parseData(day int) *RawData {
	file, _ := ioutil.ReadFile(fmt.Sprintf("day%d/input.txt", day))
	// file, _ := ioutil.ReadFile(fmt.Sprintf("day%d/testinput.txt", day))
	// file, _ := ioutil.ReadFile(fmt.Sprintf("day%d/test1.txt", day))
	data := strings.Split(string(file), "\n")

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

	// n := new(big.Int)
	// n.SetString(data[0], 16)
	// v := n.Text(2)
	// fmt.Println(v[:12], len(v))
	// v = "000000000" + v
	// t, _ := strconv.ParseInt(string(data[0][0]), 16, 64)
	// if t <= 6 {
	// 	v = "0" + v
	// 	if t <= 3 {
	// 		v = "0" + v
	// 		if t <= 1 {
	// 			v = "0" + v
	// 			if t == 0 {
	// 				v = "0" + v
	// 			}
	// 		}
	// 	}

	// }
	// fmt.Println("Payload is", v)

	// v = "11101110000000001101010000001100100000100011000001100000"
	// packets := createPacket(v)

	return &RawData{v}
}

func part1(d *RawData) {

	fmt.Println("Payload len:", len(d.d))
	fmt.Println("Payload first 40:", d.d[:50])
	fmt.Println("Payload last 10:", d.d[len(d.d)-10:len(d.d)])
	p := Packets{}
	p.createPackets(d)
	fmt.Println(p.sumVer())
	fmt.Println("Left out data:", d.d)
}

func main() {
	data := parseData(16)
	part1(data)
	// fmt.Println(newLiteral(23, "101111111000101000"))
}
