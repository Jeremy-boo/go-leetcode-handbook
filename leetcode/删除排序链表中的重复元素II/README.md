## 删除排序链表中的重复元素2

> 哈希+双向链表实现

### 题目描述
~~~
给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。

~~~

### Example
~~~
示例 1:

输入: 1->2->3->3->4->4->5
输出: 1->2->5

示例 2:

输入: 1->1->1->2->3
输出: 2->3

~~~
### 题解

~~~
1. 新建一个链表 执行head，防止链表首部出现重复元素
2. 遍历链表，判断 head.Next.Val 和head.Next.Next.Val是否相等

~~~