package service

import (
	"testing"

	"example/tic_tac_toe_app/components"
)

func TestCheckRows(t *testing.T) {
	tests := []struct {
		input    *ResultService
		input1   string
		input2   uint8
		expected bool
	}{
		{&ResultService{&BoardService{&components.Board{
			Size: 2,
			BoardCells: []*components.Cell{
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
			},
		},
		},
		}, components.XMark, 3, false},

		{&ResultService{&BoardService{&components.Board{
			Size: 3,
			BoardCells: []*components.Cell{
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
				{Mark: components.XMark},
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.NoMark},
				{Mark: components.OMark},
			},
		},
		},
		}, components.XMark, 4, true},
	}

	for _, test := range tests {
		if test.input.checkRows(test.input1, test.input2) != test.expected {
			t.Error("check row failed")
		}
	}
}

func TestCheckColumns(t *testing.T) {
	tests := []struct {
		input    *ResultService
		input1   string
		input2   uint8
		expected bool
	}{
		{&ResultService{&BoardService{&components.Board{
			Size: 2,
			BoardCells: []*components.Cell{
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
			},
		},
		},
		}, components.XMark, 3, true},

		{&ResultService{&BoardService{&components.Board{
			Size: 3,
			BoardCells: []*components.Cell{
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
				{Mark: components.XMark},
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.NoMark},
				{Mark: components.OMark},
			},
		},
		},
		}, components.XMark, 5, false},
	}

	for _, test := range tests {
		if test.input.checkColumns(test.input1, test.input2) != test.expected {
			t.Error("check column failed")
		}
	}
}

func TestCheckDiagonal(t *testing.T) {
	tests := []struct {
		input    *ResultService
		input1   string
		expected bool
	}{
		{&ResultService{&BoardService{&components.Board{
			Size: 2,
			BoardCells: []*components.Cell{
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
			},
		},
		},
		}, components.XMark, true},

		{&ResultService{&BoardService{&components.Board{
			Size: 3,
			BoardCells: []*components.Cell{
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.NoMark},
				{Mark: components.OMark},
			},
		},
		},
		}, components.OMark, true},
	}

	for _, test := range tests {
		if test.input.checkDiagonal(test.input1) != test.expected {
			t.Error(test, "check diagonal failed")
		}
	}
}
