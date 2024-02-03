package notation_test

import (
	"gochess/pkg/notation"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSquareString(t *testing.T) {
	assert.Equal(t, notation.Square_a1.String(), "a1")
	assert.Equal(t, notation.Square_a2.String(), "a2")
	assert.Equal(t, notation.Square_h8.String(), "h8")
}

func TestNewSquare(t *testing.T) {
	a1, err := notation.NewSquare(0, 0)
	assert.NoError(t, err)
	assert.Equal(t, "a1", a1.String())
	h8, err := notation.NewSquare(7, 7)
	assert.NoError(t, err)
	assert.Equal(t, "h8", h8.String())
	_, err = notation.NewSquare(-1, 0)
	assert.ErrorContains(t, err, "invalid rank")
	_, err = notation.NewSquare(8, 0)
	assert.ErrorContains(t, err, "invalid rank")
	_, err = notation.NewSquare(0, -1)
	assert.ErrorContains(t, err, "invalid file")
	_, err = notation.NewSquare(0, 8)
	assert.ErrorContains(t, err, "invalid file")
}

func TestNewSquareFromString(t *testing.T) {
	a2, err := notation.NewSquareFromString("a2")
	require.NoError(t, err)
	assert.Equal(t, a2, notation.Square_a2)
	f3, err := notation.NewSquareFromString("f3")
	require.NoError(t, err)
	assert.Equal(t, f3, notation.Square_f3)
	_, err = notation.NewSquareFromString("k9")
	assert.ErrorContains(t, err, "invalid square format")
}
