package service

import (
	"testing"

	"example/tic_tac_toe_app/components"
)

func TestCheckRows(t *testing.T) {
	tests := []struct {
		input    *components.Board
		input1   string
		expected bool
	}{
		{&components.Board{
			Size: 2,
			BoardCells: []*components.Cell{
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
			},
		}, components.XMark, false},

		{&components.Board{
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
		}, components.XMark, true},
	}

	for _, test := range tests {
		if checkRows(test.input, test.input1) != test.expected {
			t.Error("check row failed")
		}
	}
}

func TestCheckColumn(t *testing.T) {
	tests := []struct {
		input    *components.Board
		input1   string
		expected bool
	}{
		{&components.Board{
			Size: 2,
			BoardCells: []*components.Cell{
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
			},
		}, components.XMark, true},

		{&components.Board{
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
		}, components.XMark, false},
	}

	for _, test := range tests {
		if checkColumns(test.input, test.input1) != test.expected {
			t.Error("check column failed")
		}
	}
}

func TestCheckDiagonal(t *testing.T) {
	tests := []struct {
		input    *components.Board
		input1   string
		expected bool
	}{
		{&components.Board{
			Size: 2,
			BoardCells: []*components.Cell{
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
			},
		}, components.XMark, true},

		{&components.Board{
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
		}, components.OMark, true},
	}

	for _, test := range tests {
		if checkDiagonal(test.input, test.input1) != test.expected {
			t.Error("check diagonal failed")
		}
	}
}
