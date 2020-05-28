package service

import "example/tic_tac_toe_app/components"

type ResultService struct {
	*BoardService
}

func NewResultService(boardService *BoardService) *ResultService {
	return &ResultService{boardService}
}

func checkRows(b *components.Board, mark string) bool {
	ret := true
	for i := uint8(0); i < (b.Size * b.Size); i = i + b.Size {
		ret = true
		for j := i; j < (i + b.Size); j++ {
			if b.BoardCells[j].GetMark() != mark {
				ret = false
			}
		}
		if ret {
			return ret
		}
	}
	return ret
}

func checkColumns(b *components.Board, mark string) bool {
	ret := true
	for i := uint8(0); i < b.Size; i++ {
		ret = true
		for j := i; j < (b.Size * b.Size); j = j + b.Size {
			if b.BoardCells[j].GetMark() != mark {
				ret = false
			}
		}
		if ret {
			return ret
		}
	}
	return ret
}
func checkDiagonal(b *components.Board, mark string) bool {
	ret := true
	for i := uint8(0); i < b.Size; i++ {
		if b.BoardCells[b.Size*i+i].GetMark() != mark {
			ret = false
		}
	}
	if ret {
		return ret
	}
	ret = true
	for i := uint8(0); i < b.Size; i++ {
		if b.BoardCells[(b.Size*i)+(b.Size-1-i)].GetMark() != mark {
			ret = false
		}
	}
	return ret
}

func (b *ResultService) Result(player *components.Player) (bool, string) {

	if checkRows(b.Board, player.Mark) {
		return true, "win"
	} else if checkColumns(b.Board, player.Mark) {
		return true, "win"
	} else if checkDiagonal(b.Board, player.Mark) {
		return true, "win"
	} else if b.CheckBoardIsFull() {
		return true, "Draw"
	}
	return false, "Inprogress"
}
