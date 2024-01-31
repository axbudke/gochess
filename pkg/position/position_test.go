package position_test

import (
	"gochess/pkg/position"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBoard(t *testing.T) {
	startingPosition, err := position.New(position.StartingFEN)
	require.NoError(t, err)

	assert.Equal(t, true, startingPosition.IsWhitesTurn())
	assert.Equal(t, true, startingPosition.CanCastle(false, false))
	assert.Equal(t, true, startingPosition.CanCastle(false, true))
	assert.Equal(t, true, startingPosition.CanCastle(true, false))
	assert.Equal(t, true, startingPosition.CanCastle(true, true))
	assert.Equal(t, 0, startingPosition.HalfmoveCount())
	assert.Equal(t, 1, startingPosition.FullmoveCount())
}

func BenchmarkBoard(b *testing.B) {
	b.Run("Just New", func(b *testing.B) {
		_, _ = position.New(position.StartingFEN)
	})
}
