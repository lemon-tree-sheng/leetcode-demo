package edit_distance

func minDistance(s1, s2 string) int {
	var c1 = toCharArray(s1)
	var c2 = toCharArray(s2)

	if len(c1) == 0 {
		return len(c2)
	}
	if len(c2) == 0 {
		return len(c1)
	}

	var dp [][]int
	for i := 0; i < len(c1); i++ {
		dp[i][0] = i
	}
	for i := 0; i < len(c2); i++ {
		dp[0][i] = i
	}

	for i := 1; i < len(c1); i++ {
		for j := 1; j < len(c2); j++ {
			if c1[i] == c2[j] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}

	return dp[len(c1)-1][len(c2)-1]
}

func min(insert int, delete int, replace int) int {
	return 0
}

func toCharArray(s1 string) []byte {
	return []byte(s1)
}
