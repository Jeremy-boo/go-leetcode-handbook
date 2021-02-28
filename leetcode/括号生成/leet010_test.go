package 括号生成

import "testing"

func TestGenerateParenthesis(t *testing.T) {
	s := GenerateParenthesis(1)
	t.Log(s)
}

func TestRecur(t *testing.T) {
	Recur(0, 2*3, "")
}
