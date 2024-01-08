package questions

func maximumRows(matrix [][]int, numSelect int) int {
	if numSelect == len(matrix[0]) {
		return len(matrix)
	}
	selected := make([]bool, len(matrix[0]))
	return backTrackSelect(matrix, &selected, 0, numSelect)
}

func backTrackSelect(matrix [][]int, selected *[]bool, nextCol, numSelect int) int {
	if len(*selected)-nextCol <= numSelect {
		uncoveredRows := make([]bool, len(matrix))
		result := len(matrix)
		for col := 0; col < nextCol; col++ {
			if !(*selected)[col] {
				for row := 0; row < len(matrix); row++ {
					if uncoveredRows[row] {
						continue
					}
					if matrix[row][col] == 1 {
						uncoveredRows[row] = true
						result -= 1
						break
					}
				}
			}
		}
		return result
	}
	if numSelect <= 0 {
		uncoveredRows := make([]bool, len(matrix))
		result := len(matrix)
		for col, flag := range *selected {
			if !flag {
				for row := 0; row < len(matrix); row++ {
					if uncoveredRows[row] {
						continue
					}
					if matrix[row][col] == 1 {
						uncoveredRows[row] = true
						result -= 1
					}
				}
			}
		}
		return result
	}
	max := 0
	for col := nextCol; col < len(*selected); col++ {
		if len(*selected)-col < numSelect {
			break
		}
		(*selected)[col] = true
		result := backTrackSelect(matrix, selected, col+1, numSelect-1)
		if max < result {
			max = result
		}
		(*selected)[col] = false
	}
	return max
}
