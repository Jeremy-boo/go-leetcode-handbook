package 删除排序链表中的重复元素

// ListNode 单链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// deleteDuplicates 删除排序链表中的重复元素
func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
