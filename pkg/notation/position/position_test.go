package position_test

import (
	"fmt"
	"testing"

	"gochess/pkg/notation/position"

	"github.com/stretchr/testify/require"
)

func TestNewPosition(t *testing.T) {
	fmt.Println(position.StartingFEN)
	startingPosition, err := position.NewPosition(position.StartingFEN)
	require.NoError(t, err)
	fmt.Println(startingPosition.String())
}

func BenchmarkPosition(b *testing.B) {
	b.Run("Just New", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_, _ = position.NewPosition(position.StartingFEN)
		}
	})
}
