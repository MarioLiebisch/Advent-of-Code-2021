package main

import (
	"aoc2021/pkg/bits"
)

func main() {
	samples1 := []string{
		"8A004A801A8002F478",
		"620080001611562C8802118E34",
		"C0015000016115A2E0802F182340",
		"A0016C880162017C3686B18A3D4780",
	}
	samples2 := []string{
		"C200B40A82",
		"04005AC33890",
		"880086C3E88112",
		"CE00C43D881120",
		"D8005AC2A8F0",
		"F600BC2D8F",
		"9C005AC2F8F0",
		"9C0141080250320F1802104A08",
	}

	for _, sample := range samples1 {
		p, _ := bits.ReadPacketHex(sample)
		println("Solution 1 - Sample", sample, "->", p.GetVersionSum())
	}

	println()

	for _, sample := range samples2 {
		p, _ := bits.ReadPacketHex(sample)
		println("Solution 2 - Sample", sample, "->", p.Evaluate())
	}

	println()

	input := bits.LoadBits("./data/input-16.txt")
	println("Solution 1 - Input:", input.Packet.GetVersionSum())
	println("Solution 2 - Input:", input.Packet.Evaluate())
}
