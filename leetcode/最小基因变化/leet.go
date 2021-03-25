package main

func main() {
}

func minMutation(start string, end string, bank []string) int {
	validStr := []string{"A", "C", "G", "T"}
	bankMap := make(map[string]bool, 0)
	for _, b := range bank {
		bankMap[b] = true
	}
	if _, ok := bankMap[end]; !ok {
		return -1
	}
	return bfs(0, start, end, bankMap, validStr)
}

func bfs(chanTimes int, start, end string, bankMap map[string]bool, changes []string) int {
	if start == end {
		return chanTimes
	}
	var visited = make(map[string]bool)
	queue := []string{start}
	for len(queue) != 0 {
		node := queue[0]
		visited[node] = true
		var next []string
		// 取出star
		for i := 0; i < len(node); i++ {
			for _, change := range changes {
				newNode := node[:i] + change + node[i+1:]
				// 符合条件且没有遍历过的
				if _, ok := visited[newNode]; !ok && isInBank(newNode, bankMap) {
					if newNode == end {
						return chanTimes + 1
					} else {
						visited[newNode] = true
						next = append(next, newNode)
					}
				}
			}
		}
		queue = next
	}
	return -1
}

func isInBank(str string, bankMap map[string]bool) bool {
	if _, ok := bankMap[str]; ok {
		return true
	} else {
		return false
	}
}
