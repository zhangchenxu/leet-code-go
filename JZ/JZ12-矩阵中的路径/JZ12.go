package main

import "fmt"

func main() {
	board := [][]byte{{'B', 'A', 'A', 'D', 'E'}, {'C', 'E', 'E', 'P', 'G'}, {'I', 'N', 'Y', 'F', 'G'}}
	word := "AAEC"
	fmt.Println(existPath(board, word))
}

func existPath(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true
		}
		board[i][j] = '#'
		result := dfs(i+1, j, k+1) || dfs(i-1, j, k+1) || dfs(i, j+1, k+1) || dfs(i, j-1, k+1)
		// 如果匹配失败，需要还原路径上的字符，网上的答案都有问题
		if !result {
			board[i][j] = word[k]
		}
		return result
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
