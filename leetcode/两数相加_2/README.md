## 题目 两数相加 2

> tag： 链表

### 题目:
~~~
给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。

它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
~~~

Example:
~~~
输入：(7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 8 -> 0 -> 7
~~~

### 题解:
~~~
## 解法1 最开始解题思路
遍历取出两个链表的值，放入字符串中，最后转换成int进行相加,
但是需要考虑到int类型值范围溢出问题，导致case一直跑不通

## 解法2 
初始化两个栈，分别存入链表的值，然后从栈中取出元素相，最好进位判断就好
~~~