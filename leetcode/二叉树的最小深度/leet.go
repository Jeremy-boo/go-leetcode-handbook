package 二叉树的最小深度

import "math"

// TreeNode 二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minC := math.MaxInt64
	if root.Left != nil {
		minC = min(minDepth(root.Left), minC)
	}
	if root.Right != nil {
		minC = min(minDepth(root.Right), minC)
	}
	return minC + 1
}

func min(x, y int) int {
	if x >= y {
		return y
	} else {
		return x
	}
}
