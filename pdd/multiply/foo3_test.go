package multiply

func foo(n, k int) [][]int {
	res := make([][]int, 0)
	path := make([]int, k)
	dfs(n, k, 1, path, res)
	return res
}

func dfs(n, k, start int, path []int, res [][]int) {
	if len(path) == k {
		res = append(res, path)
		return
	}

	for i := start; i <= n; i++ {
		path = append(path, i)
		dfs(n, k, i+1, path, res)
		path = path[:len(path)-1]
	}
}
