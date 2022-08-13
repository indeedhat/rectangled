package rectangled

import "fmt"

// Edge defines the location/position on a rectongle of an edge
type Edge uint8

// String implements fmt.Stringer
func (e Edge) String() string {
	switch e {
	case Top:
		return "Top"
	case Right:
		return "Right"
	case Bottom:
		return "Bottom"
	case Left:
		return "Left"

	default:
		return "Unknown"
	}
}

var _ fmt.Stringer = (*Edge)(nil)

const (
	Top Edge = iota
	Right
	Bottom
	Left
)

// EdgeCoordinates is an alias of rectangle
// that represents the coordinates of single Rectangle edge
// X,Y represent the starting (top or left) most point of the line
// W,Z represent the ending (bottom or right) most point of the line
type EdgeCoordinates Rectangle[any]
