package feat1

import "fmt"

func createIndirectGraphByEdges(edges [][2]int) map[int][]int {
	graph := map[int][]int{}
	for _, edge := range edges {
		if _, ok := graph[edge[0]]; !ok {
			graph[edge[0]] = []int{}
		}
		if _, ok := graph[edge[1]]; !ok {
			graph[edge[1]] = []int{}
		}
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	fmt.Println(graph)
	return graph
}

func dfs(current int, graph map[int][]int, visited map[int]bool) {
	if visited[current] {
		return
	}
	visited[current] = true
	for _, adj := range graph[current] {
		dfs(adj, graph, visited)
	}
}

func bfs(current int, graph map[int][]int, visited map[int]bool) {
	fmt.Println("eieiza")
	queue := []int{current}
	for len(queue) > 0 {
		item := queue[0]
		visited[item] = true
		queue := queue[1:]
		for _, adj := range graph[item] {
			if visited[adj] {
				continue
			}
			queue = append(queue, adj)
		}
	}
}
