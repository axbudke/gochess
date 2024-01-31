package evaluation_test

import (
	"gochess/pkg/evaluation"
	"gochess/pkg/position"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEvaluate(t *testing.T) {
	startingPosition, err := position.New(position.StartingFEN)
	require.NoError(t, err)

	assert.Equal(t, 0, evaluation.Evaluate(startingPosition))
	assert.Equal(t, 0, evaluation.GetMaterialCount(startingPosition))
}
