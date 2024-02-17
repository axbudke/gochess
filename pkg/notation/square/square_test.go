package square_test

import (
	"fmt"
	"gochess/pkg/notation/square"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkSquare(b *testing.B) {
	b.Run("NewSquare", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			for f := square.FileA; f <= square.FileH; f++ {
				for r := square.Rank1; r <= square.Rank8; r++ {
					_ = square.NewSquare(f, r)
				}
			}
		}
	})

	b.Run("NewSquareCheck", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			for f := square.FileA; f <= square.FileH; f++ {
				for r := square.Rank1; r <= square.Rank8; r++ {
					_, _ = square.NewSquareCheck(f, r)
				}
			}
		}
	})
}

func TestSquaresInBetween(t *testing.T) {
	tests := []struct {
		name     string
		s1       square.Square
		s2       square.Square
		notEmpty bool
	}{
		{"Same Square", square.Square_a1, square.Square_a1, false},

		{"Same File", square.Square_a1, square.Square_a2, false},
		{"Same File", square.Square_a1, square.Square_a3, true},
		{"Same File", square.Square_a3, square.Square_a1, true},

		{"Same Rank", square.Square_a1, square.Square_b1, false},
		{"Same Rank", square.Square_a1, square.Square_c1, true},
		{"Same Rank", square.Square_c1, square.Square_a1, true},

		{"Positive Diagonal", square.Square_a1, square.Square_b2, false},
		{"Positive Diagonal", square.Square_a1, square.Square_c3, true},
		{"Positive Diagonal", square.Square_c3, square.Square_a1, true},

		{"Negative Diagonal", square.Square_a3, square.Square_b2, false},
		{"Negative Diagonal", square.Square_a3, square.Square_c1, true},
		{"Negative Diagonal", square.Square_c1, square.Square_a3, true},

		{"None", square.Square_a1, square.Square_b3, false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s:%s-%s", test.name, test.s1, test.s2), func(t *testing.T) {
			squares := square.SquaresInBetween(test.s1, test.s2)
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
	a1, err := square.NewSquareCheck(0, 0)
	assert.NoError(t, err)
	assert.Equal(t, "a1", a1.String())
	h8, err := square.NewSquareCheck(7, 7)
	assert.NoError(t, err)
	assert.Equal(t, "h8", h8.String())

	// Failures
	_, err = square.NewSquareCheck(-1, 0)
	assert.ErrorContains(t, err, "invalid file")
	_, err = square.NewSquareCheck(8, 0)
	assert.ErrorContains(t, err, "invalid file")
	_, err = square.NewSquareCheck(0, -1)
	assert.ErrorContains(t, err, "invalid rank")
	_, err = square.NewSquareCheck(0, 8)
	assert.ErrorContains(t, err, "invalid rank")
}

func TestNewSquareFromString(t *testing.T) {
	a2, err := square.NewSquareFromString("a2")
	require.NoError(t, err)
	assert.Equal(t, a2, square.Square_a2)
	f3, err := square.NewSquareFromString("f3")
	require.NoError(t, err)
	assert.Equal(t, f3, square.Square_f3)

	// Failures
	_, err = square.NewSquareFromString("k9")
	assert.ErrorContains(t, err, "invalid file")
	_, err = square.NewSquareFromString("a9")
	assert.ErrorContains(t, err, "invalid rank")
	_, err = square.NewSquareFromString("k92")
	assert.ErrorContains(t, err, "invalid square")
}

func TestSquareString(t *testing.T) {
	assert.Equal(t, "a1", square.Square_a1.String())
	assert.Equal(t, "a2", square.Square_a2.String())
	assert.Equal(t, "h8", square.Square_h8.String())
}
