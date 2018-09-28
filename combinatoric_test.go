package combinatoric

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCombinationsInputValues(t *testing.T) {
	tests := []struct {
		input    []int
		r        int
		expected [][]int
	}{
		{
			input:    []int{1, 2, 3},
			r:        4,
			expected: [][]int{},
		},
		{
			input:    []int{},
			r:        1,
			expected: [][]int{},
		},
		{
			input:    []int{},
			r:        0,
			expected: [][]int{},
		},
		{
			input:    []int{1, 3, 4},
			r:        0,
			expected: [][]int{},
		},
	}
	for _, tc := range tests {
		got := Combinations(tc.input, tc.r)
		if !cmp.Equal(got, tc.expected) {
			t.Error(
				"input", tc.input,
				"expected", tc.expected,
				"got", got,
			)
		}
	}
}

func TestCombinations(t *testing.T) {
	tests := []struct {
		input    []int
		r        int
		expected [][]int
	}{
		{
			input: []int{1, 3, 5},
			r:     2,
			expected: [][]int{
				[]int{1, 3},
				[]int{1, 5},
				[]int{3, 5},
			},
		},
		{
			input: []int{1, 3, 5},
			r:     3,
			expected: [][]int{
				[]int{1, 3, 5},
			},
		},
		{
			input: []int{3, 6, 8, 9},
			r:     3,
			expected: [][]int{
				[]int{3, 6, 8},
				[]int{3, 6, 9},
				[]int{3, 8, 9},
				[]int{6, 8, 9},
			},
		},
		{
			input: []int{4, 2, 8, 0, 3, 3},
			r:     4,
			expected: [][]int{
				[]int{4, 2, 8, 0},
				[]int{4, 2, 8, 3},
				[]int{4, 2, 8, 3},
				[]int{4, 2, 0, 3},
				[]int{4, 2, 0, 3},
				[]int{4, 2, 3, 3},
				[]int{4, 8, 0, 3},
				[]int{4, 8, 0, 3},
				[]int{4, 8, 3, 3},
				[]int{4, 0, 3, 3},
				[]int{2, 8, 0, 3},
				[]int{2, 8, 0, 3},
				[]int{2, 8, 3, 3},
				[]int{2, 0, 3, 3},
				[]int{8, 0, 3, 3},
			},
		},
	}

	for _, test := range tests {
		got := Combinations(test.input, test.r)
		if len(got) != len(test.expected) || !allPresent(got, test.expected) {
			t.Error(
				"input", test.input,
				"expected", test.expected,
				"got", got,
			)
		}
	}
}

func TestPermutationsInputValues(t *testing.T) {
	tests := []struct {
		input    []int
		r        int
		expected [][]int
	}{
		{
			input:    []int{1, 2, 3},
			r:        4,
			expected: [][]int{},
		},
		{
			input:    []int{},
			r:        1,
			expected: [][]int{},
		},
		{
			input:    []int{},
			r:        0,
			expected: [][]int{},
		},
		{
			input:    []int{1, 3, 4},
			r:        0,
			expected: [][]int{},
		},
	}
	for _, tc := range tests {
		got := Permutations(tc.input, tc.r)
		if !cmp.Equal(got, tc.expected) {
			t.Error(
				"input", tc.input,
				"expected", tc.expected,
				"got", got,
			)
		}
	}
}
func TestPermutations(t *testing.T) {
	tests := []struct {
		input    []int
		r        int
		expected [][]int
	}{
		{
			input: []int{4, 5, 6},
			r:     2,
			expected: [][]int{
				[]int{4, 5},
				[]int{4, 6},
				[]int{5, 4},
				[]int{5, 6},
				[]int{6, 4},
				[]int{6, 5},
			},
		},
		{
			input: []int{0, 1, 2},
			r:     3,
			expected: [][]int{
				[]int{0, 1, 2},
				[]int{0, 2, 1},
				[]int{1, 0, 2},
				[]int{1, 2, 0},
				[]int{2, 0, 1},
				[]int{2, 1, 0},
			},
		},
		{
			input: []int{0, 1, 2, 3},
			r:     4,
			expected: [][]int{
				[]int{0, 1, 2, 3},
				[]int{0, 1, 3, 2},
				[]int{0, 2, 1, 3},
				[]int{0, 2, 3, 1},
				[]int{0, 3, 1, 2},
				[]int{0, 3, 2, 1},
				[]int{1, 0, 2, 3},
				[]int{1, 0, 3, 2},
				[]int{1, 2, 0, 3},
				[]int{1, 2, 3, 0},
				[]int{1, 3, 0, 2},
				[]int{1, 3, 2, 0},
				[]int{2, 0, 1, 3},
				[]int{2, 0, 3, 1},
				[]int{2, 1, 0, 3},
				[]int{2, 1, 3, 0},
				[]int{2, 3, 0, 1},
				[]int{2, 3, 1, 0},
				[]int{3, 0, 1, 2},
				[]int{3, 0, 2, 1},
				[]int{3, 1, 0, 2},
				[]int{3, 1, 2, 0},
				[]int{3, 2, 0, 1},
				[]int{3, 2, 1, 0},
			},
		},
		{
			input: []int{1, 2, 3, 4},
			r:     2,
			expected: [][]int{
				[]int{1, 2},
				[]int{1, 3},
				[]int{1, 4},
				[]int{2, 1},
				[]int{2, 3},
				[]int{2, 4},
				[]int{3, 1},
				[]int{3, 2},
				[]int{3, 4},
				[]int{4, 1},
				[]int{4, 2},
				[]int{4, 3},
			},
		},
	}

	for _, test := range tests {
		got := Permutations(test.input, test.r)
		if len(got) != len(test.expected) || !allPresent(got, test.expected) {
			t.Error(
				"input", test.input,
				"expected", test.expected,
				"got", got,
			)
		}
	}
}

func allPresent(got, expected [][]int) bool {
	var matches int
	for _, e := range expected {
		for _, g := range got {
			if cmp.Equal(e, g) {
				matches++
				break
			}
		}
	}
	return matches == len(got)
}
