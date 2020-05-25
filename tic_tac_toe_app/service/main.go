package main

import (
	"fmt"

	"github.com/ayushman-singh-gehlot/TicTacToe/tic_tac_toe_app/components"
)

func main() {
	var board1 *components.Board
	board1 = components.CreateBoard(3)
	fmt.Println(board1)

}
