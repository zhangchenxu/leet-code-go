package main

import "fmt"

func main() {
	original := "We are happy."
	result := appendMethod(original)
	fmt.Println(result)
}

// 拼接法
func appendMethod(original string) string {
	var result []rune
	for _, char := range original {
		if char == ' ' {
			result = append(result, '%', '2', '0')
		} else {
			result = append(result, char)
		}
	}
	return string(result)
}

// 双指针
