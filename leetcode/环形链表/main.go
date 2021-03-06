package 环形链表

// ListNode 单链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// hasCycle 判断链表是否有环
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
