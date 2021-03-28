package main

import (
	"errors"
	"fmt"
)

type heap []int

// Sink 让堆中的元素通过递归的与自己的左右子节点比较下沉到合适的位置。
func (h *heap) Sink(index int) {
	// 堆的大小
	n := len(*h)
	// 左节点不超过堆的大小
	for 2*index+1 < n {
		left := 2*index + 1

		if left < n-1 && (*h)[left] > (*h)[left+1] {
			// 如果左节点比右节点大，则设置成右节点
			left++
		}
		if (*h)[index] <= (*h)[left] {
			// 如果当前节点已经比左右子节点中最小的都小，则已经满足
			break
		}
		// 交换当前节点和子结点中较小的节点
		(*h)[index], (*h)[left] = (*h)[left], (*h)[index]
		// 下探到下一层
		fmt.Println(h)
		index = left
	}
}

// Swim 让堆中的元素递归的与自己的根节点比较从而上浮到合适的位置
func (h *heap) Swim(index int) {
	for (index-1)/2 >= 0 && (*h)[index] < (*h)[(index-1)/2] {
		// 存在父节点且当前节点比父节点小，则交换节点
		(*h)[(index-1)/2], (*h)[index] = (*h)[index], (*h)[(index-1)/2]
		// 向上递归
		index = (index - 1) / 2
	}
}

// Pop 堆中取出元素
func (h *heap) Pop() (int, error) {
	n := len(*h)
	if n < 0 {
		return -1, errors.New("heap is empty")
	}
	val := (*h)[0]
	(*h)[0], (*h)[n-1] = (*h)[n-1], (*h)[0]
	*h = (*h)[:n-1]
	return val, nil
}

// Push 向堆中添加元素
func (h *heap) Push(val int) {
	// 添加元素到末尾
	*h = append(*h, val)
	h.Swim(len(*h) - 1)
}

// Len 计算堆中元素的个数
func (h *heap) Len() int {
	return len(*h)
}

//VerifyHeap
func VerifyHeap(arr []int) *heap {
	h := heap(arr)

	n := len(h)
	if n < 2 {
		return &h
	}
	for i := (n - 2) / 2; i >= 0; i-- {
		h.Sink(i)
	}
	return &h
}

func main() {
	c := []int{9, 1, 6, 2, 5, 3, 7, 4}
	nc := VerifyHeap(c)
	fmt.Println(*nc) // 结果：[1 3 2 5 9 7 4]
}
