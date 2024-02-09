package generation_test

import (
	"fmt"
	"gochess/pkg/generation"
	"gochess/pkg/notation"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateSudoLegalMoves(t *testing.T) {
	fmt.Printf("FEN: %s\n", notation.StartingFEN)
	startingPosition, err := notation.NewPosition(notation.StartingFEN)
	require.NoError(t, err)
	moves := generation.GeneratePseudoLegalMoves(startingPosition)
	fmt.Printf("Moves: %s\n", moves)

	oneMoveFen := notation.FEN("rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq - 0 1")
	fmt.Printf("FEN: %s\n", oneMoveFen)
	oneMovePosition, err := notation.NewPosition(oneMoveFen)
	require.NoError(t, err)
	moves = generation.GeneratePseudoLegalMoves(oneMovePosition)
	fmt.Printf("Moves: %s\n", moves)
}

func BenchmarkGenerateSudoLegalMoves(b *testing.B) {
	startingBoard, err := notation.NewPosition(notation.StartingFEN)
	require.NoError(b, err)
	b.Run("GenerateSudoLegalMoves", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = generation.GeneratePseudoLegalMoves(startingBoard)
		}
	})
	b.Run("GenerateNoSlideMoves", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = generation.GenerateNoSlideMoves(startingBoard, notation.Square_b3, generation.KnightMovementPairs)
		}
	})
	b.Run("GenerateMove", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = generation.GenerateMove(startingBoard, notation.Square_a2, notation.Square_a3)
		}
	})
}
