package rekt

import (
	"errors"
)

var (
	ErrZoroArea  = errors.New("rectangles must have an area")
	ErrBadPoints = errors.New("rectangle has invalid coords")
)

// Rectongle represents points in space describing a rectangle
// top left is represented by X and Y
// bottom right is represented by W and Z in place of X2 and Y2 respectively
type Rectangle[T any] struct {
	ID T
	// top left
	X int
	Y int
	// bottom right
	W int
	Z int
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

// Offset returns a copy of the Rectangle with its coords offset by the x,y of the target Rectangle
func (rect Rectangle[T]) Offset(target Rectangle[T]) Rectangle[T] {
	rect.X += target.X
	rect.W += target.X
	rect.Y += target.Y
	rect.Z += target.Y

	return rect
}

// Area calculates the area of the rectangle
func (rect Rectangle[T]) Area() int {
	return abs((rect.W - rect.X) * (rect.Z - rect.Y))
}

// Width returns the calculated width of the rectangle
func (rect Rectangle[T]) Width() int {
	return abs(rect.W - rect.X)
}

// Height returns the calculated height of the rectangle
func (rect Rectangle[T]) Height() int {
	return abs(rect.Z - rect.Y)
}

// Overlaps checks if target rectangle overlaps with this one
func (rect Rectangle[T]) Overlaps(target Rectangle[T]) bool {
	if rect.X >= target.W || target.X >= rect.W {
		return false
	}

	if rect.Y >= target.Z || target.Y >= rect.Z {
		return false
	}

	return true
}

// OverlappingArea returns (if any) the bounding box of the area where the rectangle overlaps
// with the target rectangle
// nil will be returned if the two rectangles do not overlap
func (rect Rectangle[T]) OverlappingArea(target Rectangle[T]) *Rectangle[T] {
	if !rect.Overlaps(target) {
		return nil
	}

	overlap := NewRectangle(
		// target.ID is chosen for the ID over rect as in most cases the reciever rectangle
		// is more likely to be explicit than the target rectangle, i believe it to be a
		// more useful identiier in most situations
		target.ID,
		max(rect.X, target.X),
		max(rect.Y, target.Y),
		min(rect.W, target.W),
		min(rect.Z, target.Z),
	)

	return &overlap
}

// Touches returns a slice of sides in which the reciever Rectangle touches the target
// It does not care about overlapps so an internal rectangle that has a
// touching edge will be found
func (rect Rectangle[T]) Touches(target Rectangle[T]) []Edge {
	var edges []Edge
	if touchesTop(rect, target) {
		edges = append(edges, Top)
	}

	if touchesRight(rect, target) {
		edges = append(edges, Right)
	}

	if touchesBottom(rect, target) {
		edges = append(edges, Bottom)
	}

	if touchesLeft(rect, target) {
		edges = append(edges, Left)
	}

	return edges
}

// TouchCoordinates returns the coordinates of the touch area/line between two rectangles
// if there is no touch on the given edge nil will be returned
func (rect Rectangle[T]) TouchCoordinates(target Rectangle[T], edge Edge) *EdgeCoordinates[T] {
	switch edge {
	case Top:
		if touchesTop(rect, target) {
			return &EdgeCoordinates[T]{
				ID: target.ID,
				X:  max(rect.X, target.X),
				Y:  rect.Y,
				W:  min(rect.W, target.W),
				Z:  rect.Y,
			}
		}

	case Right:
		if touchesRight(rect, target) {
			return &EdgeCoordinates[T]{
				ID: target.ID,
				X:  rect.W,
				Y:  max(rect.Y, target.Y),
				W:  rect.W,
				Z:  min(rect.Z, target.Z),
			}
		}

	case Bottom:
		if touchesBottom(rect, target) {
			return &EdgeCoordinates[T]{
				ID: target.ID,
				X:  max(rect.X, target.X),
				Y:  rect.Z,
				W:  min(rect.W, target.W),
				Z:  rect.Z,
			}
		}

	case Left:
		if touchesLeft(rect, target) {
			return &EdgeCoordinates[T]{
				ID: target.ID,
				X:  rect.X,
				Y:  max(rect.Y, target.Y),
				W:  rect.X,
				Z:  min(rect.Z, target.Z),
			}
		}
	}

	return nil
}

// touchesTop checks if the top of rect touches the top or bottom of target
func touchesTop[T any](rect, target Rectangle[T]) bool {
	return (rect.Y == target.Z || rect.Y == target.Y) &&
		rect.X < target.W && rect.W > target.X
}

// touchesRight checks if the right of rect touches the right or left of target
func touchesRight[T any](rect, target Rectangle[T]) bool {
	return (rect.W == target.X || rect.W == target.W) &&
		rect.Y < target.Z && rect.Z > target.Y
}

// touchesBottom checks if the bottom of rect touches the bottom or top of target
func touchesBottom[T any](rect, target Rectangle[T]) bool {
	return (rect.Z == target.Y || rect.Z == target.Z) &&
		rect.X < target.W && rect.W > target.X
}

// touchesLeft checks if the left of rect touches the right or left of target
func touchesLeft[T any](rect, target Rectangle[T]) bool {
	return (rect.X == target.W || rect.X == target.X) &&
		rect.Y < target.Z && rect.Z > target.Y
}

// Validate checks if the rectangle is valid
// - Area of the rectangle must not be 0
// - X,Y must be top left
// - W,Z must be bottom right
func (rect Rectangle[T]) Validate() error {
	if rect.Area() == 0 {
		return ErrZoroArea
	}

	if rect.X > rect.W || rect.Y > rect.Z {
		return ErrBadPoints
	}

	return nil
}
