package rekt_test

import (
	"fmt"
	"testing"

	"github.com/indeedhat/rekt"
	"github.com/stretchr/testify/require"
)

var rectangleOverlapTouchTests = []struct {
	rect1    rekt.Rectangle[string]
	rect2    rekt.Rectangle[string]
	overlaps bool
	touch1   []rekt.Edge
	touch2   []rekt.Edge
}{
	{
		rekt.NewRectangle("above", 0, 0, 10, 10),
		rekt.NewRectangle("below", 0, 20, 10, 30),
		false,
		nil,
		nil,
	},
	{
		rekt.NewRectangle("left", 0, 0, 10, 10),
		rekt.NewRectangle("right", 20, 0, 30, 10),
		false,
		nil,
		nil,
	},
	{
		rekt.NewRectangle("touching-above", 0, 0, 10, 10),
		rekt.NewRectangle("touching-below", 0, 10, 10, 20),
		false,
		[]rekt.Edge{rekt.Bottom},
		[]rekt.Edge{rekt.Top},
	},
	{
		rekt.NewRectangle("touching-left", 0, 0, 10, 10),
		rekt.NewRectangle("touching-right", 10, 0, 20, 10),
		false,
		[]rekt.Edge{rekt.Right},
		[]rekt.Edge{rekt.Left},
	},
	{
		rekt.NewRectangle("full-overlap", 0, 0, 10, 10),
		rekt.NewRectangle("full-overlap", 0, 0, 10, 10),
		true,
		[]rekt.Edge{rekt.Top, rekt.Right, rekt.Bottom, rekt.Left},
		[]rekt.Edge{rekt.Top, rekt.Right, rekt.Bottom, rekt.Left},
	},
	{
		rekt.NewRectangle("inside-of", 5, 5, 10, 10),
		rekt.NewRectangle("surrounding", 0, 0, 15, 15),
		true,
		nil,
		nil,
	},
	{
		rekt.NewRectangle("top-left", 0, 0, 10, 10),
		rekt.NewRectangle("top-left-surrounding", 0, 0, 15, 15),
		true,
		[]rekt.Edge{rekt.Top, rekt.Left},
		[]rekt.Edge{rekt.Top, rekt.Left},
	},
	{
		rekt.NewRectangle("bottom-right", 5, 5, 15, 15),
		rekt.NewRectangle("bottom-right-surrounding", 0, 0, 15, 15),
		true,
		[]rekt.Edge{rekt.Right, rekt.Bottom},
		[]rekt.Edge{rekt.Right, rekt.Bottom},
	},
	{
		rekt.NewRectangle("overlap-above", 0, 0, 10, 10),
		rekt.NewRectangle("overlap-below", 0, 9, 10, 19),
		true,
		[]rekt.Edge{rekt.Right, rekt.Left},
		[]rekt.Edge{rekt.Right, rekt.Left},
	},
	{
		rekt.NewRectangle("overlap-left", 0, 0, 10, 10),
		rekt.NewRectangle("overlap-right", 9, 0, 19, 10),
		true,
		[]rekt.Edge{rekt.Top, rekt.Bottom},
		[]rekt.Edge{rekt.Top, rekt.Bottom},
	},
	{
		rekt.NewRectangle("bottom-right-corner", 0, 0, 10, 10),
		rekt.NewRectangle("top-left corner", 9, 9, 19, 19),
		true,
		nil,
		nil,
	},
	{
		rekt.NewRectangle("bottom-left-corner", 9, 0, 19, 10),
		rekt.NewRectangle("top-right corner", 0, 9, 10, 19),
		true,
		nil,
		nil,
	},
}

func TestRectangleOverlap(t *testing.T) {
	for _, testCase := range rectangleOverlapTouchTests {
		t.Run(testCase.rect1.ID, func(t *testing.T) {
			require.Equal(t, testCase.overlaps, testCase.rect1.Overlaps(testCase.rect2))
		})
		t.Run(testCase.rect2.ID, func(t *testing.T) {
			require.Equal(t, testCase.overlaps, testCase.rect2.Overlaps(testCase.rect1))
		})
	}
}

func TestRectangleTouches(t *testing.T) {
	for _, testCase := range rectangleOverlapTouchTests {
		t.Run(testCase.rect1.ID, func(t *testing.T) {
			require.Equal(t, testCase.touch1, testCase.rect1.Touches(testCase.rect2))
		})
		t.Run(testCase.rect2.ID, func(t *testing.T) {
			require.Equal(t, testCase.touch2, testCase.rect2.Touches(testCase.rect1))
		})
	}
}

var rectangleOffsetTests = []struct {
	rect   rekt.Rectangle[string]
	target rekt.Rectangle[string]
}{
	{
		rekt.NewRectangle("pos offset", 0, 0, 10, 10),
		rekt.NewRectangle("pos offset", 10, 10, 0, 0),
	},
	{
		rekt.NewRectangle("neg offset", 0, 0, 10, 10),
		rekt.NewRectangle("neg offset", -10, -10, 0, 0),
	},
	{
		rekt.NewRectangle("zero offset", 0, 0, 10, 10),
		rekt.NewRectangle("zero offset", 0, 0, 0, 0),
	},
}

func TestRectangleOffset(t *testing.T) {
	for _, testCase := range rectangleOffsetTests {
		t.Run(testCase.rect.ID, func(t *testing.T) {
			offset := testCase.rect.Offset(testCase.target)

			require.Equal(t, testCase.rect.X+testCase.target.X, offset.X)
			require.Equal(t, testCase.rect.W+testCase.target.X, offset.W)
			require.Equal(t, testCase.rect.Y+testCase.target.Y, offset.Y)
			require.Equal(t, testCase.rect.Z+testCase.target.X, offset.Z)
		})
	}
}

var rectangleAreaTests = []struct {
	rect     rekt.Rectangle[string]
	expected int
}{
	{rekt.NewRectangle("zero height", 0, 0, 10, 0), 0},
	{rekt.NewRectangle("zero width", 0, 0, 0, 10), 0},
	{rekt.NewRectangle("zero size", 0, 0, 0, 0), 0},
	{rekt.NewRectangle("positive area", 0, 0, 10, 10), 100},
	{rekt.NewRectangle("negative area", 0, 0, -10, -10), 100},
}

func TestRectangleArea(t *testing.T) {
	for _, testCase := range rectangleAreaTests {
		t.Run(testCase.rect.ID, func(t *testing.T) {
			require.Equal(t, testCase.expected, testCase.rect.Area())
		})
	}
}

var rectangleValidityTests = []struct {
	rect     rekt.Rectangle[string]
	expected error
}{
	{rekt.NewRectangle("zero height", 0, 0, 10, 0), rekt.ErrZoroArea},
	{rekt.NewRectangle("zero width", 0, 0, 0, 10), rekt.ErrZoroArea},
	{rekt.NewRectangle("zero size", 0, 0, 0, 0), rekt.ErrZoroArea},
	{rekt.NewRectangle("flipped points", 0, 0, -10, -10), rekt.ErrBadPoints},
	{rekt.NewRectangle("valid", 0, 0, 10, 10), nil},
	{rekt.NewRectangle("valid in negative", -10, -10, 0, 0), nil},
}

func TestRectangleValidate(t *testing.T) {
	for _, testCase := range rectangleValidityTests {
		t.Run(testCase.rect.ID, func(t *testing.T) {
			if testCase.expected == nil {
				require.Nil(t, testCase.rect.Validate())
			} else {
				require.ErrorIs(t, testCase.expected, testCase.rect.Validate())
			}
		})
	}
}

var rectangleOverlappingAreaTests = []struct {
	rect1    rekt.Rectangle[string]
	rect2    rekt.Rectangle[string]
	expected *rekt.Rectangle[string]
}{
	{
		rekt.NewRectangle("overlap-1", 0, 0, 10, 10),
		rekt.NewRectangle("overlap-2", 5, 5, 15, 15),
		&rekt.Rectangle[string]{"overlap-2", 5, 5, 10, 10},
	},
	{
		rekt.NewRectangle("overlap-2", 5, 5, 15, 15),
		rekt.NewRectangle("overlap-1", 0, 0, 10, 10),
		&rekt.Rectangle[string]{"overlap-1", 5, 5, 10, 10},
	},
	{
		rekt.NewRectangle("no-overlap-1", 0, 0, 10, 10),
		rekt.NewRectangle("no-overlap-2", 10, 10, 20, 20),
		nil,
	},
}

func TestRectangleOverlappingArea(t *testing.T) {
	for _, testCase := range rectangleOverlappingAreaTests {
		var testName = fmt.Sprintf("%s -> %s", testCase.rect1.ID, testCase.rect2.ID)

		t.Run(testName, func(t *testing.T) {
			overlap := testCase.rect1.OverlappingArea(testCase.rect2)

			if testCase.expected == nil {
				require.Nil(t, overlap)
				return
			}

			require.NotNil(t, overlap)
			require.Equal(t, *overlap, *testCase.expected)
		})
	}
}

var rectangleWidthHeightTests = []struct {
	rect           rekt.Rectangle[string]
	expectedWidth  int
	expectedHeight int
}{
	{rekt.NewRectangle("simple", 0, 0, 10, 10), 10, 10},
	{rekt.NewRectangle("negative", -20, -20, -10, -10), 10, 10},
	{rekt.NewRectangle("negative->posative", -10, -10, 10, 10), 20, 20},
	// although this requires a Rectangle that would not pass the Validate checks its worth testing imo
	{rekt.NewRectangle("zero-size", 0, 0, 0, 0), 0, 0},
}

func TestRectangleWidth(t *testing.T) {
	for _, testCase := range rectangleWidthHeightTests {
		t.Run(testCase.rect.ID, func(t *testing.T) {
			require.Equal(t, testCase.expectedWidth, testCase.rect.Width())
		})
	}
}

func TestRectangleHeight(t *testing.T) {
	for _, testCase := range rectangleWidthHeightTests {
		t.Run(testCase.rect.ID, func(t *testing.T) {
			require.Equal(t, testCase.expectedHeight, testCase.rect.Height())
		})
	}
}

func edge(id string, x, y, w, z int) *rekt.EdgeCoordinates[string] {
	return &rekt.EdgeCoordinates[string]{
		ID: id,
		X:  x,
		Y:  y,
		W:  w,
		Z:  z,
	}
}

var rectangleTouchCoordinatesTests = []struct {
	rect      rekt.Rectangle[string]
	target    rekt.Rectangle[string]
	coords    [4]*rekt.EdgeCoordinates[string]
	coordsRev [4]*rekt.EdgeCoordinates[string]
}{
	{
		rekt.NewRectangle("below", 0, 10, 20, 20),
		rekt.NewRectangle("above", 5, 0, 15, 10),
		[4]*rekt.EdgeCoordinates[string]{
			edge("above", 5, 10, 15, 10), nil, nil, nil,
		},
		[4]*rekt.EdgeCoordinates[string]{
			nil, nil, edge("below", 5, 10, 15, 10), nil,
		},
	},
	{
		rekt.NewRectangle("left", 0, 0, 10, 20),
		rekt.NewRectangle("right", 10, 5, 15, 15),
		[4]*rekt.EdgeCoordinates[string]{
			nil, edge("right", 10, 5, 10, 15), nil, nil,
		},
		[4]*rekt.EdgeCoordinates[string]{
			nil, nil, nil, edge("left", 10, 5, 10, 15),
		},
	},
}

func TestRectangleEdgeCoordinates(t *testing.T) {
	for _, testCase := range rectangleTouchCoordinatesTests {
		t.Run(fmt.Sprintf("%s -> %s", testCase.rect.ID, testCase.target.ID), func(t *testing.T) {
			require.Equal(t, testCase.coords[0], testCase.rect.TouchCoordinates(testCase.target, rekt.Top))
			require.Equal(t, testCase.coords[1], testCase.rect.TouchCoordinates(testCase.target, rekt.Right))
			require.Equal(t, testCase.coords[2], testCase.rect.TouchCoordinates(testCase.target, rekt.Bottom))
			require.Equal(t, testCase.coords[3], testCase.rect.TouchCoordinates(testCase.target, rekt.Left))
		})

		t.Run(fmt.Sprintf("%s -> %s", testCase.target.ID, testCase.rect.ID), func(t *testing.T) {
			require.Equal(t, testCase.coordsRev[0], testCase.target.TouchCoordinates(testCase.rect, rekt.Top))
			require.Equal(t, testCase.coordsRev[1], testCase.target.TouchCoordinates(testCase.rect, rekt.Right))
			require.Equal(t, testCase.coordsRev[2], testCase.target.TouchCoordinates(testCase.rect, rekt.Bottom))
			require.Equal(t, testCase.coordsRev[3], testCase.target.TouchCoordinates(testCase.rect, rekt.Left))
		})
	}
}
