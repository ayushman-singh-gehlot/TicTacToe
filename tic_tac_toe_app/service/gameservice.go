package service

import (
	"example/tic_tac_toe_app/components"
	"fmt"
)

type GameService struct {
	*ResultService
	player1 *components.Player
	player2 *components.Player
}

var turn int = 0

func NewGameService(rs *ResultService, pl1, pl2 *components.Player) *GameService {
	return &GameService{rs, pl1, pl2}

}

func (gs *GameService) Play(pos uint8) (Result, error) {
	if pos < 0 || pos >= gs.Size {
		return Result{false, false}, fmt.Errorf("please enter integer in range (0,%d)", gs.Size)
	}
	var res Result
	if turn%2 == 0 {
		err := gs.PutMarkInPosition(gs.player1, pos)
		if err != nil {
			return Result{false, false}, err
		}
		res = gs.GetResult(gs.player1, pos)
	} else if turn%2 == 1 {
		err := gs.PutMarkInPosition(gs.player2, pos)
		if err != nil {
			return Result{false, false}, err
		}
		res = gs.GetResult(gs.player2, pos)
	}
	turn++
	return res, nil
}