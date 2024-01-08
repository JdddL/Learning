package questions

import "testing"

func Test2397(t *testing.T) {
	matrix := [][]int{
		{1, 0},
		{0, 1},
		{1, 1},
	}
	result := maximumRows(matrix, 1)
	t.Log(result)
}
