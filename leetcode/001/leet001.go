package _01

import (
	"math"
	"strconv"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode,l2 *ListNode) *ListNode {
	// 取出第一个数的数据
	arr1 := make([]int,0)
	for l1 != nil {
		arr1 = append(arr1,l1.Val)
		l1 = l1.Next
	}
	var x float64
	for i, val:= range arr1 {
		x = x+(float64(val)*math.Pow10(len(arr1)-i-1))
	}
	// 取出第二个链表中的数据
	arr2 := make([]int,0)
	for l2 != nil {
		arr2 = append(arr2,l2.Val)
		l2 = l2.Next
	}
	var y float64
	for i, val:= range arr2 {
		y = y+(float64(val)*math.Pow10(len(arr2)-i-1))
	}
	sum := int(x +y)
	sumStr := strconv.Itoa(sum)
	l3 := &ListNode{}
	for i :=0;i<len(sumStr);i++ {
		valStr := sumStr[i:i+1]
		val,_ := strconv.Atoi(valStr)
		if i == 0 {
			l3.Val = val
			continue
		}
		node := ListNode{
			Val: int(val) ,
			Next: nil,
		}
		if l3.Next == nil {
			l3.Next = &node
			continue
		}
		prev := l3.Next
		for prev != nil {
			prev = prev.Next
		}

	}
	return l3
}