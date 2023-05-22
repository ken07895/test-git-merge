package feat1

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
	return graph
}
