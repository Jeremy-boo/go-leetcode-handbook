package leet012

import "math"

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isValidRecu(root, math.MinInt64, math.MaxInt64)
}

func isValidRecu(root *TreeNode, low int, upper int) bool {
	// 边界条件
	if root == nil {
		return true
	}
	// 处理当前层的逻辑
	if root.Val <= low || root.Val >= upper {
		return false
	}
	// 下探到下一层
	return isValidRecu(root.Left, low, root.Val) && isValidRecu(root.Right, root.Val, upper)
}
