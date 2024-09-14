package token

import (
	"testing"
)

var tests = []struct {
	char     byte
	expected bool
}{
	{'+', true},
	{'=', true},
	{',', true},
	{'}', true},
	{'a', false},
	{'1', false},
}

func TestIsSymbol(t *testing.T) {
	for i, value := range tests {
		result := IsSymbol(value.char)
		if result != value.expected {
			t.Fatalf("tests[%d] - result wrong. expected=%t, got=%t",
				i, value.expected, result)
		}
	}
}

// same test for the recursive version
func TestIsSymbolRec(t *testing.T) {
	for i, value := range tests {
		result := IsSymbolRec(value.char)
		if result != value.expected {
			t.Fatalf("tests[%d] - result wrong. expected=%t, got=%t",
				i, value.expected, result)
		}
	}
}
