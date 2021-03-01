package main

func quickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
}

func partition(list []int, left, right int) int {
	pivot := list[left]
	for left < right {
		// 右指针往前走
		for right >= left && list[right] > pivot {
			right--
		}
	}
	return left
}
