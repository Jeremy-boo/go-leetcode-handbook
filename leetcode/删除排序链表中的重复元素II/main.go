package 删除排序链表中的重复元素II

// ListNode 单链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// deleteDuplicates 删除排序链表中的重复元素II
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	node := &ListNode{Next: head}
	var tmpVal int

	// 遍历链表寻找重复元素
	cur := node
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			tmpVal = cur.Next.Val
			// 删除重复节点
			for cur.Next != nil && cur.Next.Val == tmpVal {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return node.Next
}
