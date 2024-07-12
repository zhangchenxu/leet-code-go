package main

import (
	"fmt"
	"sort"
)

func main() {
	var envelopes [][]int
	envelopes = [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}
	maxLen := maxEnvelopes(envelopes)
	fmt.Println(maxLen)
}

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	heighArr := make([]int, len(envelopes))
	for i := 0; i < len(envelopes); i++ {
		heighArr[i] = envelopes[i][1]
	}
	return lengthOfLIS(heighArr)
}

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	maxLen := 0
	for i := 0; i < len(nums); i++ {
		maxLen = max(maxLen, dp[i])
	}
	return maxLen
}
