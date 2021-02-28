package 括号生成

import "fmt"

// GenerateParenthesis 括号生成
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

func Recur(level int, max int, str string) {
	// 递归终结条件
	if level >= max {
		fmt.Println(str)
		return
	}
	// 处理当前层逻辑
	// 下探到下一层
	Recur(level+1, max, str+"(")
	Recur(level+1, max, str+")")
	// 清楚
}

// generateParenthesis
func generateParenthesis(n int) []string {
	// 返回结果
	var res []string
	// 递归函数
	var recur func(left, right, max int, str string)
	recur = func(left, right, max int, str string) {
		// 终止条件
		if left == max && right == max {
			res = append(res, str)
			return
		}
		// 判断括号是否合法
		if left <= max {
			recur(left+1, right, max, str+"(")
		}
		if right < left && right <= max {
			recur(left, right+1, max, str+")")
		}
	}
	// 执行函数
	recur(0, 0, n, "")
	return res
}
