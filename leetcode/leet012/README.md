## 题目: 验证二叉搜索树

> 递归

### 题目说明
````
给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

- 节点的左子树只包含小于当前节点的数。
- 节点的右子树只包含大于当前节点的数。
- 所有左子树和右子树自身必须也是二叉搜索树。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/validate-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
````

### Example:
```
输入:
    2
   / \
  1   3
输出: true

输入:
    5
   / \
  1   4
     / \
    3   6

输出: false
解释: 输入为: [5,1,4,null,null,3,6]。
     根节点的值为 5 ，但是其右子节点值为 4 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/validate-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

```
### 题解:
```
判断二叉树是否为二叉搜索树，根据二叉树的判断条件为我们可以相出来两种解决办法：
1. 使用递归方法，我们可以分别判断他的左右子树的节点是否满足上述条件，从而判断
整个树是否为二叉搜索树

2. 如果我们直接采用中序遍历，可以得到一个递增的数组，通过判断得到后的结果是否为递增、
来判断是否为二叉搜索树
```