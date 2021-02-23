package 二叉树的中序遍历

// NodeTree  Definition for a binary tree node.
type NodeTree struct {
	Val   int
	Left  *NodeTree
	Right *NodeTree
}

// inorderTraversal 树的中序遍历
func inorderTraversal(root *NodeTree) []int {
	res := make([]int, 0)
	inorder(root, &res)
	return res
}

// inorder
func inorder(root *NodeTree, res *[]int) {
	if root == nil {
		return
	}
	// 左 根 右
	inorder(root.Left, res)
	*res = append(*res, root.Val)
	inorder(root.Right, res)
}
