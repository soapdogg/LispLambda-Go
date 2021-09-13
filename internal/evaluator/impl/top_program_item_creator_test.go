package impl

import (
	"github.com/stretchr/testify/assert"
	"lisp_lambda-go/internal/core/datamodels"
	"testing"
)

func TestCreateTopProgramItem(t *testing.T) {
	functionName := "functionName"
	atomNode := datamodels.NewAtomNode(functionName)
	expressionListNode := datamodels.NewExpressionListNode([]datamodels.Node{atomNode})
	variableMap := map[string]datamodels.Node{}
	programStack := datamodels.NewProgramItemStack()

	topProgramItemCreator := NewTopProgramItemCreator()

	topProgramItemCreator.CreateTopProgramItem(
		expressionListNode,
		variableMap,
		programStack,
	)

	assert.Equal(t, 1, programStack.Size())
}