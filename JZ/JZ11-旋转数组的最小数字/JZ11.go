package main

import "fmt"

// 解题思路：二分查找
func main() {
	nums := []int{3, 4, 5, 1, 2}
	fmt.Println(minArray(nums))
}

func minArray(nums []int) int {
	i, j := 0, len(nums)-1
	for i < j {
		mid := (i + j) >> 1
		if nums[mid] > nums[i] {
			i = mid + 1
		} else if nums[j] > nums[mid] {
			j = mid
		} else {
			j--
		}
	}
	return nums[i]
}
