package main

import "fmt"

func main() {
	n := 128
	fmt.Printf("hammingWeight(%d) = %d\n", n, hammingWeight(n))
}

// 汉明重量算法，使用清除最低有效位的方案
func hammingWeight(n int) int {
	count := 0
	for n != 0 {
		n = n & (n - 1)
		count++
	}
	return count
}
