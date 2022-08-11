package rectangled

// Edge defines the location/position on a rectongle of an edge
type Edge uint8

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

const (
	Top Edge = iota
	Right
	Bottom
	Left
	UnknownEdge
)
