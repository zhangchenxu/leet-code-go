package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Print(removeNode(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}, 3))
}

func removeNode(head *ListNode, i int) *ListNode {
	// 定义哑节点，方便应对head节点为nil的情况
	dummy := &ListNode{0, head}
	// 定义指针，对cur的操作和对head操作是一样的，因为指向的都是同一个指针
	cur := dummy
	for cur.Next == nil {
		if cur.Next.Val == i {
			cur.Next = cur.Next.Next
			break
		}
		cur = cur.Next
	}
	return dummy.Next
}
