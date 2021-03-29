package main

import "fmt"

func mergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	mid := len(arr) / 2
	l := mergeSort(arr[:mid])
	r := mergeSort(arr[mid:])
	return merge(l, r)
}

func merge(leftArr, rightArr []int) []int {
	res := make([]int, 0)
	lIndex, rIndex := 0, 0
	for lIndex < len(leftArr) && rIndex < len(rightArr) {
		if leftArr[lIndex] > rightArr[rIndex] {
			res = append(res, rightArr[rIndex])
			rIndex++
		} else {
			res = append(res, leftArr[lIndex])
			lIndex++
		}
	}
	if lIndex < len(leftArr) {
		res = append(res, leftArr[lIndex:]...)
	}
	if rIndex < len(rightArr) {
		res = append(res, rightArr[rIndex:]...)
	}
	return res
}

func main() {
	s := []int{6, 7, 8, 10, 4, 6, 99}
	res := mergeSort(s)
	fmt.Println(res)
}
