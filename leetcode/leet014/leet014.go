package leet014

import "fmt"

// subsets 子集
func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	set := make([]int, 0)
	dfs(0, nums, &set, &ans)
	return ans
}

// 递归来做
func dfs(cur int, nums []int, mid *[]int, res *[][]int) {
	// 终止条件
	if cur == len(nums) {
		fmt.Println(append([]int(nil), *mid...))
		*res = append(*res, append([]int(nil), *mid...))
		return
	}
	// 处理当前逻辑 每一个节点都有两个分支，一个选一个不选

	// 下探到下一层，什么也不加,
	// 走不选这个分支
	dfs(cur+1, nums, mid, res)

	// 选这个分支
	*mid = append(*mid, nums[cur])
	dfs(cur+1, nums, mid, res)
	// 清理
	*mid = (*mid)[:len(*mid)-1]
}

// 迭代方式来做
