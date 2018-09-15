package combinatoric

import (
	"math"
)

//Permutations return successive len(values) length permutations of elements in values.
//Using Heap's algorithm, https://en.wikipedia.org/wiki/Heap%27s_algorithm
//TODO: this should be updated to support r length permutations
func Permutations(values []int) [][]int {
	var c []int
	var permutations [][]int
	n := len(values)

	for i := 0; i < n; i++ {
		c = append(c, 0)
	}

	tmp := make([]int, len(values))
	copy(tmp, values)
	permutations = append(permutations, tmp)

	for i := 0; i < n; {
		if c[i] < i {
			if i%2 == 0 {
				swap(values, 0, i)
			} else {
				swap(values, c[i], i)
			}
			tmp := make([]int, len(values))
			copy(tmp, values)
			permutations = append(permutations, tmp)
			c[i]++
			i = 0
		} else {
			c[i] = 0
			i++
		}
	}
	return permutations
}

// Combinations retruns unique combinations of values of length r
// Implementation from python itertools https://docs.python.org/3.6/library/itertools.html#itertools.combinations
func Combinations(values []int, r int) [][]int {
	var combinations [][]int
	pool := make([]int, len(values))
	copy(pool, values)
	n := len(pool)
	if r > n {
		return combinations
	}
	indices := Range(0, r)

	combinations = append(combinations, createCombination(indices, pool))

	for {
		ii := math.MinInt32
		for _, i := range Reversed(Range(0, r)) {
			if indices[i] != i+n-r {
				ii = i
				break
			}
		}
		// Python version uses a for else structure which we don't have
		// this is to catch the case where we are no longer finding
		// a new indices base
		if ii == math.MinInt32 {
			return combinations
		}
		indices[ii]++
		for _, j := range Range(ii+1, r) {
			indices[j] = indices[j-1] + 1
		}
		combinations = append(combinations, createCombination(indices, pool))
	}
}

// Range return a list of integers from low .. high-1
func Range(low, high int) []int {
	var r []int
	for i := low; i < high; i++ {
		r = append(r, i)
	}
	return r
}

// Reversed returns a list of integers in reverse order
func Reversed(values []int) []int {
	var reversed []int
	for i := len(values) - 1; i >= 0; i-- {
		reversed = append(reversed, values[i])
	}
	return reversed
}

func createCombination(indices, values []int) []int {
	var combination []int
	for _, i := range indices {
		combination = append(combination, values[i])
	}
	return combination
}

func swap(values []int, f, t int) {
	values[f], values[t] = values[t], values[f]
}
