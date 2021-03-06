package 移除链表元素

// ListNode 单链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummyNode := &ListNode{Next: head}
	cur := dummyNode
	for cur.Next != nil {
		if cur.Next.Val == val {
			// 移除节点
			for cur.Next != nil && cur.Next.Val == val {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummyNode.Next
}
