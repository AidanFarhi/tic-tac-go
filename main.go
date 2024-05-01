package main

import (
	"tic-tac-go/display"
	"tic-tac-go/model"
)

func main() {

	display.PrintIntroMessage()

	game_over := false
	board := model.NewBoard()

	for {
		if game_over {
			break
		}
		display.PrintEnterMoveMessage()
		display.DisplayBoard(board)
		game_over = true
	}
}
