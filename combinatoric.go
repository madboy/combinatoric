package combinatoric

import (
	"math"
)

// Permutations returns permutations of values of length r
// Implementation based on https://docs.python.org/3.6/library/itertools.html#itertools.permutations
func Permutations(values []int, r int) [][]int {
	n := len(values)
	if r > n || n == 0 || r == 0 {
		return [][]int{}
	}

	nbrOfPermutations := fact(n, r)
	permutations := make([][]int, 0, nbrOfPermutations)
	pool := make([]int, n)
	copy(pool, values)

	indices := Range(0, n)
	cycles := RRange(n, n-r)
	permutations = append(permutations, getValues(indices[:r], pool))
	for {
		ii := math.MinInt32
		for i := r - 1; i >= 0; i-- {
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
	n := len(values)
	if r > n || n == 0 || r == 0 {
		return [][]int{}
	}
	nbrOfCombinations := fact(n, n) / (fact(r, r) * fact(n-r, n-r))
	combinations := make([][]int, 0, nbrOfCombinations)
	pool := make([]int, n)
	copy(pool, values)

	indices := Range(0, r)

	combinations = append(combinations, getValues(indices, pool))

	for {
		ii := math.MinInt32
		for i := r - 1; i >= 0; i-- {
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
		for j := ii + 1; j < r; j++ {
			indices[j] = indices[j-1] + 1
		}
		combinations = append(combinations, getValues(indices, pool))
	}
}

// Range return a list of integers from low .. high-1
func Range(low, high int) []int {
	r := make([]int, 0, high-low)
	for i := low; i < high; i++ {
		r = append(r, i)
	}
	return r
}

// RRange return a list of integers from high .. low-1
func RRange(high, low int) []int {
	r := make([]int, 0, high-low)
	for i := high; i > low; i-- {
		r = append(r, i)
	}
	return r
}

func getValues(indices, values []int) []int {
	combination := make([]int, 0, len(indices))
	for _, i := range indices {
		combination = append(combination, values[i])
	}
	return combination
}

func swap(values []int, f, t int) {
	values[f], values[t] = values[t], values[f]
}

func shuffleIndices(indices []int, i int) []int {
	tmp := make([]int, 0, len(indices))
	tmp = append(tmp, indices[:i]...)
	tmp = append(tmp, indices[i+1:]...)
	return append(tmp, indices[i:i+1]...)
}

func fact(start, nbr int) int {
	v := 1
	for i := 0; i < nbr; i++ {
		v *= start
		start--
	}
	return v
}
