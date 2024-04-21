package snippet

func lcs(w1, w2 string) int {
	seqC := 0
	dp := make([][]int, len(w1))
	for i := 0; i < len(w1); i++ {
		dp[i] = make([]int, len(w2))
	}
	for i := 0; i < len(w1); i++ {
		for j := 0; j < len(w2); j++ {
			if w1[i] == w2[j] {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i-1][j-1] + 1
				}

				if seqC < dp[i][j] {
					seqC = dp[i][j]
				}
			}
		}
	}
	return seqC
}
