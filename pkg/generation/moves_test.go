package generation_test

import (
	"fmt"
	"gochess/pkg/board"
	"gochess/pkg/generation"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateMoves(t *testing.T) {
	startingBoard, err := board.New(board.StartingFEN)
	require.NoError(t, err)

	moves := generation.GenerateMoves(startingBoard)
	fmt.Printf("Moves: %+v\n", moves)
}
