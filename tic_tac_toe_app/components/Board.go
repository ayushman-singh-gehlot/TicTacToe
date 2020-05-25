package components

type Board struct {
	size       uint8
	boardCells []*Cell
}

func CreateBoard(size uint8) *Board {
	var board Board
	board.size = size
	for i := 0; i < int(board.size*board.size); i++ {
		board.boardCells = append(board.boardCells, NewCell())
	}
	return &board
}
