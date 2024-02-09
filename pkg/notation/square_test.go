package notation_test

import (
	"fmt"
	"gochess/pkg/notation"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSquaresInBetween(t *testing.T) {
	tests := []struct {
		name     string
		s1       notation.Square
		s2       notation.Square
		notEmpty bool
	}{
		{"Same Square", notation.Square_a1, notation.Square_a1, false},

		{"Same File", notation.Square_a1, notation.Square_a2, false},
		{"Same File", notation.Square_a1, notation.Square_a3, true},
		{"Same File", notation.Square_a3, notation.Square_a1, true},

		{"Same Rank", notation.Square_a1, notation.Square_b1, false},
		{"Same Rank", notation.Square_a1, notation.Square_c1, true},
		{"Same Rank", notation.Square_c1, notation.Square_a1, true},

		{"Positive Diagonal", notation.Square_a1, notation.Square_b2, false},
		{"Positive Diagonal", notation.Square_a1, notation.Square_c3, true},
		{"Positive Diagonal", notation.Square_c3, notation.Square_a1, true},

		{"Negative Diagonal", notation.Square_a3, notation.Square_b2, false},
		{"Negative Diagonal", notation.Square_a3, notation.Square_c1, true},
		{"Negative Diagonal", notation.Square_c1, notation.Square_a3, true},

		{"None", notation.Square_a1, notation.Square_b3, false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s:%s-%s", test.name, test.s1, test.s2), func(t *testing.T) {
			squares := notation.SquaresInBetween(test.s1, test.s2)
			fmt.Printf("%s - %s : %s\n", test.s1, test.s2, squares)
			if test.notEmpty {
				assert.NotEmpty(t, squares)
			} else {
				assert.Empty(t, squares)
			}
		})
	}
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

func TestSquareString(t *testing.T) {
	assert.Equal(t, notation.Square_a1.String(), "a1")
	assert.Equal(t, notation.Square_a2.String(), "a2")
	assert.Equal(t, notation.Square_h8.String(), "h8")
}
