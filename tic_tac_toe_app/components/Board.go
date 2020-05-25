package components

type Board struct {
	Size       uint8
	BoardCells []*Cell
}

func CreateBoard(size uint8) *Board {
	var board Board
	board.Size = size
	for i := 0; i < int(board.Size*board.Size); i++ {
		board.BoardCells = append(board.BoardCells, NewCell())
	}
	return &board
}

// func SetBoardCell(index uint8, mark string) {

// }
