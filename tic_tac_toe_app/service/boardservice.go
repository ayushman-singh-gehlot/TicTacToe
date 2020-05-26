package service

import (
	"fmt"

	"example/tic_tac_toe_app/components"
)

type BoardService struct {
	*components.Board
}

func NewBoardService(size uint8) *BoardService {
	return &BoardService{components.CreateBoard(size)}
}

func (b *BoardService) PutMarkInPosition(index uint8, mark string) error {
	err := b.BoardCells[index].SetMark(mark)
	return err
}

func (b *BoardService) PrintBoard() string {
	tempStr := ""
	for i := uint8(0); i < (b.Size * b.Size); i++ {
		if i%b.Board.Size == 0 {
			tempStr = tempStr + fmt.Sprint("\n\t")
		}
		tempStr = tempStr + fmt.Sprint(b.BoardCells[i].GetMark(), " ")
	}
	return tempStr
}

func (b *BoardService) CheckBoardIsFull() bool {
	for _, cell := range b.BoardCells {
		if cell.GetMark() == components.NoMark {
			return false
		}
	}
	return true
}
