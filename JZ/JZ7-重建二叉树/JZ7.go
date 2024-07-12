package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	fmt.Println(buildTree(preorder, inorder))
}

// 前序遍历：根节点，左子树，右子树
// 中序遍历：左子树，根节点，右子树
// 后序遍历：左子树，右子树，根节点
func buildTree(preorder []int, inorder []int) *TreeNode {
	indexMap := make(map[int]int)
	for i, v := range inorder {
		indexMap[v] = i
	}

	var dfs func(i, j, n int) *TreeNode
	dfs = func(i, j, n int) *TreeNode {
		if n == 0 {
			return nil
		}
		root := &TreeNode{Val: preorder[i]}
		l := indexMap[preorder[i]] - j
		// 0 + 1  = 1
		// 1 - 1 = 0
		// 1 = 0 = 1
		root.Left = dfs(i+1, j, l)
		// 0 + 1 + 1 - 0 = 2
		// 1 + 1 = 2
		// 5 - 1 - 1 = 3
		root.Right = dfs(i+1+l, indexMap[preorder[i]]+1, n-l-1)
		return root
	}
	return dfs(0, 0, len(preorder))
}
