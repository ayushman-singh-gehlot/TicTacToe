package main

import (
	"example/components"
	"example/service"
	"fmt"
)

func main() {
	temp := service.NewBoardService(3)
	temp.MarkCell(0, components.XMark)
	temp.MarkCell(3, components.OMark)
	fmt.Println(temp.DisplayBoard())
}
