package notation_test

import (
	"gochess/pkg/notation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquare(t *testing.T) {
	assert.Equal(t, notation.Square_a1.String(), "a1")
	assert.Equal(t, notation.Square_h8.String(), "h8")

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
