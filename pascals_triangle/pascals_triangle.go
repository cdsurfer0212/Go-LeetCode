package pascalstriangle

import "math"

func generate(numRows int) [][]int {
	if numRows <= 0 {
		return nil
	}
	result := [][]int{{1}}
	for i := 1; i < numRows; i++ {
		tmp := make([]int, i+1)
		tmp[0] = 1
		tmp[i] = 1
		for j := 1; j < int(math.Ceil(float64(i+1)/float64(2))); j++ {
			tmp[j] = result[i-1][j-1] + result[i-1][j]
			tmp[i-j] = tmp[j]
		}
		result = append(result, tmp)
	}
	return result
}
