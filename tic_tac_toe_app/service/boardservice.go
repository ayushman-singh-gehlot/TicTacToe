package service

import (
	"fmt"

	"example/tic_tac_toe_app/components"
)

type BoardService struct {
	*components.Board
}

func NewBoardService(board *components.Board) *BoardService {
	return &BoardService{board}
}

func (b *BoardService) PutMarkInPosition(player *components.Player, position uint8) error {
	err := b.BoardCells[position].SetMark(player.Mark)
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
