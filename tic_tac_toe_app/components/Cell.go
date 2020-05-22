package components

import (
	"errors"
)

type Cell struct {
	mark string
}

const (
	NoMark = " "
	XMark  = "X"
	OMark  = "O"
)

func NewCell() *Cell {
	return &Cell{mark: NoMark}
}

func SetMark(cell *Cell, mark string) error {
	if cell.mark == NoMark {
		cell.mark = mark
		return nil
	}
	return errors.New("cell is already marked ")
}

func GetMark(cell *Cell) string {
	return cell.mark
}
