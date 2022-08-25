package rekt

import (
	"errors"
)

var (
	ErrNegativePositionInSet = errors.New("rectangles must have posotive coords")
)

// Set defines a group of rectangles
// wasnt sure what to call this, it was eitge Set or Murder
type Set[T any] struct {
	Rectangle[T]
	children []Rectangle[T]
}

// NewSet fills out the fields of the set struct with the given types
func NewSet[T any](id T, x, y int, children []Rectangle[T]) (*Set[T], error) {
	set := &Set[T]{
		Rectangle: Rectangle[T]{
			ID: id,
			X:  x,
			Y:  y,
		},
	}

	for _, rect := range children {
		if err := set.AddRectangle(rect); err != nil {
			return nil, err
		}
	}

	return set, nil
}

// AddRectangle adds a rectangle to the set and recalculates the sets dimensions
func (set *Set[T]) AddRectangle(rect Rectangle[T]) error {
	if err := rect.Validate(); err != nil {
		return err
	}

	if rect.X < 0 || rect.Y < 0 {
		return ErrNegativePositionInSet
	}

	set.children = append(set.children, rect)
	resizeSetToContent(set)

	return nil
}

// resizeSetToContent calculates and sets the bottom right corner and therefore size of
// a set based on the Rectangle's in it
func resizeSetToContent[T any](set *Set[T]) {
	for _, rect := range set.children {
		set.W = max(set.W, rect.X+rect.W)
		set.Z = max(set.Z, rect.Y+rect.Z)
	}
}

// OverlapsChildren returns a slice of child Rectangle's  of the provided Set that overlap with the
// bounding box of recievers Set
// this doesnt check if the bounding boxes of the sets overlap, that can be done with
//
//  set.Overlaps(target.Rectangle)
func (set *Set[T]) OverlapsChildren(target Set[T]) []Rectangle[T] {
	var overlapping []Rectangle[T]

	for _, rect := range target.children {
		if set.Overlaps(rect) {
			overlapping = append(overlapping, rect)
		}
	}

	return overlapping
}

// TouchesChildren returns a slice of child Rectangle's  of the provided Set that touch an edge
// against the bounding box of recievers Set
// this doesnt check if the bounding boxes of the sets touch, that can be done with
//
//  set.Tocches(target.Rectangle)
func (set *Set[T]) TouchesChildren(target Set[T]) []Rectangle[T] {
	var touching []Rectangle[T]

	for _, rect := range target.children {
		if set.Touches(rect) != nil {
			touching = append(touching, rect)
		}
	}

	return touching
}

// Children will return a copy of the child rectangle slice for this set
// Coordinates of the children will be relative to the Set space
func (set *Set[T]) Children() []Rectangle[T] {
	return set.children
}

// OffsetChildren will return a copy of the child rectangle slice for this set
// Coordinates of the children will be offset by the coordinates of the set making them
// relative to world space
func (set *Set[T]) OffsetChildren() []Rectangle[T] {
	var offsetChildren = make([]Rectangle[T], 0, len(set.children))

	for _, rect := range set.children {
		offsetChildren = append(offsetChildren, Rectangle[T]{
			ID: rect.ID,
			X:  rect.X + set.X,
			Y:  rect.Y + set.Y,
			W:  rect.W + set.X,
			Z:  rect.Z + set.Y,
		})
	}

	return offsetChildren
}

// ChildOnEdge finds the Rectangle closest to the priorityEdge within the Set
// Should more than one Rectangle be equally close then the one closest to the secondaryEdge
// will be picked
func (set *Set[T]) ChildOnEdge(priorityEdge, secondaryEdge Edge) *Rectangle[T] {
	var target *Rectangle[T]

	for i := range set.children {
		child := &set.children[i]

		if target == nil {
			target = child
			continue
		}

		if closer := findClosetsToEdge(target, child, priorityEdge); closer != nil {
			target = closer
		} else if closer := findClosetsToEdge(target, child, secondaryEdge); closer != nil {
			target = closer
		}
	}

	return target
}

// findClosetsToEdge compares the coords of the given Rectangles and returns the Rectangle
// that is closest to the given edge
//
// if they are equally close the nil will be returned
func findClosetsToEdge[T any](a, b *Rectangle[T], edge Edge) *Rectangle[T] {
	switch edge {
	case Top:
		fallthrough
	default:
		if a.Y < b.Y {
			return a
		} else if a.Y > b.Y {
			return b
		}
		return nil

	case Right:
		if a.W > b.W {
			return a
		} else if a.W < b.W {
			return b
		}
		return nil

	case Bottom:
		if a.Z > b.Z {
			return a
		} else if a.Z < b.Z {
			return b
		}
		return nil

	case Left:
		if a.X < b.X {
			return a
		} else if a.X > b.X {
			return b
		}
		return nil
	}
}
