package main

import (
	"fmt"

	"github.com/ayushman-singh-gehlot/TicTacToe/components"
	"github.com/ayushman-singh-gehlot/TicTacToe/service"
)

func main() {
	temp := service.NewBoardService(3)
	temp.MarkCell(0, components.XMark)
	temp.MarkCell(3, components.OMark)
	fmt.Println(temp.DisplayBoard())
}
