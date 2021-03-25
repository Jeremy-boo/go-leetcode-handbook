package main

import "fmt"

// TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, res *[]int) {
	// dfs 的先序遍历
	if root != nil {
		*res = append(*res, root.Val)
		dfs(root.Left, res)
		dfs(root.Right, res)
	}
}

func main() {
	l := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 4,
		},
		Right: &TreeNode{
			Val: 5,
		},
	}
	r := &TreeNode{
		Val: 7,
		Left: &TreeNode{
			Val: 8,
		},
		Right: &TreeNode{
			Val: 9,
		},
	}
	root := &TreeNode{
		Val: 3,
	}
	root.Left = l
	root.Right = r
	res := make([]int, 0)
	dfs(root, &res)
	fmt.Println(res)
}
