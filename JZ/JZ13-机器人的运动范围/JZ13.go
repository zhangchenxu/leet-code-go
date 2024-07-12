package main

import "fmt"

func main() {
	fmt.Print(movingCount(2, 3, 1))
}

func movingCount(m, n, cnt int) int {
	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}
	var dfs func(i, j, k int) int
	dfs = func(i, j, k int) int {
		if i < 0 || i >= m || j < 0 || j >= n || k < 0 || vis[i][j] || (i/10+i%10+j/10+j%10) > k {
			return 0
		}
		// 将运动过的路径打上标记，防止被重复计算
		vis[i][j] = true
		return 1 + dfs(i+1, j, k) + dfs(i, j+1, k)
	}
	return dfs(0, 0, cnt)
}
