package service

import (
	"example/tic_tac_toe_app/components"
	"fmt"
)

type GameService struct {
	*ResultService
	players [2]*components.Player
}

var turn int = 0

func NewGameService(rs *ResultService, players [2]*components.Player) *GameService {
	return &GameService{rs, players}

}

func (gs *GameService) Play(pos uint8) (Result, error) {
	if pos < 0 || pos >= gs.Size*gs.Size {
		return Result{"", false, false}, fmt.Errorf("please enter integer in range (0,%d)", gs.Size*gs.Size-1)
	}
	var res Result
	if turn%2 == 0 {
		err := gs.PutMarkInPosition(gs.players[0], pos)
		if err != nil {
			return Result{"", false, false}, err
		}
		res = gs.GetResult(gs.players[0], pos)
	} else if turn%2 == 1 {
		err := gs.PutMarkInPosition(gs.players[1], pos)
		if err != nil {
			return Result{"", false, false}, err
		}
		res = gs.GetResult(gs.players[1], pos)
	}
	turn++
	return res, nil
}
