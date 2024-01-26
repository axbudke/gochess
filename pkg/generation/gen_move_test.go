package generation_test

import (
	"fmt"
	"gochess/pkg/board"
	"gochess/pkg/generation"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateSudoLegalMoves(t *testing.T) {
	startingBoard, err := board.New(board.StartingFEN)
	require.NoError(t, err)

	moves := generation.GenerateSudoLegalMoves(startingBoard)
	fmt.Printf("Moves: %s\n", moves)
}

func BenchmarkGenerateSudoLegalMoves(b *testing.B) {

}
