package questions

func lengthOfLongestSubstring(s string) int {
	l, r := 0, 0
	n := len(s)
	hash := make(map[byte]int)
	maxLen := 0

	for l < n && l <= r {
		for r < n {
			if _, exist := hash[s[r]]; !exist {
				hash[s[r]] = r
				r++
				continue
			} else {
				break
			}
		}

		tmpLen := r - l
		if maxLen < tmpLen {
			maxLen = tmpLen
		}

		if r >= n {
			break
		}

		dupIndex := hash[s[r]]
		for i := l; i <= dupIndex; i++ {
			delete(hash, s[i])
		}
		l = dupIndex + 1
	}

	return maxLen
}
