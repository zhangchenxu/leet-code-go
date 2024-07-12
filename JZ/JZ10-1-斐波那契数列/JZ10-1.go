package main

import "fmt"

func main() {
	n := 10
	fmt.Printf("Fibonacci(%d) = %d\n", n, Fibonacci(n))
}

func Fibonacci(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, (a+b)%1000000007
	}
	return a
}
