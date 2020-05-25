package service

import (
	"example/components"
	"fmt"
)

func (b *components.Board) MarkCell(index uint8, mark string) error {
	err := b.boardCells[index].SetMark(mark)
	return err
}

func (b *components.Board) DisplayBoard() string {
	tempStr := ""
	for i := uint8(0); i < (b.size * b.size); i++ {
		if i%b.size == 0 {
			tempStr = tempStr + fmt.Sprint("\n\t")
		}
		tempStr = tempStr + fmt.Sprint(b.boardCells[i].GetMark(), " ")
	}
	return tempStr
}
