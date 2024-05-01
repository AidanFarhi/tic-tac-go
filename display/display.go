package display

import (
	"fmt"
	"tic-tac-go/model"
)

func PrintIntroMessage() {
	fmt.Println("Welcome to Tic-Tac-Go!")
}

func PrintEnterMoveMessage() {
	fmt.Println("Enter move:")
}

func DisplayBoard(board model.Board) {
	fmt.Println(board)
}
