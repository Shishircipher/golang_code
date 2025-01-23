package main

import (
	"fmt"
)

// Function to generate all permutations by swapping elements
func permute(a []int, l int, r int) [][]int {
	var result [][]int

	// Base case: when left index equals right index, add the permutation to the result
	if l == r {
		// Create a copy of 'a' to avoid modifying the original slice
		perm := make([]int, len(a))
		copy(perm, a)
		result = append(result, perm)
	} else {
		// Loop through the slice and swap elements to generate permutations
		for i := l; i <= r; i++ {
			// Swap elements
			a[l], a[i] = a[i], a[l]

			// Recursively generate permutations
			result = append(result, permute(a, l+1, r)...)

			// Backtrack to restore the original slice
			a[l], a[i] = a[i], a[l]
		}
	}
	return result
}

func main() {
	a := []int{1, 2, 3}
	// Generate all permutations
	permutations := permute(a, 0, len(a)-1)

	// Print the permutations
	for _, p := range permutations {
		fmt.Println(p)
	}
}
