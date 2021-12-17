package probe

import (
	"aoc2021/pkg/io"
	"aoc2021/pkg/numbers"
	"strconv"
	"strings"
)

type Probe struct {
	Position numbers.Vector2
	Velocity numbers.Vector2
}

type TargetArea struct {
	Start numbers.Vector2
	End   numbers.Vector2
}

type ProbeData struct {
	Target TargetArea
}

func ReadProbeData(file string) ProbeData {
	var pd ProbeData
	for _, line := range io.ReadLines(file) {
		data := strings.SplitN(line, ": ", 2)
		if data[0] == "target area" {
			data = strings.SplitN(data[1], ", ", 2)
			// Let's just assume fixed layout to make this quicker
			xd := strings.SplitN(data[0][2:], "..", 2)
			yd := strings.SplitN(data[1][2:], "..", 2)
			pd.Target.Start.X, _ = strconv.Atoi(xd[0])
			pd.Target.Start.Y, _ = strconv.Atoi(yd[0])
			pd.Target.End.X, _ = strconv.Atoi(xd[1])
			pd.Target.End.Y, _ = strconv.Atoi(yd[1])
		}
	}
	return pd
}

func (pd *ProbeData) FindMaxHeight() int {
	// For maximum height you'd want maximum velocity
	// Perfect Parabola: Since there's no drag/resistance, v_launch = v_impact (crossing y = 0)
	// -> Highest velocity goes from 0 to min_y in 1 step, x velocity doesn't matter
	// -> Vertical launch velocity is -min_y
	// -> height_max = v + (v - 1) + (v - 2) + ... + 1
	// -> height_max = v * (v + 1) / 2
	// Inserted and with swapped signs (orientation)...
	// -> height_max = -min_y * (-min_y - 1) / 2
	return -pd.Target.Start.Y * (-pd.Target.Start.Y - 1) / 2
}

func (pd *ProbeData) FindLaunchVelocities() []numbers.Vector2 {
	var launches []numbers.Vector2

	// Similar approach, but swapped axes, maximum width here:
	// reach = vx * (vx - 1) / 2
	// Now just find all within range first

	vx_min := 0 // Don't think we'd have to shoot to the left :P
	// Find the first vx that will actually reach the target area
	// This is basically a maximum height (now: width) thing again
	for ; vx_min*(vx_min+1)/2 < pd.Target.Start.X; vx_min++ {
	}
	vx_max := pd.Target.End.X // More would be a guaranteed overshoot

	// Again, assuming that anything slower/faster will just under-overshoot
	vy_min := pd.Target.Start.Y
	vy_max := -pd.Target.Start.Y

	for vy := vy_min; vy <= vy_max; vy++ {
		for vx := vx_min; vx <= vx_max; vx++ {
			valid := false
			for x, y, tvx, tvy := 0, 0, vx, vy; x <= pd.Target.End.X && y >= pd.Target.Start.Y; {
				if x >= pd.Target.Start.X && x <= pd.Target.End.X && y >= pd.Target.Start.Y && y <= pd.Target.End.Y {
					valid = true
					break
				}

				x += tvx
				y += tvy

				if tvx > 0 {
					tvx--
				} else if tvx < 0 {
					tvx++
				}
				tvy--
			}
			if valid {
				launches = append(launches, numbers.Vector2{X: vx, Y: vy})
			}
		}
	}

	return launches
}
