package leet010

func GenerateParenthesis(n int) []string {
	arrS := make([]string, 0)
	_generate(0, 0, n, "", &arrS)
	return arrS
}

// _generate 递归求解
func _generate(left, right int, max int, s string, arr *[]string) {
	if left == max && right == max {
		*arr = append(*arr, s)
	}
	if left <= max {
		_generate(left+1, right, max, s+"(", arr)
	}
	if right < left {
		_generate(left, right+1, max, s+")", arr)
	}
}
