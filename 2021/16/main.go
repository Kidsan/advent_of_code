package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

var hexValuesInBinary = map[string]string{
	"0": "0000",
	"1": "0001",
	"2": "0010",
	"3": "0011",
	"4": "0100",
	"5": "0101",
	"6": "0110",
	"7": "0111",
	"8": "1000",
	"9": "1001",
	"A": "1010",
	"B": "1011",
	"C": "1100",
	"D": "1101",
	"E": "1110",
	"F": "1111",
}

type binaryString []string

type packet struct {
	subPackets   []packet
	packetTypeId int
	version      int
	value        int
}

func (p packet) sumVersions() int {
	var sum int = p.version
	for _, packet := range p.subPackets {
		sum += packet.sumVersions()
	}
	return sum
}

func (b *binaryString) Pop(count int) binaryString {
	oldB := *b
	*b = (*b)[count:]
	return oldB[:count]
}

func (b *binaryString) PopAndDecodeBits(count int) int {
	leadBits := b.Pop(count)
	return leadBits.Decode()
}

func (b binaryString) Decode() int {
	return b.valueOf(0, len(b))
}

func (b binaryString) valueOf(start int, count int) int {
	bits := b[start : start+count]
	str := strings.Join(bits, "")
	num, _ := strconv.ParseInt(str, 2, 64)

	return int(num)
}

func (b *binaryString) Parse() packet {
	version := int(b.PopAndDecodeBits(3))
	packetTypeId := int(b.PopAndDecodeBits(3))

	switch packetTypeId {
	case 4:
		value := binaryString{}
		for i := "1"; i != "0"; {
			bits := b.Pop(5)
			value = append(value, bits[1:]...)
			i = bits[0]
		}
		parsed := int(value.Decode())
		return packet{
			version:      version,
			packetTypeId: packetTypeId,
			value:        parsed,
		}
	default:
		// parsing an operator packet
		lengthTypeId := b.Pop(1)[0]
		children := []packet{}

		switch lengthTypeId {
		case "0":
			lengthOfSubpacketString := b.PopAndDecodeBits(15)
			subPacketString := b.Pop(lengthOfSubpacketString)
			for len(subPacketString) > 0 {
				children = append(children, subPacketString.Parse())
			}
		case "1":
			numberOfChildren := b.PopAndDecodeBits(11)
			for i := 0; i < numberOfChildren; i++ {
				children = append(children, b.Parse())
			}
		default:
			panic(fmt.Errorf("unknown lengthTypeId: %v", lengthTypeId))
		}
		return packet{
			version:      version,
			packetTypeId: packetTypeId,
			subPackets:   children,
		}
	}
}

func hexToBinary(input string) string {
	result := ""
	parsed := strings.Split(input, "")
	for _, v := range parsed {
		result += hexValuesInBinary[v]
	}
	return result
}

func part1(input string) int {
	data := hexToBinary(input)

	binary := binaryString{}
	for _, v := range strings.Split(data, "") {
		binary = append(binary, strings.Split(v, "")...)
	}

	parsed := binary.Parse()
	return parsed.sumVersions()

}

func part2(input string) int {
	return 1
}

func main() {
	start := time.Now()
	input, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	content, err := ioutil.ReadAll(input)
	if err != nil {
		panic(err)
	}

	inputLists := strings.Split(string(content), "\n")
	fmt.Printf("Part One: %v (took %s)\n", part1(inputLists[0]), time.Since(start))
}
