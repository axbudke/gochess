package notation_test

import (
	"fmt"
	"gochess/pkg/notation"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkSquare(b *testing.B) {
	b.Run("NewSquare", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			for f := notation.FileA; f <= notation.FileH; f++ {
				for r := notation.Rank1; r <= notation.Rank8; r++ {
					_ = notation.NewSquare(f, r)
				}
			}
		}
	})

	b.Run("NewSquareCheck", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			for f := notation.FileA; f <= notation.FileH; f++ {
				for r := notation.Rank1; r <= notation.Rank8; r++ {
					_, _ = notation.NewSquareCheck(f, r)
				}
			}
		}
	})
}

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

func TestNewSquareCheck(t *testing.T) {
	a1, err := notation.NewSquareCheck(0, 0)
	assert.NoError(t, err)
	assert.Equal(t, "a1", a1.String())
	h8, err := notation.NewSquareCheck(7, 7)
	assert.NoError(t, err)
	assert.Equal(t, "h8", h8.String())

	// Failures
	_, err = notation.NewSquareCheck(-1, 0)
	assert.ErrorContains(t, err, "invalid file")
	_, err = notation.NewSquareCheck(8, 0)
	assert.ErrorContains(t, err, "invalid file")
	_, err = notation.NewSquareCheck(0, -1)
	assert.ErrorContains(t, err, "invalid rank")
	_, err = notation.NewSquareCheck(0, 8)
	assert.ErrorContains(t, err, "invalid rank")
}

func TestNewSquareFromString(t *testing.T) {
	a2, err := notation.NewSquareFromString("a2")
	require.NoError(t, err)
	assert.Equal(t, a2, notation.Square_a2)
	f3, err := notation.NewSquareFromString("f3")
	require.NoError(t, err)
	assert.Equal(t, f3, notation.Square_f3)

	// Failures
	_, err = notation.NewSquareFromString("k9")
	assert.ErrorContains(t, err, "invalid file")
	_, err = notation.NewSquareFromString("a9")
	assert.ErrorContains(t, err, "invalid rank")
	_, err = notation.NewSquareFromString("k92")
	assert.ErrorContains(t, err, "invalid square")
}

func TestSquareString(t *testing.T) {
	assert.Equal(t, "a1", notation.Square_a1.String())
	assert.Equal(t, "a2", notation.Square_a2.String())
	assert.Equal(t, "h8", notation.Square_h8.String())
}
