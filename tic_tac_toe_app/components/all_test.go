package components

import (
	"errors"
	"fmt"
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
		{&Cell{mark: XMark}, OMark, errors.New("cell is already marked")},  // why i m getting error?
		{&Cell{mark: OMark}, NoMark, errors.New("cell is already marked")}, // why i m getting error?
	}
	for _, test := range tests {
		actual := SetMark(test.input1, test.input2)
		fmt.Println(actual)
		fmt.Println(test.expected)
		if actual != test.expected {
			t.Error(actual, test.expected)
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
		actual := GetMark(test.input1)
		if actual != test.expected {
			t.Error(actual, test.expected)
		}
	}

}
