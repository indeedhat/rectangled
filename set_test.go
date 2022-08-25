package rekt_test

import (
	"fmt"
	"testing"

	"github.com/indeedhat/rekt"
	"github.com/stretchr/testify/require"
)

var setAddRectangleTests = []struct {
	rect         rekt.Rectangle[string]
	expected     error
	expectedArea int
}{
	{rekt.NewRectangle("zero height", 0, 0, 10, 0), rekt.ErrZoroArea, 0},
	{rekt.NewRectangle("zero width", 0, 0, 0, 10), rekt.ErrZoroArea, 0},
	{rekt.NewRectangle("zero size", 0, 0, 0, 0), rekt.ErrZoroArea, 0},
	{rekt.NewRectangle("flipped points", 0, 0, -10, -10), rekt.ErrBadPoints, 0},
	{rekt.NewRectangle("negative position", -10, -10, 0, 0), rekt.ErrNegativePositionInSet, 0},
	{rekt.NewRectangle("valid", 0, 0, 10, 10), nil, 100},
	{rekt.NewRectangle("valid offset", 10, 10, 100, 100), nil, 12100},
}

func TestNewSet(t *testing.T) {
	for _, testCase := range setAddRectangleTests {
		t.Run(testCase.rect.ID, func(t *testing.T) {
			set, err := rekt.NewSet("test set", 0, 0, []rekt.Rectangle[string]{testCase.rect})
			if testCase.expected != nil {
				require.ErrorIs(t, testCase.expected, err)
				return
			}

			t.Log(set)
			require.Nil(t, err)
			require.Equal(t, testCase.expectedArea, set.Area())
		})
	}
}

func TestSetAddRectangle(t *testing.T) {
	set, _ := rekt.NewSet("test set", 0, 0, nil)

	for _, testCase := range setAddRectangleTests {
		t.Run(testCase.rect.ID, func(t *testing.T) {
			if testCase.expected != nil {
				require.ErrorIs(t, testCase.expected, set.AddRectangle(testCase.rect))
				return
			}

			require.Nil(t, set.AddRectangle(testCase.rect))
			require.Equal(t, testCase.expectedArea, set.Area())
		})
	}
}

func TestSetOverlapsChildren(t *testing.T) {
	var set1, _ = rekt.NewSet("set-1", 0, 0, []rekt.Rectangle[string]{
		rekt.NewRectangle("set-1_rect-1", 0, 0, 10, 10),
		rekt.NewRectangle("set-1_rect-2", 10, 0, 100, 100),
	})
	var set2, _ = rekt.NewSet("set-2", 10, 10, []rekt.Rectangle[string]{
		rekt.NewRectangle("set-2_rect-1", 0, 10, 10, 20),
		rekt.NewRectangle("set-2_rect-2", 10, 10, 110, 110),
	})

	testCases := []struct {
		receiver *rekt.Set[string]
		target   *rekt.Set[string]
		expected []string
	}{
		{set1, set2, []string{"set-2_rect-1", "set-2_rect-2"}},
		{set2, set1, []string{"set-1_rect-2"}},
	}

	for _, testCase := range testCases {
		var testName = fmt.Sprintf("%s -> %s", testCase.receiver.ID, testCase.target.ID)

		t.Run(testName, func(t *testing.T) {
			overlaps := testCase.receiver.OverlapsChildren(*testCase.target)
			require.Len(t, overlaps, len(testCase.expected))
			for i, rect := range overlaps {
				require.Equal(t, testCase.expected[i], rect.ID)
			}
		})
	}
}

func TestSetTouchesChildren(t *testing.T) {
	var set1, _ = rekt.NewSet("set-1", 0, 0, []rekt.Rectangle[string]{
		rekt.NewRectangle("set-1_rect-1", 0, 0, 100, 100),
		rekt.NewRectangle("set-1_rect-2", 100, 0, 110, 10),
	})
	var set2, _ = rekt.NewSet("set-2", 110, 0, []rekt.Rectangle[string]{
		rekt.NewRectangle("set-2_rect-1", 0, 0, 10, 10),
		rekt.NewRectangle("set-2_rect-2", 0, 10, 10, 20),
	})

	testCases := []struct {
		receiver *rekt.Set[string]
		target   *rekt.Set[string]
		expected []string
	}{
		{set1, set2, []string{"set-2_rect-1", "set-2_rect-2"}},
		{set2, set1, []string{"set-1_rect-2"}},
	}

	for _, testCase := range testCases {
		var testName = fmt.Sprintf("%s -> %s", testCase.receiver.ID, testCase.target.ID)

		t.Run(testName, func(t *testing.T) {
			overlaps := testCase.receiver.TouchesChildren(*testCase.target)
			require.Len(t, overlaps, len(testCase.expected))
			for i, rect := range overlaps {
				require.Equal(t, testCase.expected[i], rect.ID)
			}
		})
	}
}

func TestSetChildren(t *testing.T) {
	var multipleChildren = []rekt.Rectangle[string]{
		rekt.NewRectangle("rect-1", 0, 0, 10, 10),
		rekt.NewRectangle("rect-1", 10, 10, 15, 20),
	}
	var multiple, _ = rekt.NewSet("multiple", 10, 10, multipleChildren)

	var singleChild = []rekt.Rectangle[string]{
		rekt.NewRectangle("rect-1", 0, 0, 10, 10),
	}
	var single, _ = rekt.NewSet("single", 10, 10, singleChild)

	var noneChild []rekt.Rectangle[string] = nil
	var none, _ = rekt.NewSet("none", 10, 10, noneChild)

	var testCases = []struct {
		set      *rekt.Set[string]
		children []rekt.Rectangle[string]
	}{
		{multiple, multipleChildren},
		{single, singleChild},
		{none, noneChild},
	}

	for _, testCase := range testCases {
		t.Run(testCase.set.ID, func(t *testing.T) {
			require.Equal(t, testCase.children, testCase.set.Children())
		})
	}
}

func TestSetOffsetChildren(t *testing.T) {
	var multipleChildren = []rekt.Rectangle[string]{
		rekt.NewRectangle("rect-1", 0, 0, 10, 10),
		rekt.NewRectangle("rect-1", 10, 10, 15, 20),
	}
	var multiple, _ = rekt.NewSet("multiple", 10, 10, multipleChildren)

	var singleChild = []rekt.Rectangle[string]{
		rekt.NewRectangle("rect-1", 0, 0, 10, 10),
	}
	var single, _ = rekt.NewSet("single", 10, 10, singleChild)

	var noneChild []rekt.Rectangle[string] = nil
	var none, _ = rekt.NewSet("none", 10, 10, noneChild)

	var testCases = []struct {
		set      *rekt.Set[string]
		children []rekt.Rectangle[string]
	}{
		{multiple, multipleChildren},
		{single, singleChild},
		{none, noneChild},
	}

	for _, testCase := range testCases {
		t.Run(testCase.set.ID, func(t *testing.T) {
			var offsetChildren = testCase.set.OffsetChildren()

			require.Len(t, testCase.set.OffsetChildren(), len(testCase.children))
			if testCase.children == nil {
				return
			}

			for i, child := range testCase.children {
				require.Equal(t, child.ID, offsetChildren[i].ID)
				require.Equal(t, child.X+testCase.set.X, offsetChildren[i].X)
				require.Equal(t, child.Y+testCase.set.Y, offsetChildren[i].Y)
				require.Equal(t, child.W+testCase.set.X, offsetChildren[i].W)
				require.Equal(t, child.Z+testCase.set.Y, offsetChildren[i].Z)
			}
		})
	}
}

func TestSetChildOnEdge(t *testing.T) {
	var set, _ = rekt.NewSet("bottom-right-most-child", 0, 0, []rekt.Rectangle[string]{
		rekt.NewRectangle("top-middle", 5, 0, 15, 10),
		rekt.NewRectangle("bottom-middle", 5, 10, 15, 20),
		rekt.NewRectangle("right-middle", 10, 5, 20, 15),
		rekt.NewRectangle("left-middle", 0, 5, 10, 15),

		rekt.NewRectangle("top-left", 0, 0, 10, 10),
		rekt.NewRectangle("top-right", 10, 0, 20, 10),
		rekt.NewRectangle("bottom-left", 0, 10, 10, 20),
		rekt.NewRectangle("bottom-right", 10, 10, 20, 20),

		rekt.NewRectangle("center", 5, 5, 15, 15),
	})

	require.Equal(t, "top-right", set.ChildOnEdge(rekt.Top, rekt.Right).ID)
	require.Equal(t, "top-left", set.ChildOnEdge(rekt.Top, rekt.Left).ID)

	require.Equal(t, "bottom-right", set.ChildOnEdge(rekt.Bottom, rekt.Right).ID)
	require.Equal(t, "bottom-left", set.ChildOnEdge(rekt.Bottom, rekt.Left).ID)

	require.Equal(t, "top-right", set.ChildOnEdge(rekt.Right, rekt.Top).ID)
	require.Equal(t, "bottom-right", set.ChildOnEdge(rekt.Right, rekt.Bottom).ID)

	require.Equal(t, "top-left", set.ChildOnEdge(rekt.Left, rekt.Top).ID)
	require.Equal(t, "bottom-left", set.ChildOnEdge(rekt.Left, rekt.Bottom).ID)
}
