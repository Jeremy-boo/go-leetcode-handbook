package leet016

// numToStr 数字到字母的映射关系
var numToStr = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

// letterCombinations 电话号码的组合
func letterCombinations(digits string) []string {
	res := make([]string, 0)
	search(0, "", digits, &res)
	return res
}

// search 递归函数
func search(level int, s, digits string, res *[]string) {
	// 递归终止条件
	if level == len(digits) {
		*res = append(*res, s)
		return
	}
	// process
	letter := numToStr[string(digits[level])]
	for j := 0; j < len(letter); j++ {
		// 下探到下一层
		search(level+1, s+string(letter[j]), digits, res)
	}
	// reverse
}
