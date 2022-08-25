package rekt

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

var maxTests = []struct {
	a        int
	b        int
	expected int
}{
	{0, 1, 1},
	{1, 0, 1},
	{0, 0, 0},
	{1, 1, 1},
	{-1, 1, 1},
	{1, -1, 1},
	{-2, -1, -1},
	{-1, -2, -1},
}

func TestMax(t *testing.T) {
	for _, testCase := range maxTests {
		t.Run(fmt.Sprintf("a(%d) b(%d)", testCase.a, testCase.b), func(t *testing.T) {
			require.Equal(t, testCase.expected, max(testCase.a, testCase.b))
		})
	}
}

var minTests = []struct {
	a        int
	b        int
	expected int
}{
	{0, 1, 0},
	{1, 0, 0},
	{0, 0, 0},
	{1, 1, 1},
	{-1, 1, -1},
	{1, -1, -1},
	{-2, -1, -2},
	{-1, -2, -2},
}

func TestMin(t *testing.T) {
	for _, testCase := range minTests {
		t.Run(fmt.Sprintf("a(%d) b(%d)", testCase.a, testCase.b), func(t *testing.T) {
			require.Equal(t, testCase.expected, min(testCase.a, testCase.b))
		})
	}
}

var absTests = []struct {
	n        int
	expected int
}{
	{1, 1},
	{0, 0},
	{-1, 1},
}

func TestAbs(t *testing.T) {
	for _, testCase := range absTests {
		t.Run(strconv.Itoa(testCase.n), func(t *testing.T) {
			require.Equal(t, testCase.expected, abs(testCase.n))
		})
	}
}
