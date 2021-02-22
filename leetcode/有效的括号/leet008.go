package 有效的括号

import "errors"

// IsValid 有效的括号
func IsValid(s string) bool {
	// 奇数肯定不是
	if len(s)%2 != 0 {
		return false
	}
	var slice []interface{}
	for _, val := range s {
		if string(val) == "(" || string(val) == "{" || string(val) == "[" {
			slice = push(string(val), slice)
		} else {
			switch string(val) {
			case ")":
				x, err := pop(&slice)
				if err != nil {
					return false
				}
				if x.(string) != "(" {
					return false
				}
			case "]":
				x, err := pop(&slice)
				if err != nil {
					return false
				}
				if x.(string) != "[" {
					return false
				}
			case "}":
				x, err := pop(&slice)
				if err != nil {
					return false
				}
				if x.(string) != "{" {
					return false
				}
			}
		}

	}
	return len(slice) == 0
}

// pop 出栈
func pop(arr *[]interface{}) (val interface{}, err error) {
	if len(*arr) <= 0 {
		return nil, errors.New("Stack index out of length")
	}
	val = (*arr)[0]
	*arr = (*arr)[1:]
	return val, nil
}

// push 进栈
func push(val interface{}, arr []interface{}) []interface{} {
	newArr := make([]interface{}, len(arr)+1, len(arr)+1)
	newArr[0] = val
	copy(newArr[1:], arr)
	return newArr
}
