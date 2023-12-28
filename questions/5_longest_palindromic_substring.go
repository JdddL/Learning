package questions

func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	// init dp
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	// fill dp
	for length := 2; length <= n; length++ {
		for l := 0; l < n-length+1; l++ {
			r := l + length - 1
			if length == 2 {
				if s[l] == s[r] {
					dp[l][r] = true
				}
				continue
			}
			dp[l][r] = dp[l+1][r-1] && s[l] == s[r]
		}
	}
	maxL, maxR := 0, 0
	for l := 0; l < n; l++ {
		for r := 0; r < n; r++ {
			if dp[l][r] && maxR-maxL < r-l {
				maxL, maxR = l, r
			}
		}
	}
	return s[maxL : maxR+1]
}
