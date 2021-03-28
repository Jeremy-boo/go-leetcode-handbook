package main

import "fmt"

func main() {
	s := []int{6, 3, 2, 62, 4, 51}
	qSort(s)
	fmt.Println(s)
}
func qSort(s []int) {
	len := len(s)
	if len < 2 {
		return
	}
	head, trip := 0, len-1
	value := s[head]
	for head < trip { //s[head]就是我们的标尺，
		if s[head+1] > value { //标尺元素遇到大于它的，就把这个元素丢到最右边trip
			s[head+1], s[trip] = s[trip], s[head+1]
			trip--
		} else if s[head+1] < s[head] { //标尺元素遇到小于它的，就换位置，标尺右移动一位。
			s[head], s[head+1] = s[head+1], s[head]
			head++
		} else { //相等不用交换
			head++
		}
	}
	//进过上面的处理，保证了标尺左边的元素都小于等于标尺元素（s[head]），右边的元素大于等于标尺元素。
	qSort(s[:head])
	qSort(s[head+1:])
}
