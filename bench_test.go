package combinatoric

import (
	"fmt"
	"testing"
)

func BenchmarkCombinations(b *testing.B) {
	bmarks := []struct {
		name string
		l    []int
	}{
		{
			name: "five-elements",
			l:    []int{1, 2, 3, 4, 5},
		},
		{
			name: "seven-elements",
			l:    []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name: "nine-elements",
			l:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, bmark := range bmarks {
		for r := 3; r < 6; r++ {
			b.Run(fmt.Sprintf("%s:%d", bmark.name, r), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = Combinations(bmark.l, r)
				}
			})
		}

	}
}

func BenchmarkPermutations(b *testing.B) {
	bmarks := []struct {
		name string
		l    []int
	}{
		{
			name: "five-elements",
			l:    []int{1, 2, 3, 4, 5},
		},
		{
			name: "seven-elements",
			l:    []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name: "nine-elements",
			l:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, bmark := range bmarks {
		for r := 3; r < 6; r++ {
			b.Run(fmt.Sprintf("%s:%d", bmark.name, r), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = Permutations(bmark.l, r)
				}
			})
		}

	}
}
