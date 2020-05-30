package service

import (
	"errors"
	"example/tic_tac_toe_app/components"
	"testing"
)

func TestPlay(t *testing.T){
	tests:=[]struct{
		input *GameService
		input1 uint8
		expected error 
	}{
		{&GameService{&ResultService{&BoardService{&components.Board{ 3,[]*components.Cell{nil}}}},[2]*components.Player{nil,nil}},10, errors.New("please enter integer in range (0,8)")},
		{&GameService{&ResultService{&BoardService{&components.Board{2,[]*components.Cell{nil}}}},[2]*components.Player{nil,nil}},8, errors.New("please enter integer in range (0,3)")},

	}
	for _,test:=range tests{
		_,err:=test.input.Play(test.input1)
		if err.Error()!=test.expected.Error(){
			t.Error("fail")
		}
	}
}