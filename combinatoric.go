package combinatoric

import (
	"math"
)

// Permutations returns permutations of values of length r
// Implementation based on https://docs.python.org/3.6/library/itertools.html#itertools.permutations
func Permutations(values []int, r int) [][]int {
	var permutations [][]int
	pool := make([]int, len(values))
	copy(pool, values)
	n := len(pool)

	if r > n {
		return permutations
	}

	indices := Range(0, n)
	cycles := RRange(n, n-r)
	permutations = append(permutations, getValues(indices[:r], pool))
	for {
		ii := math.MinInt32
		for _, i := range Reversed(Range(0, r)) {
			cycles[i]--
			if cycles[i] == 0 {
				indices = shuffleIndices(indices, i)
				cycles[i] = n - i
			} else {
				j := cycles[i]
				swap(indices, i, len(indices)-j)
				permutations = append(permutations, getValues(indices[:r], pool))
				ii = j
				break
			}
		}
		if ii == math.MinInt32 {
			return permutations
		}
	}
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

	combinations = append(combinations, getValues(indices, pool))

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
		combinations = append(combinations, getValues(indices, pool))
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

// RRange return a list of integers from high .. low-1
func RRange(high, low int) []int {
	var r []int
	for i := high; i > low; i-- {
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

func getValues(indices, values []int) []int {
	var combination []int
	for _, i := range indices {
		combination = append(combination, values[i])
	}
	return combination
}

func swap(values []int, f, t int) {
	values[f], values[t] = values[t], values[f]
}

func shuffleIndices(indices []int, i int) []int {
	var tmp []int
	tmp = append(tmp, indices[:i]...)
	tmp = append(tmp, indices[i+1:]...)
	return append(tmp, indices[i:i+1]...)
}
