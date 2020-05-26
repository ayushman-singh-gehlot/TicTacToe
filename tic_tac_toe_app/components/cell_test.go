package components

import (
	"errors"
	"testing"
)

func TestNewCell(t *testing.T) {
	test := NewCell()
	expected := NoMark
	if test.mark != expected {
		t.Error(test.mark, expected)
	}
}

func TestSetMark(t *testing.T) {
	tests := []struct {
		input1   *Cell
		input2   string
		expected error
	}{
		{&Cell{mark: NoMark}, XMark, nil},
		{&Cell{mark: XMark}, OMark, errors.New("cell is already marked")},
		{&Cell{mark: OMark}, NoMark, errors.New("cell is already marked")},
	}
	for _, test := range tests {
		actual := test.input1.SetMark(test.input2)
		if actual != nil {
			if actual.Error() != test.expected.Error() {
				t.Error(actual, test.expected)
			}
		}
	}
}

func TestGetMark(t *testing.T) {
	tests := []struct {
		input1   *Cell
		expected string
	}{
		{&Cell{mark: NoMark}, NoMark},
		{&Cell{mark: OMark}, OMark},
		{&Cell{mark: XMark}, XMark},
	}
	for _, test := range tests {
		actual := test.input1.GetMark()
		if actual != test.expected {
			t.Error(actual, test.expected)
		}
	}

}
