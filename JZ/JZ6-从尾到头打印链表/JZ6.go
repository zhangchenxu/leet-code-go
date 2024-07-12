package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	fmt.Println(reversePrint1(head))
}

// 遍历为数组，然后对数组进行反转
func reversePrint1(head *ListNode) []int {
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	for i, j := 0, len(arr); i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

// 递归ListNode，从最后一个开始放入数组
func reversePrint2(head *ListNode) (ans []int) {
	if head == nil {
		return
	}
	ans = reversePrint2(head.Next)
	ans = append(ans, head.Val)
	return ans
}
