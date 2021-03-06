package 旋转链表

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// rotateRight 旋转链表
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	// step one: 找链表尾节点和 计算链表的长度
	length, tail := 1, head
	for tail.Next != nil {
		tail = tail.Next
		length++
	}
	// 形成环形链表
	tail.Next = head

	k = k % length
	for i := 0; i < length-k; i++ {
		tail = tail.Next
	}
	head, tail.Next = tail.Next, nil
	return head
}
