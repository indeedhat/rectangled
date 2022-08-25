package rectangled_test

import (
	"fmt"
	"testing"

	"github.com/indeedhat/rectangled"
	"github.com/stretchr/testify/require"
)

var setAddRectangleTests = []struct {
	rect         rectangled.Rectangle[string]
	expected     error
	expectedArea int
}{
	{rectangled.NewRectangle("zero height", 0, 0, 10, 0), rectangled.ErrZoroArea, 0},
	{rectangled.NewRectangle("zero width", 0, 0, 0, 10), rectangled.ErrZoroArea, 0},
	{rectangled.NewRectangle("zero size", 0, 0, 0, 0), rectangled.ErrZoroArea, 0},
	{rectangled.NewRectangle("flipped points", 0, 0, -10, -10), rectangled.ErrBadPoints, 0},
	{rectangled.NewRectangle("negative position", -10, -10, 0, 0), rectangled.ErrNegativePositionInSet, 0},
	{rectangled.NewRectangle("valid", 0, 0, 10, 10), nil, 100},
	{rectangled.NewRectangle("valid offset", 10, 10, 100, 100), nil, 12100},
}

func TestNewSet(t *testing.T) {
	for _, testCase := range setAddRectangleTests {
		t.Run(testCase.rect.ID, func(t *testing.T) {
			set, err := rectangled.NewSet("test set", 0, 0, []rectangled.Rectangle[string]{testCase.rect})
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
	set, _ := rectangled.NewSet("test set", 0, 0, nil)

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
	var set1, _ = rectangled.NewSet("set-1", 0, 0, []rectangled.Rectangle[string]{
		rectangled.NewRectangle("set-1_rect-1", 0, 0, 10, 10),
		rectangled.NewRectangle("set-1_rect-2", 10, 0, 100, 100),
	})
	var set2, _ = rectangled.NewSet("set-2", 10, 10, []rectangled.Rectangle[string]{
		rectangled.NewRectangle("set-2_rect-1", 0, 10, 10, 20),
		rectangled.NewRectangle("set-2_rect-2", 10, 10, 110, 110),
	})

	testCases := []struct {
		receiver *rectangled.Set[string]
		target   *rectangled.Set[string]
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
	var set1, _ = rectangled.NewSet("set-1", 0, 0, []rectangled.Rectangle[string]{
		rectangled.NewRectangle("set-1_rect-1", 0, 0, 100, 100),
		rectangled.NewRectangle("set-1_rect-2", 100, 0, 110, 10),
	})
	var set2, _ = rectangled.NewSet("set-2", 110, 0, []rectangled.Rectangle[string]{
		rectangled.NewRectangle("set-2_rect-1", 0, 0, 10, 10),
		rectangled.NewRectangle("set-2_rect-2", 0, 10, 10, 20),
	})

	testCases := []struct {
		receiver *rectangled.Set[string]
		target   *rectangled.Set[string]
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
	var multipleChildren = []rectangled.Rectangle[string]{
		rectangled.NewRectangle("rect-1", 0, 0, 10, 10),
		rectangled.NewRectangle("rect-1", 10, 10, 15, 20),
	}
	var multiple, _ = rectangled.NewSet("multiple", 10, 10, multipleChildren)

	var singleChild = []rectangled.Rectangle[string]{
		rectangled.NewRectangle("rect-1", 0, 0, 10, 10),
	}
	var single, _ = rectangled.NewSet("single", 10, 10, singleChild)

	var noneChild []rectangled.Rectangle[string] = nil
	var none, _ = rectangled.NewSet("none", 10, 10, noneChild)

	var testCases = []struct {
		set      *rectangled.Set[string]
		children []rectangled.Rectangle[string]
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
	var multipleChildren = []rectangled.Rectangle[string]{
		rectangled.NewRectangle("rect-1", 0, 0, 10, 10),
		rectangled.NewRectangle("rect-1", 10, 10, 15, 20),
	}
	var multiple, _ = rectangled.NewSet("multiple", 10, 10, multipleChildren)

	var singleChild = []rectangled.Rectangle[string]{
		rectangled.NewRectangle("rect-1", 0, 0, 10, 10),
	}
	var single, _ = rectangled.NewSet("single", 10, 10, singleChild)

	var noneChild []rectangled.Rectangle[string] = nil
	var none, _ = rectangled.NewSet("none", 10, 10, noneChild)

	var testCases = []struct {
		set      *rectangled.Set[string]
		children []rectangled.Rectangle[string]
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
	var set, _ = rectangled.NewSet("bottom-right-most-child", 0, 0, []rectangled.Rectangle[string]{
		rectangled.NewRectangle("top-middle", 5, 0, 15, 10),
		rectangled.NewRectangle("bottom-middle", 5, 10, 15, 20),
		rectangled.NewRectangle("right-middle", 10, 5, 20, 15),
		rectangled.NewRectangle("left-middle", 0, 5, 10, 15),

		rectangled.NewRectangle("top-left", 0, 0, 10, 10),
		rectangled.NewRectangle("top-right", 10, 0, 20, 10),
		rectangled.NewRectangle("bottom-left", 0, 10, 10, 20),
		rectangled.NewRectangle("bottom-right", 10, 10, 20, 20),

		rectangled.NewRectangle("center", 5, 5, 15, 15),
	})

	require.Equal(t, "top-right", set.ChildOnEdge(rectangled.Top, rectangled.Right).ID)
	require.Equal(t, "top-left", set.ChildOnEdge(rectangled.Top, rectangled.Left).ID)

	require.Equal(t, "bottom-right", set.ChildOnEdge(rectangled.Bottom, rectangled.Right).ID)
	require.Equal(t, "bottom-left", set.ChildOnEdge(rectangled.Bottom, rectangled.Left).ID)

	require.Equal(t, "top-right", set.ChildOnEdge(rectangled.Right, rectangled.Top).ID)
	require.Equal(t, "bottom-right", set.ChildOnEdge(rectangled.Right, rectangled.Bottom).ID)

	require.Equal(t, "top-left", set.ChildOnEdge(rectangled.Left, rectangled.Top).ID)
	require.Equal(t, "bottom-left", set.ChildOnEdge(rectangled.Left, rectangled.Bottom).ID)
}
