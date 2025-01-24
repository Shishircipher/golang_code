package main

import (
	"fmt"
	"math"
)

// totalCost calculates the minimum cost to visit all cities
func totalCost(mask int, curr int, n int, cost [][]int, memo [][]int) int {
	// Base case: all cities visited
	if mask == (1<<n)-1 {
		return cost[curr][0]
	}

	// If already computed, return the stored value
	if memo[curr][mask] != -1 {
		return memo[curr][mask]
	}

	// Initialize the minimum cost to a very high value
	ans := math.MaxInt32

	// Try visiting all unvisited cities
	for i := 0; i < n; i++ {
		if (mask & (1 << i)) == 0 { // If city `i` is not visited
			// Update mask to include city `i` and recursively calculate the cost
			ans = min(ans, cost[curr][i]+totalCost(mask|(1<<i), i, n, cost, memo))
		}
	}

	// Save the result in the memoization table
	memo[curr][mask] = ans
	return ans
}

// tsp initializes memoization and calls totalCost
func tsp(cost [][]int) int {
	n := len(cost)

	// Memoization table, initialized to -1
	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, 1<<n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	// Start from city 0, with only city 0 visited initially
	return totalCost(1, 0, n, cost, memo)
}

// Helper function to find the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Cost matrix
	cost := [][]int{
		{0, 10, 15, 20},
		{10, 0, 35, 25},
		{15, 35, 0, 30},
		{20, 25, 30, 0},
	}

	// Calculate the result
	res := tsp(cost)
	fmt.Println("The minimum cost to visit all cities is:", res)
}

