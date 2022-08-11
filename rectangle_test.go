package rectangled_test

import (
	"testing"

	"github.com/indedhat/rectangled"
	"github.com/stretchr/testify/require"
)

var rectangleOverlapTests = []struct {
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
	for _, testCase := range rectangleOverlapTests {
		t.Run(testCase.rect1.ID, func(t *testing.T) {
			require.Equal(t, testCase.overlaps, testCase.rect1.Overlaps(testCase.rect2))
		})
		t.Run(testCase.rect2.ID, func(t *testing.T) {
			require.Equal(t, testCase.overlaps, testCase.rect2.Overlaps(testCase.rect1))
		})
	}
}

func TestRectangleTouches(t *testing.T) {
	for _, testCase := range rectangleOverlapTests {
		t.Run(testCase.rect1.ID, func(t *testing.T) {
			require.Equal(t, testCase.touch1.String(), testCase.rect1.Touches(testCase.rect2).String())
		})
		t.Run(testCase.rect2.ID, func(t *testing.T) {
			require.Equal(t, testCase.touch2.String(), testCase.rect2.Touches(testCase.rect1).String())
		})
	}
}
