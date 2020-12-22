## 题目: Pow(x, n)

> 递归-分治法

### 题目说明
````
实现 pow(x, n) ，即计算 x 的 n 次幂函数。
````

### Example:
```
示例 1:
输入: 2.00000, 10
输出: 1024.00000

示例 2:
输入: 2.10000, 3
输出: 9.26100

示例 3:

输入: 2.00000, -2
输出: 0.25000
解释: 2-2 = 1/22 = 1/4 = 0.25

```
### 题解:
```
Solution one: 暴力求解, 但是会超时

Solution two: 分治法
例子： 比如 2^10 = 2^5 * 2^5 = 2^2*2^2 * 2 * 2^2*2^2 * 2

Solution three：牛顿迭代法
```