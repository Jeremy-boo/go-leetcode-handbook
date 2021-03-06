package 链表的中间节点

// ListNode 单链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// middleNode 链表的中间节点
func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	// 哈希表存储链表的每个索引和节点
	cache := make(map[int]*ListNode)
	count := 0
	for head != nil {
		count++
		cache[count] = head
		head = head.Next
	}
	return cache[count/2+1]
}
