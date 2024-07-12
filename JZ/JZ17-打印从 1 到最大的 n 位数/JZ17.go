package main

import (
	"fmt"
)

func main() {
	// 测试字符串长度
	// s := "1234"
	// fmt.Println(s[3])
	// fmt.Println(len(s))
	// fmt.Println(s[:len(s)-1])

	// 测试append方法
	// fmt.Printf(string(append([]byte{'1'}, byte('1' + 2))))

	fmt.Println(print(2))
}

// 时间复杂度O(2n)，空间复杂度O(n)
func print(n int) []string {
	ans := []string{}

	s := []byte{}

	// 定义深度优先搜索函数，i为当前位数，j为目标位数，比如从个位数(i = 0)开始打印，到十位数(i = 1)
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i == j {
			ans = append(ans, string(s))
			return
		}
		k := 0
		if i == 0 {
			k++
		}
		// 每位数从1循环到9
		for k < 10 {
			// 使用ASICII码进行表示，整数的ASIC码是连续的，所以使用byte('0' + k)
			// byte(k)是获取对应的ASIC字符串，不能直接使用
			s = append(s, byte('0'+k))
			dfs(i+1, j)
			s = s[:len(s)-1]
			k++
		}
	}

	for i := 1; i <= n; i++ {
		dfs(0, i)
	}

	return ans
}
