package bits

import (
	"aoc2021/pkg/io"
	"log"
)

var hexbits map[rune]string = map[rune]string{
	'0': "0000",
	'1': "0001",
	'2': "0010",
	'3': "0011",
	'4': "0100",
	'5': "0101",
	'6': "0110",
	'7': "0111",
	'8': "1000",
	'9': "1001",
	'A': "1010",
	'B': "1011",
	'C': "1100",
	'D': "1101",
	'E': "1110",
	'F': "1111",
}

type BitsTransmission struct {
	Packet BitsPacket
}

type BitsPackets []BitsPacket

type BitsPacket struct {
	version int
	typeId  int
	value   int
	packets []BitsPacket
}

func HexToBits(hex string) string {
	bits := ""
	for _, c := range hex {
		bits += hexbits[c]
	}
	return bits
}

func BitsToDec(bits string) int {
	dec := 0
	for i := 0; i < len(bits); i++ {
		dec *= 2
		if bits[i] == '1' {
			dec++
		}
	}
	return dec
}

func LoadBits(file string) BitsTransmission {
	var bt BitsTransmission
	data := io.ReadLines(file)[0]
	bt.Packet, _ = ReadPacketHex(data)
	return bt
}

func ReadPacketHex(data string) (BitsPacket, int) {
	return ReadPacket(HexToBits(data))
}

func ReadPacket(bits string) (BitsPacket, int) {
	var p BitsPacket
	consumed := 0
	if len(bits) < 6 {
		log.Fatalln("Not enough bits in packet data!")
		return p, consumed
	}
	consumed += 6
	p.version = BitsToDec(bits[0:3])
	p.typeId = BitsToDec(bits[3:6])
	switch p.typeId {
	case 4: // literal
		value := ""
		for {
			part := bits[consumed : consumed+5]
			consumed += 5
			value += part[1:5]
			// Continue as long as the first bit is 1
			if part[0] == '0' {
				break
			}
		}
		p.value = BitsToDec(value)
	default: // operator
		lengthTypeId := BitsToDec(bits[consumed : consumed+1])
		consumed++

		// Determine content length and read/append sub packets
		if lengthTypeId == 0 {
			length := BitsToDec(bits[consumed : consumed+15])
			consumed += 15
			ec := 0
			// Continue until `length` bits were consumed
			for ec < length {
				sp, c := ReadPacket(bits[consumed+ec:])
				p.packets = append(p.packets, sp)
				ec += c
			}
			consumed += ec
		} else {
			count := BitsToDec(bits[consumed : consumed+11])
			consumed += 11
			// Continue until `count` packets were read
			for i := 0; i < count; i++ {
				sp, c := ReadPacket(bits[consumed:])
				p.packets = append(p.packets, sp)
				consumed += c
			}
		}

		// Perform the actual operation now that we have the sub packets
		switch p.typeId {
		case 0: // sum
			sum := 0
			for _, sp := range p.packets {
				sum += sp.value
			}
			p.value = sum
		case 1: // product
			prod := 1
			for _, sp := range p.packets {
				prod *= sp.value
			}
			p.value = prod
		case 2: // minimum
			min := p.packets[0].value
			for _, sp := range p.packets {
				if sp.value < min {
					min = sp.value
				}
			}
			p.value = min
		case 3: // maximum
			max := p.packets[0].value
			for _, sp := range p.packets {
				if sp.value > max {
					max = sp.value
				}
			}
			p.value = max
		case 5: // greater than
			if p.packets[0].value > p.packets[1].value {
				p.value = 1
			} else {
				p.value = 0
			}
		case 6: // less than
			if p.packets[0].value < p.packets[1].value {
				p.value = 1
			} else {
				p.value = 0
			}
		case 7: // equal to
			if p.packets[0].value == p.packets[1].value {
				p.value = 1
			} else {
				p.value = 0
			}
		}
	}
	return p, consumed
}

func (p *BitsPacket) GetVersionSum() int {
	sum := p.version
	for _, sp := range p.packets {
		sum += sp.GetVersionSum()
	}
	return sum
}

func (p *BitsPacket) GetValue() int {
	return p.value
}
