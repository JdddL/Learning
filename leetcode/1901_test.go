package questions

import (
	"fmt"
	"testing"
)

func Test1901(t *testing.T) {
	arr := [][]int{
		{10, 50, 40, 30, 20},
		{1, 500, 2, 3, 4},
	}
	res := findPeakGrid(arr)
	fmt.Println(res)
}
