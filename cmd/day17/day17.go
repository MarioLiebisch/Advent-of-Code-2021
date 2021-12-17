package main

import "aoc2021/pkg/probe"

func main() {
	datas := probe.ReadProbeData("./data/sample-17.txt")
	datai := probe.ReadProbeData("./data/input-17.txt")

	println("Solution 1 - Sample:", datas.FindMaxHeight())
	println("Solution 1 - Input:", datai.FindMaxHeight())

	println("Solution 2 - Sample:", len(datas.FindLaunchVelocities()))
	println("Solution 2 - Input:", len(datai.FindLaunchVelocities()))
}
