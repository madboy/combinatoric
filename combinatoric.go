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

	indices := Range(0, n)
	cycles := RRange(n, n-r)
	permutations = append(permutations, getValues(indices[:r], values))
	for {
		ii := math.MinInt32
		for i := r - 1; i >= 0; i-- {
			cycles[i]--
			if cycles[i] == 0 {
				shuffleIndices(indices, i)
				cycles[i] = n - i
			} else {
				j := cycles[i]
				swap(indices, i, len(indices)-j)
				permutations = append(permutations, getValues(indices[:r], values))
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
	indices := Range(0, r)

	combinations = append(combinations, getValues(indices, values))

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
		combinations = append(combinations, getValues(indices, values))
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

func shuffleIndices(indices []int, i int) {
	for j := i + 1; j < len(indices); j++ {
		swap(indices, i, j)
		i++
	}
}

func fact(start, nbr int) int {
	v := 1
	for i := 0; i < nbr && start > 1; i++ {
		v *= start
		start--
	}
	return v
}
