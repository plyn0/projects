package route

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func roundFloat(a float64) float64 {
	return a
}

func TestRoundFloat(t *testing.T) {
	assert.Equal(t, roundFloat(1.234), 1.23)
	assert.Equal(t, roundFloat(1.400001), 1.4) // 1 decimal result
}
