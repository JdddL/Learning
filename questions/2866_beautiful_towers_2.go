package questions

func maximumSumOfHeights(maxHeights []int) int64 {
	if len(maxHeights) == 0 {
		return 0
	}

	prefix, suffix := make([]int, len(maxHeights)), make([]int, len(maxHeights))
	prefixStack, suffixStack := []int{}, []int{}

	for i := 0; i < len(maxHeights); i++ {
		peak := maxHeights[i]
		for len(prefixStack) > 0 && peak < maxHeights[prefixStack[len(prefixStack)-1]] {
			prefixStack = prefixStack[:len(prefixStack)-1]
		}
		if len(prefixStack) == 0 {
			prefix[i] = (i + 1) * peak
		} else {
			last := prefixStack[len(prefixStack)-1]
			prefix[i] = prefix[last] + (i-last)*peak
		}
		prefixStack = append(prefixStack, i)
	}

	result := 0
	for j := len(maxHeights) - 1; j >= 0; j-- {
		peak := maxHeights[j]
		for len(suffixStack) > 0 && peak < maxHeights[suffixStack[len(suffixStack)-1]] {
			suffixStack = suffixStack[:len(suffixStack)-1]
		}
		if len(suffixStack) == 0 {
			suffix[j] = (len(maxHeights) - j) * peak
		} else {
			last := suffixStack[len(suffixStack)-1]
			suffix[j] = suffix[last] + (last-j)*peak
		}
		suffixStack = append(suffixStack, j)
		result = max(result, prefix[j]+suffix[j]-peak)
	}

	return int64(result)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
