package components

import (
	"errors"
)

type Cell struct {
	mark string
}

const (
	NoMark = "-"
	XMark  = "X"
	OMark  = "O"
)

func NewCell() *Cell {
	return &Cell{mark: NoMark}
}

func (c *Cell) SetMark(mark string) error {
	if c.mark == NoMark {
		c.mark = mark
		return nil
	}
	return errors.New("cell is already marked ")
}

func (c *Cell) GetMark() string {
	return cell.mark
}
