package n_queue

func solveQueues(n int) [][]string {
	var res [][]string
	var board [][]byte
	board = constructBoard(n)
	backtrack(board, 0, res)
	return res
}

func backtrack(board [][]byte, row int, res [][]string) {
	if row == len(board) {
		board2 := convertBoard(board)
		res = append(res, board2)
		return
	}

	for i := 0; i < len(board); i++ {
		board[row][i] = 'Q'
		if isValid(board) {
			backtrack(board, row+1, res)
		}
		board[row][i] = '.'
	}
}

func isValid(board [][]byte) bool {
	// 每一行不能有两个 Q
	for i := 0; i < len(board); i++ {
		flag := false
		for j := 0; j < len(board); j++ {
			if board[i][j] == 'Q' {
				if flag {
					return false
				} else {
					flag = true
				}
			}
		}
	}

	// 每一列不能有两个 Q
	// 略

	// 每一个对角线不能有两个 Q
	// 从 0 行开启的斜线，从 0 列开启的斜线
	// 这里比较难写一些
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			flag := false
			if board[i][j] == 'Q' {
				if flag {
					return false
				} else {
					flag = true
				}
			}
		}
	}
	return false
}

func convertBoard(board [][]byte) []string {
	var res []string
	for i := 0; i < len(board); i++ {
		res = append(res, string(board[i]))
	}
	return res
}

func constructBoard(n int) [][]byte {
	var res [][]byte
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res[i][j] = '.'
		}
	}
	return res
}
