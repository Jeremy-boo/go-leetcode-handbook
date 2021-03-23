## 深度优先遍历(dfs)

> 深度优先遍历采用栈 或者 递归方式解决

通过遍历数的节点，一直到该节点的末端，然后重新检查是否还有别的节点，一直遍历下去

### 深度优先遍历模板

```
# python 模板代码-递归写法

visited = set()

def dfs(node, visited)
    // termitor
    if node in visited :
        // 已经访问过的节点
        return
    visited.add(node)
    
    // 处理当前成节点
    for next_node in node.children():
        if not  next_node in visited:
            dfs(next_node,visited)
            
# python 模板代码- 栈的写法
def DFS(self, tree): 

	if tree.root is None: 
		return [] 

	visited, stack = [], [tree.root]

	while stack: 
		node = stack.pop() 
		visited.add(node)

		process (node) 
		nodes = generate_related_nodes(node) 
		stack.push(nodes) 

	# other processing work 
```