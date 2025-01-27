package spiralmatrix

func SpiralOrder(matrix [][]int) []int {

	res := make([]int, len(matrix)*len(matrix[0]))
	t, b, l, r := 0, len(matrix)-1, 0, len(matrix[0])-1
	idx := 0

	// for spiral matrix we shrink in to centre 
	// we traverse from right to left, top to bottom, left to right, bottom to top
	// then we bring in the boundaries

	for t <= b && l <= r {
		for i := l; i <=r; i++ {
			// row is constant, column changes
			res[idx] = matrix[t][i]
			idx++
		}
		// bring top down
		t++
		// move from top to bottom
		for i := t; i <= b; i++ {
			// column is constant, row changes
			res[idx] = matrix[i][r]
			idx++
		}
		// bring right in
		r--

		if t <= b{
			// move from right to left
			for i := r; i >= l; i-- {
				// row is constant, column changes
				res[idx] = matrix[b][i]
				idx++
			}
			// bring bottom up
			b--
		}

		if l <= r {
			//move from bottom to top
			for i := b; i >= t; i-- {
				// column is constant, row changes
				res[idx] = matrix[i][l]
				idx++
			}
			// bring left in
			l++
		}

	}

	return res
}