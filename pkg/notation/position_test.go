package notation_test

import (
	"fmt"
	"gochess/pkg/notation"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPosition(t *testing.T) {
	startingPosition, err := notation.New(notation.StartingFEN)
	require.NoError(t, err)
	fmt.Println(startingPosition.String())

	assert.Equal(t, true, startingPosition.IsWhitesTurn())
	assert.Equal(t, true, startingPosition.CanCastle(false, false))
	assert.Equal(t, true, startingPosition.CanCastle(false, true))
	assert.Equal(t, true, startingPosition.CanCastle(true, false))
	assert.Equal(t, true, startingPosition.CanCastle(true, true))
	assert.Equal(t, 0, startingPosition.HalfmoveCount())
	assert.Equal(t, 1, startingPosition.FullmoveCount())
}

func BenchmarkPosition(b *testing.B) {
	b.Run("Just New", func(b *testing.B) {
		_, _ = notation.New(notation.StartingFEN)
	})
}
