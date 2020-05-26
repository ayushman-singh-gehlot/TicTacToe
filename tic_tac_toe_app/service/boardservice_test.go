package service

import (
	"errors"
	"example/tic_tac_toe_app/components"
	"testing"
)

func TestNewServiceBoard(t *testing.T) {
	tests := []struct {
		input1    uint8
		expected  uint8
		expected1 string
	}{
		{2, 4, components.NoMark},
		{3, 9, components.NoMark},
		{4, 16, components.NoMark},
	}
	for _, test := range tests {
		actual := NewBoardService(test.input1)
		if len(actual.BoardCells) != int(test.expected) {
			t.Error(len(actual.BoardCells), test.expected)
		}
		for _, cell := range actual.BoardCells {
			if cell.Mark != test.expected1 {
				t.Error(actual, test.expected1)
			}

		}
	}
}

func TestPutMarkInPosition(t *testing.T) {
	tests := []struct {
		input1   *BoardService
		index    uint8
		mark     string
		expected error
	}{
		{&BoardService{components.CreateBoard(3)}, 0, components.OMark, nil},
		{&BoardService{&components.Board{
			Size: 2,
			BoardCells: []*components.Cell{
				&components.Cell{Mark: components.OMark},
				&components.Cell{Mark: components.XMark},
				&components.Cell{Mark: components.NoMark},
				&components.Cell{Mark: components.XMark},
			},
		},
		}, 0, components.XMark, errors.New("cell is already marked")},
	}
	for _, test := range tests {
		actual := test.input1.PutMarkInPosition(test.index, test.mark)
		if actual != nil {
			if actual.Error() != test.expected.Error() {
				t.Error(actual, test.expected)
			}
		}
	}
}

func TestPrintBoard(t *testing.T) {
	tests := []struct {
		input1   *BoardService
		expected string
	}{
		{&BoardService{components.CreateBoard(3)}, "\n\t- - - \n\t- - - \n\t- - - "},
		{&BoardService{&components.Board{
			Size: 2,
			BoardCells: []*components.Cell{
				&components.Cell{Mark: components.OMark},
				&components.Cell{Mark: components.XMark},
				&components.Cell{Mark: components.NoMark},
				&components.Cell{Mark: components.XMark},
			},
		},
		}, "\n\tO X \n\t- X "},
	}
	for _, test := range tests {
		actual := test.input1.PrintBoard()
		if actual != test.expected {
			t.Error(actual, test.expected)
		}
	}
}

func TestCheckBoardIsFull(t *testing.T) {
	tests := []struct {
		input1   *BoardService
		expected bool
	}{
		{&BoardService{components.CreateBoard(3)}, false},
		{&BoardService{&components.Board{
			Size: 2,
			BoardCells: []*components.Cell{
				&components.Cell{Mark: components.OMark},
				&components.Cell{Mark: components.XMark},
				&components.Cell{Mark: components.OMark},
				&components.Cell{Mark: components.XMark},
			},
		},
		}, true},
	}
	for _, test := range tests {
		actual := test.input1.CheckBoardIsFull()
		if actual != test.expected {
			t.Error(actual, test.expected)
		}
	}
}
