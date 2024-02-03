package notation_test

import (
	"fmt"
	"gochess/pkg/notation"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewPosition(t *testing.T) {
	fmt.Println(notation.StartingFEN)
	startingPosition, err := notation.NewPosition(notation.StartingFEN)
	require.NoError(t, err)
	fmt.Println(startingPosition.String())
}

func BenchmarkPosition(b *testing.B) {
	b.Run("Just New", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_, _ = notation.NewPosition(notation.StartingFEN)
		}
	})
}
