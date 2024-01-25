package board_test

import (
	"gochess/pkg/board"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFEN(t *testing.T) {
	startingBoard, err := board.New(board.StartingFEN)
	require.NoError(t, err)

	assert.Equal(t, true, startingBoard.IsWhitesTurn())
	assert.Equal(t, true, startingBoard.CanCastle(false, false))
	assert.Equal(t, true, startingBoard.CanCastle(false, true))
	assert.Equal(t, true, startingBoard.CanCastle(true, false))
	assert.Equal(t, true, startingBoard.CanCastle(true, true))
	assert.Equal(t, 0, startingBoard.HalfmoveCount())
	assert.Equal(t, 1, startingBoard.FullmoveCount())
}
