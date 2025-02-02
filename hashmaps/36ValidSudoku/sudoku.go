package validsudoku

func IsValidSudoku(board [][]byte) bool {
	// create rows, cols, boxes maps
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	boxes := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			cell := board[i][j]
			if cell == '.' {
				continue
			}
			boxIdx := (i/3)*3 + j/3
			if rows[i][cell] || cols[j][cell] || boxes[boxIdx][cell] {
				return false
			}
			rows[i][cell] = true
			cols[j][cell] = true
			boxes[boxIdx][cell] = true
		}
	}
	return true
}