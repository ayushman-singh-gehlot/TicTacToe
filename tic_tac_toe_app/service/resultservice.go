package service

import (
	"example/tic_tac_toe_app/components"
)

type ResultService struct {
	*BoardService
}

func NewResultService(boardService *BoardService) *ResultService {
	return &ResultService{boardService}
}

func (rs *ResultService) checkRows(mark string, pos uint8) bool {
	ret := true
	i := (pos - pos%rs.Size)
	//fmt.Printf("checking row %d", i/rs.Size+1)
	for j := i; j < (i + rs.Size); j++ {
		if rs.BoardCells[j].GetMark() != mark {
			ret = false
		}
	}
	return ret
}

func (rs *ResultService) checkColumns(mark string, pos uint8) bool {
	ret := true
	//fmt.Printf("checking column %d", pos%rs.Size+1)
	for j := pos % rs.Size; j < (rs.Size * rs.Size); j = j + rs.Size {
		if rs.BoardCells[j].GetMark() != mark {
			ret = false
		}
	}
	return ret
}
func (rs *ResultService) checkDiagonal(mark string) bool {
	ret := true
	for i := uint8(0); i < rs.Size; i++ {
		if rs.BoardCells[rs.Size*i+i].GetMark() != mark {
			ret = false
		}
	}
	if ret {
		return ret
	}
	ret = true
	for i := uint8(0); i < rs.Size; i++ {
		if rs.BoardCells[(rs.Size*i)+(rs.Size-1-i)].GetMark() != mark {
			ret = false
		}
	}
	return ret
}

func (rs *ResultService) Result(player *components.Player, pos uint8) (bool, string) {

	if rs.checkRows(player.Mark, pos) {
		return true, "win"
	} else if rs.checkColumns(player.Mark, pos) {
		return true, "win"
	} else if rs.checkDiagonal(player.Mark) {
		return true, "win"
	} else if rs.CheckBoardIsFull() {
		return true, "Draw"
	}
	return false, "Inprogress"
}
