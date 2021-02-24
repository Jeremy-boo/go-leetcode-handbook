package tree

// NodeTree  Definition for a binary tree node.
type NodeTree struct {
	Val   interface{}
	Left  *NodeTree
	Right *NodeTree
}

// inOrderTraversal 树的中序遍历 左根右
func inOrderTraversal(root *NodeTree) []int {
	arr := make([]int, 0)
	inOrder(root, &arr)
	return arr
}

// inOrder 树的中序遍历 左->根->右
func inOrder(root *NodeTree, res *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, res)
	*res = append(*res, root.Val.(int))
	inOrder(root.Right, res)
}

// preOrder 树的先序遍历 根->左->右
func preOrder(root *NodeTree, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val.(int))
	preOrder(root.Left, res)
	preOrder(root.Right, res)
}

// postOrder 树的后续遍历 左->右->根
func postOrder(root *NodeTree, res *[]int) {
	if root == nil {
		return
	}
	postOrder(root.Left, res)
	postOrder(root.Right, res)
	*res = append(*res, root.Val.(int))
}
