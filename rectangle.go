package rectangled

// Rectongle represents points in space describing a rectangle
// top left is represented by X and Y
// bottom right is represented by W and Z in place of X2 and Y2 respectively
type Rectangle[T any] struct {
	ID T
	W  int
	X  int
	Y  int
	Z  int
}

// NewRectangle simply fills out the fields of a Rectangle struct
// It is just a convinience method to avoid explicit constructon
func NewRectangle[T any](id T, x, y, w, z int) Rectangle[T] {
	return Rectangle[T]{
		ID: id,
		W:  w,
		X:  x,
		Y:  y,
		Z:  z,
	}
}

// Overlaps checks if another rectangle overlaps with this one
func (rect Rectangle[T]) Overlaps(other Rectangle[T]) bool {
	if rect.X >= other.W || other.X >= rect.W {
		return false
	}

	if rect.Y >= other.Z || other.Y >= rect.Z {
		return false
	}

	return true
}

// Touches returns the side that this rectangle touches the given one
// It does not care about overlapps so an internal rectangle that has a
// touching edge will be found
func (rect Rectangle[T]) Touches(other Rectangle[T]) Edge {
	if (rect.Y == other.Z || rect.Y == other.Y) &&
		rect.X < other.W && rect.W > other.X {

		return Top
	}

	if (rect.W == other.X || rect.W == other.W) &&
		rect.Y < other.Z && rect.Z > other.Y {

		return Right
	}

	if (rect.Z == other.Y || rect.Z == other.Z) &&
		rect.X < other.W && rect.W > other.X {

		return Bottom
	}

	if (rect.X == other.W || rect.X == other.X) &&
		rect.Y < other.Z && rect.Z > other.Y {

		return Left
	}

	return UnknownEdge
}
