package array

func generateMatrix(n int) [][]int {
	loop := n / 2
	startx, starty := 0, 0
	count, offset := 1, 1
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	for loop > 0 {
		i, j := startx, starty
		for j = starty; j < n-offset; j++ {
			result[startx][j] = count
			count++
		}

		for i = startx; i < n-offset; i++ {
			result[i][j] = count
			count++
		}

		for ; j > starty; j-- {
			result[i][j] = count
			count++
		}

		for ; i > startx; i-- {
			result[i][j] = count
			count++
		}
		offset++
		starty++
		startx++
		loop--
	}

	if n%2 == 1 {
		result[n/2][n/2] = count
	}

	return result
}
