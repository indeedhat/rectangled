package rectangled_test

import (
	"fmt"
	"testing"

	"github.com/indeedhat/rectangled"
	"github.com/stretchr/testify/require"
)

var rectangleOverlapTouchTests = []struct {
	rect1    rectangled.Rectangle[string]
	rect2    rectangled.Rectangle[string]
	overlaps bool
	touch1   rectangled.Edge
	touch2   rectangled.Edge
}{
	{
		rectangled.NewRectangle("above", 0, 0, 10, 10),
		rectangled.NewRectangle("below", 0, 20, 10, 30),
		false,
		rectangled.UnknownEdge,
		rectangled.UnknownEdge,
	},
	{
		rectangled.NewRectangle("left", 0, 0, 10, 10),
		rectangled.NewRectangle("right", 20, 0, 30, 10),
		false,
		rectangled.UnknownEdge,
		rectangled.UnknownEdge,
	},
	{
		rectangled.NewRectangle("touching-above", 0, 0, 10, 10),
		rectangled.NewRectangle("touching-below", 0, 10, 10, 20),
		false,
		rectangled.Bottom,
		rectangled.Top,
	},
	{
		rectangled.NewRectangle("touching-left", 0, 0, 10, 10),
		rectangled.NewRectangle("touching-right", 10, 0, 20, 10),
		false,
		rectangled.Right,
		rectangled.Left,
	},
	{
		rectangled.NewRectangle("full-overlap", 0, 0, 10, 10),
		rectangled.NewRectangle("full-overlap", 0, 0, 10, 10),
		true,
		rectangled.Top,
		rectangled.Top,
	},
	{
		rectangled.NewRectangle("inside-of", 5, 5, 10, 10),
		rectangled.NewRectangle("surrounding", 0, 0, 15, 15),
		true,
		rectangled.UnknownEdge,
		rectangled.UnknownEdge,
	},
	{
		rectangled.NewRectangle("top-left", 0, 0, 10, 10),
		rectangled.NewRectangle("top-left-surrounding", 0, 0, 15, 15),
		true,
		rectangled.Top,
		rectangled.Top,
	},
	{
		rectangled.NewRectangle("bottom-right", 5, 5, 15, 15),
		rectangled.NewRectangle("bottom-right-surrounding", 0, 0, 15, 15),
		true,
		rectangled.Right,
		rectangled.Right,
	},
	{
		rectangled.NewRectangle("overlap-above", 0, 0, 10, 10),
		rectangled.NewRectangle("overlap-below", 0, 9, 10, 19),
		true,
		rectangled.Right,
		rectangled.Right,
	},
	{
		rectangled.NewRectangle("overlap-left", 0, 0, 10, 10),
		rectangled.NewRectangle("overlap-right", 9, 0, 19, 10),
		true,
		rectangled.Top,
		rectangled.Top,
	},
	{
		rectangled.NewRectangle("bottom-right-corner", 0, 0, 10, 10),
		rectangled.NewRectangle("top-left corner", 9, 9, 19, 19),
		true,
		rectangled.UnknownEdge,
		rectangled.UnknownEdge,
	},
	{
		rectangled.NewRectangle("bottom-left-corner", 9, 0, 19, 10),
		rectangled.NewRectangle("top-right corner", 0, 9, 10, 19),
		true,
		rectangled.UnknownEdge,
		rectangled.UnknownEdge,
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
			require.Equal(t, testCase.touch1.String(), testCase.rect1.Touches(testCase.rect2).String())
		})
		t.Run(testCase.rect2.ID, func(t *testing.T) {
			require.Equal(t, testCase.touch2.String(), testCase.rect2.Touches(testCase.rect1).String())
		})
	}
}

var rectangleAreaTests = []struct {
	rect     rectangled.Rectangle[string]
	expected int
}{
	{rectangled.NewRectangle("zero height", 0, 0, 10, 0), 0},
	{rectangled.NewRectangle("zero width", 0, 0, 0, 10), 0},
	{rectangled.NewRectangle("zero size", 0, 0, 0, 0), 0},
	{rectangled.NewRectangle("positive area", 0, 0, 10, 10), 100},
	{rectangled.NewRectangle("negative area", 0, 0, -10, -10), 100},
}

func TestRectangleArea(t *testing.T) {
	for _, testCase := range rectangleAreaTests {
		t.Run(testCase.rect.ID, func(t *testing.T) {
			require.Equal(t, testCase.expected, testCase.rect.Area())
		})
	}
}

var rectangleValidityTests = []struct {
	rect     rectangled.Rectangle[string]
	expected error
}{
	{rectangled.NewRectangle("zero height", 0, 0, 10, 0), rectangled.ErrZoroArea},
	{rectangled.NewRectangle("zero width", 0, 0, 0, 10), rectangled.ErrZoroArea},
	{rectangled.NewRectangle("zero size", 0, 0, 0, 0), rectangled.ErrZoroArea},
	{rectangled.NewRectangle("flipped points", 0, 0, -10, -10), rectangled.ErrBadPoints},
	{rectangled.NewRectangle("valid", 0, 0, 10, 10), nil},
	{rectangled.NewRectangle("valid in negative", -10, -10, 0, 0), nil},
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
	rect1    rectangled.Rectangle[string]
	rect2    rectangled.Rectangle[string]
	expected *rectangled.Rectangle[string]
}{
	{
		rectangled.NewRectangle("overlap-1", 0, 0, 10, 10),
		rectangled.NewRectangle("overlap-2", 5, 5, 15, 15),
		&rectangled.Rectangle[string]{"overlap-2", 5, 5, 10, 10},
	},
	{
		rectangled.NewRectangle("overlap-2", 5, 5, 15, 15),
		rectangled.NewRectangle("overlap-1", 0, 0, 10, 10),
		&rectangled.Rectangle[string]{"overlap-1", 5, 5, 10, 10},
	},
	{
		rectangled.NewRectangle("no-overlap-1", 0, 0, 10, 10),
		rectangled.NewRectangle("no-overlap-2", 10, 10, 20, 20),
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
