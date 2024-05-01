package model

type Board struct {
	Grid [][]string
}

func NewBoard() Board {
	board := Board{Grid: make([][]string, 0, 0)}
	for i := 0; i < 3; i++ {
		row := []string{" ", " ", " "}
		board.Grid = append(board.Grid, row)
	}
	return board
}
