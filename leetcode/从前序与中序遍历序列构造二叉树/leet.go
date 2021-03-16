package main

// TreeNode 二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	// 递归终止条件
	if len(preorder) == 0 {
		return nil
	}
	// 确定左右子树
	root := &TreeNode{Val: preorder[0]}
	// 中序遍历中找到根节点的位置
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	// 遍历左子树
	root.Left = buildTree(preorder[1:i+1], inorder[:i])
	// 遍历右子树
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return root
}
func main() {

}
