package 两数相加

import (
	"errors"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

// AddTwoNumbers 两数相加
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 取出链表中的元素 [2,4,3] -> 3,4,2 到栈中
	var s1 []interface{}
	var s2 []interface{}
	for l1 != nil {
		s1 = push(s1, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		s2 = push(s2, l2.Val)
		l2 = l2.Next
	}
	carry := 0
	// 两数相加运算
	getTwoSum := func(x, y int) int {

		sum := x + y + carry
		if sum > 9 {
			sum = sum - 10
			carry = 1
		} else {
			carry = 0
		}
		return sum
	}

	var sumL *ListNode
	insertL := func(val int) {
		if sumL == nil {
			sumL = &ListNode{
				Val:  val,
				Next: nil,
			}
			return
		}
		node := &ListNode{
			Val:  val,
			Next: sumL,
		}
		sumL = node
	}
	for {
		v1, err1 := pop(&s1)
		v2, err2 := pop(&s2)
		if err1 != nil && err2 != nil {
			// s1,s2都为空
			break
		} else if err1 == nil && err2 == nil {
			// s1,s2都不为空
			insertL(getTwoSum(v1.(int), v2.(int)))
		} else if err1 == nil {
			// s1 为空
			insertL(getTwoSum(v1.(int), 0))
		} else {
			// s2 为空
			insertL(getTwoSum(0, v2.(int)))
		}
	}
	if carry != 0 {
		insertL(1)
	}
	return reverseList(sumL)
}

// reverseList 反转链表
func reverseList(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode = nil
	for cur != nil {
		prev, cur, cur.Next = cur, cur.Next, prev
	}
	return prev
}

// pop 出栈,从栈尾出
func pop(s *[]interface{}) (interface{}, error) {
	if len(*s) <= 0 {
		return nil, errors.New("栈内无元素")
	}
	value := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return value, nil
}

// push 入栈
func push(s []interface{}, val interface{}) []interface{} {
	newS := make([]interface{}, len(s)+1, len(s)+1)
	newS[0] = val
	copy(newS[1:], s)
	return newS
}
