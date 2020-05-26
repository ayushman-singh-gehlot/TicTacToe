package components

import "testing"

func TestCreateBoard(t *testing.T) {
	tests := []struct {
		input1    uint8
		expected  int
		expected1 string
	}{
		{2, 4, NoMark},
		{3, 9, NoMark},
		{4, 16, NoMark},
	}
	for _, test := range tests {
		actual := CreateBoard(test.input1)
		if len(actual.BoardCells) != test.expected {
			t.Error(len(actual.BoardCells), test.expected)
		}
		for _, cell := range actual.BoardCells {
			if cell.mark != test.expected1 {
				t.Error(actual, test.expected1)
			}

		}
	}
}
