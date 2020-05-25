package service

import (
	"fmt"

	"github.com/ayushman-singh-gehlot/TicTacToe/tic_tac_toe_app/components"
)

type BoardService struct {
	*components.Board
}

func NewBoardService(size uint8) *BoardService {
	return &BoardService{components.CreateBoard(size)}
}

func (b *BoardService) MarkCell(index uint8, mark string) error {
	err := b.Board.BoardCells[index].SetMark(mark)
	return err
}

func (b *BoardService) DisplayBoard() string {
	tempStr := ""
	for i := uint8(0); i < (b.Board.Size * b.Board.Size); i++ {
		if i%b.Board.Size == 0 {
			tempStr = tempStr + fmt.Sprint("\n\t")
		}
		tempStr = tempStr + fmt.Sprint(b.Board.BoardCells[i].GetMark(), " ")
	}
	return tempStr
}
