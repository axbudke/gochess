package evaluation_test

import (
	"gochess/pkg/board"
	"gochess/pkg/evaluation"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEvaluate(t *testing.T) {
	startingBoard, err := board.New(board.StartingFEN)
	require.NoError(t, err)

	assert.Equal(t, 0, evaluation.Evaluate(startingBoard))
	assert.Equal(t, 0, evaluation.GetMaterialCount(startingBoard))
}
