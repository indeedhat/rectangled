package rectangled

// Set defines a group of rectangles
// wasnt sure what to call this, it was eitge Set or Murder
type Set[S, R any] struct {
	ID       S
	Children []Rectangle[R]
}

// NewSet fills out the fields of the set struct with the given types
func NewSet[S, R any](id S, children []Rectangle[R]) Set[S, R] {
	return Set[S, R]{
		ID:       id,
		Children: children,
	}
}
