package rotateimage

func Rotate(matrix [][]int)  {

	// to rotate the matrix 90 degrees clockwise
	// we first transpose the matrix
	// then reverse each row

	// transpose the matrix
	for i := 0; i < len(matrix); i++ {
		for j := i+1; j < len(matrix[0]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// now we reverse each row
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0])/2; j++ {
			matrix[i][j], matrix[i][len(matrix[0])-1-j] = matrix[i][len(matrix[0])-1-j], matrix[i][j]
		}
	}
}