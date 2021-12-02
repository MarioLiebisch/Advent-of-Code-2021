package navigation

import "fmt"

type NavState struct {
	Horizontal int
	Depth      int
	Aim        int
}

func (v NavState) Checksum() int {
	return v.Horizontal * v.Depth
}

func (v NavState) String() string {
	return fmt.Sprintf("(%d,%d)", v.Horizontal, v.Depth)
}
