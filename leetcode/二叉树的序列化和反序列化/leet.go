package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TreeNode 二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	data []string
}

// Constructor 构建方法
func Constructor() Codec {
	return Codec{}
}

// serialize 序列化
func (this *Codec) serialize(root *TreeNode) string {
	return this.recurSerialize(root, "")
}

// deserialize 反序列化
func (this *Codec) deserialize(str string) *TreeNode {
	strArrs := strings.Split(str, ",")
	for i := 0; i < len(strArrs); i++ {
		if strArrs[i] != "" {
			this.data = append(this.data, strArrs[i])
		}
	}
	return this.recurDeserialize()
}

func (this *Codec) recurDeserialize() *TreeNode {
	// 反序列化，拼接data
	if this.data[0] == "null" {
		this.data = this.data[1:]
		return nil
	}
	val, _ := strconv.Atoi(this.data[0])
	root := &TreeNode{Val: val}
	this.data = this.data[1:]
	root.Left = this.recurDeserialize()
	root.Right = this.recurDeserialize()
	return root
}

func (this *Codec) recurSerialize(root *TreeNode, str string) string {
	if root == nil {
		str += "null,"
	} else {
		// 处理当前层逻辑
		str += strconv.Itoa(root.Val) + ","
		// 下探到下一层
		str = this.recurSerialize(root.Left, str)
		str = this.recurSerialize(root.Right, str)
	}
	return str
}

func main() {

	l := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	r := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   4,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
	}
	root := &TreeNode{
		Val:   1,
		Left:  l,
		Right: r,
	}
	ser := Constructor()
	data := ser.serialize(root)
	fmt.Println(data)
	deser := Constructor()
	ans := deser.deserialize(data)
	fmt.Println(ans)
}
