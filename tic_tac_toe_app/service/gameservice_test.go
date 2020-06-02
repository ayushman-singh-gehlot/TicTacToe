package service

import (
	"errors"
	"example/tic_tac_toe_app/components"
	"testing"
)

func TestPlay(t *testing.T) {
	tests := []struct {
		input         *GameService
		input1        uint8
		expected1     error
		expectedmark1 string
		input2        uint8
		expected2     error
		expectedmark2 string
	}{
		{&GameService{&ResultService{&BoardService{&components.Board{
			Size: 2,
			BoardCells: []*components.Cell{
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
			},
		}}}, [2]*components.Player{
			{Name: "ayushman", Mark: components.OMark},
			{Name: "fardin", Mark: components.XMark},
		}, 1}, 6, errors.New("please enter integer in range (0,3)"), "", 3, nil, components.XMark},

		{&GameService{&ResultService{&BoardService{&components.Board{
			Size: 3,
			BoardCells: []*components.Cell{
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
			},
		}}}, [2]*components.Player{
			{Name: "ayushman", Mark: components.OMark},
			{Name: "fardin", Mark: components.XMark},
		}, 0}, 7, nil, components.OMark, 4, errors.New("cell is already marked"), ""},
	}
	for _, test := range tests {
		_, err := test.input.Play(test.input1)
		if err != nil {
			if err.Error() != test.expected1.Error() {
				t.Error("errors didn't matched", err.Error(), test.expected1.Error())
			}
		} else if test.input.BoardCells[test.input1].GetMark() != test.expectedmark1 {
			t.Error("mark didn't matched")
		}
		_, err = test.input.Play(test.input2)
		if err != nil {
			if err.Error() != test.expected2.Error() {
				t.Error("errors didn't matched", err.Error(), test.expected2.Error())
			}
		} else if test.input.BoardCells[test.input2].GetMark() != test.expectedmark2 {
			t.Error("mark didn't matched")
		}

	}
}
