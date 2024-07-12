package main

import "fmt"

func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	repeatNum := findRepeatNumber(nums)
	fmt.Println(repeatNum)
}

// 依次遍历数组，将value放入到数组的第i位，如果第i位有相同的值则说明value = i是重复的数字
// 遍历到第i位试，一定要保证i = i
func findRepeatNumber(nums []int) int {
	for i, num := range nums {
		for {
			if num == i {
				break
			}
			temp := nums[num]
			if num == temp {
				return num
			}
			nums[num] = num
			nums[i] = temp
			num = nums[i]
		}
		temp := nums[num]
		nums[num] = num
		nums[i] = temp
	}
	return -1
}
