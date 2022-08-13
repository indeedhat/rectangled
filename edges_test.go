package rectangled_test

import (
	"testing"

	"github.com/indeedhat/rectangled"
	"github.com/stretchr/testify/require"
)

var edgesStringTests = []struct {
	edge     rectangled.Edge
	expected string
}{
	{rectangled.Top, "Top"},
	{rectangled.Right, "Right"},
	{rectangled.Bottom, "Bottom"},
	{rectangled.Left, "Left"},
	{rectangled.Edge(100), "Unknown"},
}

func TestEdgeString(t *testing.T) {
	for _, testCase := range edgesStringTests {
		t.Run(testCase.expected, func(t *testing.T) {
			require.Equal(t, testCase.expected, testCase.edge.String())
		})
	}
}
