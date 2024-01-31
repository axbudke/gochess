package generation_test

import (
	"fmt"
	"gochess/pkg/generation"
	"gochess/pkg/position"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateSudoLegalMoves(t *testing.T) {
	fmt.Printf("FEN: %s\n", position.StartingFEN)
	startingPosition, err := position.New(position.StartingFEN)
	require.NoError(t, err)
	moves := generation.GenerateSudoLegalMoves(startingPosition)
	fmt.Printf("Moves: %s\n", moves)

	oneMoveFen := position.FEN("rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq - 0 1")
	fmt.Printf("FEN: %s\n", oneMoveFen)
	oneMovePosition, err := position.New(oneMoveFen)
	require.NoError(t, err)
	moves = generation.GenerateSudoLegalMoves(oneMovePosition)
	fmt.Printf("Moves: %s\n", moves)
}

func BenchmarkGenerateSudoLegalMoves(b *testing.B) {
	startingBoard, err := position.New(position.StartingFEN)
	require.NoError(b, err)
	b.Run("GenerateSudoLegalMoves", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = generation.GenerateSudoLegalMoves(startingBoard)
		}
	})
	b.Run("GenerateNoSlideMoves", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = generation.GenerateNoSlideMoves(startingBoard, position.Square_b3, generation.KnightMovementPairs)
		}
	})
	b.Run("GenerateMove", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_, _ = generation.GenerateMove(startingBoard, position.Square_a2, position.Rank3, position.FileA)
		}
	})
}