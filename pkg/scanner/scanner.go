package scanner

import (
	"aoc2021/pkg/io"
	"aoc2021/pkg/numbers"
	"strconv"
	"strings"
)

type Scan []numbers.Vector3
type Scanner struct {
	Pings    Scan
	Position numbers.Vector3
}
type Scanners []Scanner

func ReadScanners(file string) Scanners {
	scanner_id := 0
	scanners := Scanners{}
	for _, line := range io.ReadLines(file) {
		if len(line) == 0 {
			continue
		}
		if line[0] == '-' && line[1] == '-' {
			tmp := strings.Split(line, " ")
			if tmp[1] == "scanner" {
				scanner_id, _ = strconv.Atoi(tmp[2])
				scanners = append(scanners, Scanner{Pings: Scan{}})
			}
			continue
		}
		tmp := strings.SplitN(line, ",", 3)
		x, _ := strconv.Atoi(tmp[0])
		y, _ := strconv.Atoi(tmp[1])
		z, _ := strconv.Atoi(tmp[2])
		scanners[scanner_id].Pings = append(scanners[scanner_id].Pings, numbers.Vector3{X: x, Y: y, Z: z})
	}
	return scanners
}

func (s *Scanners) GetBeaconCount() int {
	beacons := []numbers.Vector3{}
	for _, scanner := range *s {
		for _, ping := range scanner.Pings {
			found := false
			for _, b := range beacons {
				if b.X == ping.X && b.Y == ping.Y && b.Z == ping.Z {
					found = true
					break
				}
			}
			if !found {
				beacons = append(beacons, ping)
			}
		}
	}
	return len(beacons)
}

func (s *Scanners) GetScannerSpan() int {
	max_distance := 0
	for i, s1 := range *s {
		for j, s2 := range *s {
			if i != j {
				distance := numbers.Abs(s2.Position.X-s1.Position.X) + numbers.Abs(s2.Position.Y-s1.Position.Y) + numbers.Abs(s2.Position.Z-s1.Position.Z)
				if distance > max_distance {
					max_distance = distance
				}
			}
		}
	}
	return max_distance
}

func (s *Scanners) Align() {
	scanner_count := len(*s)
	oriented := make(map[int]bool)
	oriented[0] = true

	// Since all sensors are guaranteed overlaps, just loop until we found them all
	for len(oriented) < scanner_count {
		for ref_scanner_id := 0; ref_scanner_id < scanner_count; ref_scanner_id++ {
			if !oriented[ref_scanner_id] {
				continue
			}

			ref_scanner := &(*s)[ref_scanner_id]

			for test_scanner_id := 0; test_scanner_id < scanner_count; test_scanner_id++ {
				if oriented[test_scanner_id] {
					continue
				}

				test_scanner := &(*s)[test_scanner_id]
				matches := 0
				rots := 0

				// Just brute-force rotate trial and error
				for x_rot := 0; x_rot < 4; x_rot++ {
					for y_rot := 0; y_rot < 4; y_rot++ {
						for z_rot := 0; z_rot < 4; z_rot++ {

							// Pick a reference point for the offset
							for _, ref_point := range ref_scanner.Pings {
								// Pick a potential match for the reference point
								for _, test_point := range test_scanner.Pings {
									ox := test_point.X - ref_point.X
									oy := test_point.Y - ref_point.Y
									oz := test_point.Z - ref_point.Z
									matches = 0
									// Now go through both sets of points...
									for _, match_point := range ref_scanner.Pings {
										for _, other_point := range test_scanner.Pings {
											// ... and see if they match with the given offset
											if other_point.X-ox == match_point.X && other_point.Y-oy == match_point.Y && other_point.Z-oz == match_point.Z {
												matches++
												break
											}
										}
									}
									// If we find at least 12 matches, this offset works
									if matches >= 12 {
										// Adjust all beacon positions to the first sensor's origin
										for i, _ := range test_scanner.Pings {
											test_scanner.Pings[i].X -= ox
											test_scanner.Pings[i].Y -= oy
											test_scanner.Pings[i].Z -= oz
										}
										// Save the scanner's position (=offset) for part 2
										test_scanner.Position.X = ox
										test_scanner.Position.Y = oy
										test_scanner.Position.Z = oz
										oriented[test_scanner_id] = true
										break
									}
								}
								if matches >= 12 {
									break
								}
							}

							if matches >= 12 {
								break
							}

							// Didn't find enough matches, roate around Z and try again
							for i, _ := range test_scanner.Pings {
								tmp := test_scanner.Pings[i].X
								test_scanner.Pings[i].X = test_scanner.Pings[i].Y
								test_scanner.Pings[i].Y = -tmp
							}
							rots++
						}
						if oriented[test_scanner_id] {
							break
						}

						// Didn't find enough matches, roate around Y and try again
						for i, _ := range test_scanner.Pings {
							tmp := test_scanner.Pings[i].X
							test_scanner.Pings[i].X = test_scanner.Pings[i].Z
							test_scanner.Pings[i].Z = -tmp
						}
						rots++
					}
					if oriented[test_scanner_id] {
						break
					}

					// Didn't find enough matches, roate around X and try again
					for i, _ := range test_scanner.Pings {
						tmp := test_scanner.Pings[i].Z
						test_scanner.Pings[i].Z = test_scanner.Pings[i].Y
						test_scanner.Pings[i].Y = -tmp
					}
					rots++
				}
				if oriented[test_scanner_id] {
					println("✔️ Aligned scanner", test_scanner_id, "using", ref_scanner_id, "with", matches, "matches after", rots, "rotations")
				} else {
					// This line was just for testing, getts a bit spammy...
					// println("❌ Failed to align scanner", test_scanner_id, "using", ref_scanner_id, "and", rots, "rotations")
				}
			}
		}
		println("Completion:", len(oriented)*100/scanner_count, "%")
	}
}
