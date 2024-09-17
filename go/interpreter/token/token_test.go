package token

import (
	"testing"

	"github.com/stretchr/testify/require"
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
	for _, value := range tests {
		result := IsSymbol(value.char)
		require.Equal(t, value.expected, result)
	}
}

// same test for the recursive version
func TestIsSymbolRec(t *testing.T) {
	for _, value := range tests {
		result := IsSymbolRec(value.char)
		require.Equal(t, value.expected, result)
	}
}
