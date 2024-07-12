package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%d\n", cutRope1(10))
}

// 贪心算法
func cutRope1(n int) int {
	if n < 2 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 3
	}
	// 拆分为足够多的2和3，当n = 4时，拆分为2*2
	timesOf3 := n / 3
	if n-timesOf3*3 == 1 {
		timesOf3--
	}
	timesOf2 := (n - timesOf3*3) / 2
	return int((int(math.Pow(3, float64(timesOf3)))) * (int(math.Pow(2, float64(timesOf2)))))
}
