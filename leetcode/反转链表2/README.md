## 题目: 反转链表2

> 链表

### 题目说明
~~~
反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

说明:
1 ≤ m ≤ n ≤ 链表长度。

链接:https://leetcode-cn.com/problems/reverse-linked-list-ii/
~~~

### Example:
~~~
输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL
~~~

### 解法
~~~
关键点：m的前一个节点；m和n区间的反转；n之后的节点
1->2->3->4->5->6,假设m=3,n=5
prem = 2, m = 3, n = 5, n + 1= 6

1.定义一个next指向head的虚拟指针（为了最后返回head），并定义prem
2.使用迭代法反转m到n的链表，pre为反转后n的节点，current为n+1节点
链表变为：1->2->3<-4<-5 6
3.1 原m节点为prem.Next需要指向n+1节点current（2->3原链表还没断）
prem.Next.Next = current
3.2 反转后的m节点为pre，prem需要重新指向反转后的m节点(2->3链表已经断，更新为了2->5)
prem.Next = pre
最终链表变为：1->2->5->4->3->6

~~~