package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	differenceMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		difference := target - nums[i]
		if _, ok := differenceMap[difference]; ok {
			return []int{differenceMap[difference], i}
		}
		differenceMap[nums[i]] = i
	}
	return nil
}
