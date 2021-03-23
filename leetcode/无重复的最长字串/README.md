## 无重复的最长字串

> 滑动窗口

### 题目描述
~~~
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
~~~

### Example
~~~
示例 1:

输入: s = "abcabcbb"
输出: 3 
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。



示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
~~~
### 题解

```cassandraql
// 1. 使用暴力破解法
//     双层for 循环来做

// 2. 通过滑动窗口的形式来做
// abcabcbb 从a开始，right指针后移，用map遇到重复字符串
```