package main

import (
	"example/tic_tac_toe_app/components"
	"example/tic_tac_toe_app/service"
	"fmt"
)

func main() {
	fmt.Println("select board size  ")
	temp := components.CreateBoard(3)
	boardServiceObj := service.NewBoardService(temp)
	resultserviceObj := service.NewResultService(boardServiceObj)

	player1 := components.NewPlayer("ayushman", components.XMark)
	player2 := components.NewPlayer("fardin", components.OMark)
	boardServiceObj.PutMarkInPosition(player1, 4)
	fmt.Println(boardServiceObj.PrintBoard())
	fmt.Println(resultserviceObj.Result(player1))
	boardServiceObj.PutMarkInPosition(player2, 3)
	fmt.Println(boardServiceObj.PrintBoard())

}
