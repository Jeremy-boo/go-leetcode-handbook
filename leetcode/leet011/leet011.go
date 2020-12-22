package leet011

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// invertTree 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	// 递归终止条件
	if root == nil {
		return nil
	}
	// 处理当前逻辑，右子树变成左子树
	root.Left, root.Right = root.Right, root.Left
	// 下探到下一层
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
