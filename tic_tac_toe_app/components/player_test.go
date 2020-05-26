package components

import "testing"

func TestNewPlayer(t *testing.T) {
	tests := []struct {
		input1    string
		input2    string
		expected1 string
		expected2 string
	}{
		{"ayushman", OMark, "ayushman", OMark},
		{"fardin", XMark, "fardin", XMark},
	}
	for _, test := range tests {
		temp := NewPlayer(test.input1, test.input2)
		if temp.name != test.expected1 || temp.mark != test.expected2 {
			t.Error(temp, test.expected1, test.expected2)
		}
	}
}
