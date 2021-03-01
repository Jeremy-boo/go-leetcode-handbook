package main

import "fmt"

func main() {
	arr := []int{3, 4, 2, 1, 5, 7, 9}
	fmt.Println(bubbleSort1(arr))
}

// bubbleSort 冒泡排序
func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func bubbleSort1(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		flag := true
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = false
			}
		}
		if flag {
			break
		}
	}
	return arr
}
