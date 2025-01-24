package main

import (
	"container/list"
	"fmt"
)

func findMinSteps(mat [][]rune) int {
	n := len(mat)
	m := len(mat[0])

	// Transforming the grid
	grid := make([][]int, n)
	for i := range grid {
		grid[i] = make([]int, m)
	}

	houseID := 0
	startX, startY := 0, 0

	// Assign unique IDs to houses, and find starting position
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mat[i][j] == '.' {
				grid[i][j] = 0
			} else if mat[i][j] == '*' {
				grid[i][j] = houseID
				houseID++
			} else if mat[i][j] == 'o' {
				startX, startY = i, j
			} else {
				grid[i][j] = -1
			}
		}
	}

	// If no houses to visit, return 0
	if houseID == 0 {
		return 0
	}

	// BFS setup: visited array and queue
	vis := make([][][]bool, 1<<houseID)
	for i := range vis {
		vis[i] = make([][]bool, n)
		for j := range vis[i] {
			vis[i][j] = make([]bool, m)
		}
	}

	queue := list.New()
	vis[0][startX][startY] = true
	queue.PushBack([]int{0, startX, startY})

	// Directions for BFS traversal
	dx := []int{0, 0, 1, -1}
	dy := []int{1, -1, 0, 0}

	steps := 0

	// BFS loop
	for queue.Len() > 0 {
		qSize := queue.Len()

		for i := 0; i < qSize; i++ {
			curr := queue.Remove(queue.Front()).([]int)
			mask := curr[0]
			x, y := curr[1], curr[2]

			// If all houses are visited, return the steps
			if mask == (1<<houseID)-1 {
				return steps
			}

			// Explore all possible moves
			for k := 0; k < 4; k++ {
				newX, newY := x+dx[k], y+dy[k]

				// Check boundaries and obstacles
				if newX >= 0 && newX < n && newY >= 0 && newY < m && grid[newX][newY] != -1 {
					if mat[newX][newY] == '*' {
						// If it's a house, update the visited mask
						newMask := mask | (1 << grid[newX][newY])
						if vis[newMask][newX][newY] {
							continue
						}
						vis[newMask][newX][newY] = true
						queue.PushBack([]int{newMask, newX, newY})
					} else {
						// If it's an empty cell, continue without updating the mask
						if vis[mask][newX][newY] {
							continue
						}
						vis[mask][newX][newY] = true
						queue.PushBack([]int{mask, newX, newY})
					}
				}
			}
		}

		steps++
	}

	// If unable to visit all houses, return -1
	return -1
}

func main() {
	mat := [][]rune{
		{'.', '.', '.', '.', '.', '.', '.'},
		{'.', 'o', '.', '.', '.', '*', '.'},
		{'.', '.', '.', '.', '.', '.', '.'},
		{'.', '*', '.', '.', '.', '*', '.'},
		{'.', '.', '.', '.', '.', '.', '.'},
	}

	res := findMinSteps(mat)
	fmt.Println(res)
}

