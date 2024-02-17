package generation_test

import (
	"fmt"
	"testing"

	"gochess/pkg/generation"
	"gochess/pkg/notation/position"
	"gochess/pkg/notation/square"

	"github.com/stretchr/testify/require"
)

func TestGenerateSudoLegalMoves(t *testing.T) {
	fmt.Printf("FEN: %s\n", position.StartingFEN)
	startingPosition, err := position.NewPosition(position.StartingFEN)
	require.NoError(t, err)
	moves := generation.GeneratePseudoLegalMoves(startingPosition)
	fmt.Printf("Moves: %s\n", moves)

	oneMoveFen := position.FEN("rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq - 0 1")
	fmt.Printf("FEN: %s\n", oneMoveFen)
	oneMovePosition, err := position.NewPosition(oneMoveFen)
	require.NoError(t, err)
	moves = generation.GeneratePseudoLegalMoves(oneMovePosition)
	fmt.Printf("Moves: %s\n", moves)
}

func BenchmarkGenerateSudoLegalMoves(b *testing.B) {
	startingBoard, err := position.NewPosition(position.StartingFEN)
	require.NoError(b, err)
	b.Run("GenerateSudoLegalMoves", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = generation.GeneratePseudoLegalMoves(startingBoard)
		}
	})
	b.Run("GenerateNoSlideMoves", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = generation.GenerateNoSlideMoves(startingBoard, square.Square_b3, generation.KnightMovementPairs)
		}
	})
	b.Run("GenerateMove", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = generation.GenerateMove(startingBoard, square.Square_a2, square.Square_a3)
		}
	})
}
