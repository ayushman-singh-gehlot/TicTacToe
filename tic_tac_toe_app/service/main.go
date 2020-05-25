package main

import (
	"fmt"

	"github.com/ayushman-singh-gehlot/TicTacToe/tic_tac_toe_app/components"
	"github.com/ayushman-singh-gehlot/TicTacToe/tic_tac_toe_app/service"
)

func main() {
	board1 := service.NewBoardService(3)
	board1.MarkCell(0, components.XMark)
	fmt.Println(board1.DisplayBoard())

}
