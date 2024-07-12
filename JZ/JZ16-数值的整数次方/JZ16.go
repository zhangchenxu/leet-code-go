package main

import "fmt"

func main() {
	fmt.Println(myPow(2.00000, 10))
}

// 使用快速幂算法，循环处理，先判断是否是奇偶数，
// 如果是奇数，那么需要将乘一次当前数，如果是偶数，就对当前数进行平方运算
// 时间复杂度：O(logn)
func myPow(x float64, n int) float64 {
	qpow := func(a float64, n int) float64 {
		ans := 1.0
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				ans *= a
			}
			a *= a
		}
		return ans
	}
	if n >= 0 {
		return qpow(x, n)
	}
	return 1 / qpow(x, -n)
}
