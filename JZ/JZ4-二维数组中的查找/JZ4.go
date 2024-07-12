package main

import "fmt"

func main() {
	matrix := [][]int{{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30}}
	result := findNumber(matrix, 20)
	fmt.Print(result)
}

// findNumber searches for a target integer in a 2D matrix.
//
// Parameters:
// - matrix: a 2D slice of integers representing the matrix.
// - target: an integer representing the target value to search for.
//
// Returns:
// - bool: true if the target is found in the matrix, false otherwise.
func findNumber(matrix [][]int, target int) bool {
	rowMax := len(matrix[0])
	colMax := len(matrix)

	row := 0
	col := colMax - 1
	for {
		if row >= rowMax || col < 0 {
			break
		}
		if matrix[row][col] == target {
			return true
		}
		if matrix[row][col] < target {
			row++
			continue
		}
		if matrix[row][col] > target {
			col--
			continue
		}
	}
	return false
}
