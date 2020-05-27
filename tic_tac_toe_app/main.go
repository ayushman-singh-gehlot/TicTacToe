package main

import (
	"example/tic_tac_toe_app/components"
	"example/tic_tac_toe_app/service"
	"fmt"
)

func main() {
	temp := components.CreateBoard(3)
	boardServiceObj := service.NewBoardService(temp)
	player1 := components.NewPlayer("ayushman", components.XMark)
	player2 := components.NewPlayer("fardin", components.OMark)
	boardServiceObj.PutMarkInPosition(player1, 4)
	fmt.Println(boardServiceObj.PrintBoard())
	boardServiceObj.PutMarkInPosition(player2, 3)
	fmt.Println(boardServiceObj.PrintBoard())

}
