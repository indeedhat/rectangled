package rekt_test

import (
	"testing"

	"github.com/indeedhat/rekt"
	"github.com/stretchr/testify/require"
)

var edgesStringTests = []struct {
	edge     rekt.Edge
	expected string
}{
	{rekt.Top, "Top"},
	{rekt.Right, "Right"},
	{rekt.Bottom, "Bottom"},
	{rekt.Left, "Left"},
	{rekt.Edge(100), "Unknown"},
}

func TestEdgeString(t *testing.T) {
	for _, testCase := range edgesStringTests {
		t.Run(testCase.expected, func(t *testing.T) {
			require.Equal(t, testCase.expected, testCase.edge.String())
		})
	}
}
