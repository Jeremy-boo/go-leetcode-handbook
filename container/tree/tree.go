package tree

// NodeTree  Definition for a binary tree node.
type NodeTree struct {
	val   interface{}
	left  *NodeTree
	right *NodeTree
}

// inOrderTraversal 树的中序遍历 左根右
func inOrderTraversal(root *NodeTree) []int {
	arr := make([]int, 0)
	inOrder(root, &arr)
	return arr
}

func inOrder(root *NodeTree, res *[]int) {
	if root == nil {
		return
	}
	inOrder(root.left, res)
	*res = append(*res, root.val.(int))
	inOrder(root.right, res)
}
