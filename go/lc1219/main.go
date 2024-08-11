package main
// FIXME: not working correctly 
// [[0,0 ,0 ,22,0 ,24],
// [34,23,18,0 ,23,2],
// [11,39,20,12,0 ,0],
// [39,8 ,0 ,2 ,0 ,1],
// [19,32,26,20,20,30],
// [0 ,38,26,0 ,29,31]] (should be 478)

import (
	"fmt"
	"slices"
)

func getMaximumGold(grid [][]int) int {
	maxGold := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != 0 {
				maxGold = max(maxGold, maxFromStart(deepcopy(grid), i, j), 0, 0, 0)
			}
		}
	}
	return maxGold
}

func deepcopy(grid [][]int) [][]int {
	newGrid := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		newGrid[i] = slices.Clone(grid[i])
	}
	return newGrid
}

func maxFromStart(grid [][]int, x, y int) int {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
		return 0
	}
	if grid[x][y] == 0 {
		return -10000000
	}

	localVal := grid[x][y]
	grid[x][y] = 0

	v1 := maxFromStart(grid, x-1, y)
	v2 := maxFromStart(grid, x, y-1)
	v3 := maxFromStart(grid, x+1, y)
	v4 := maxFromStart(grid, x, y+1)
	outVal := max(v1+localVal, v2+localVal, v3+localVal, v4+localVal, localVal)
	switch outVal {
	case v1 + localVal:
		return v1 + localVal
	case v2 + localVal:
		return v2 + localVal
	case v3 + localVal:
		return v3 + localVal
	case v4 + localVal:
		return v4 + localVal
	default:
		return localVal
	}
}

func max(a, b, c, d, e int) int {
	if a > b && a > c && a > d && a > e {
		return a
	}
	if b > c && b > d && b > e {
		return b
	}
	if c > d && c > e {
		return c
	}
	if d > e {
		return d
	}
	return e
}

func main() {
	fmt.Println(getMaximumGold([][]int{{0, 6, 0}, {5, 8, 7}, {0, 9, 0}}))
	fmt.Println(getMaximumGold([][]int{{1, 0, 7}, {2, 0, 6}, {3, 4, 5}, {0, 3, 0}, {9, 0, 20}}))
}
