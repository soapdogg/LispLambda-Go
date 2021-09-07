package parser

import (
	"github.com/stretchr/testify/assert"
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"testing"
)

func TestParseExpressionListNode(t *testing.T) {
	openToken := constants.OPEN_TOKEN_STR
	openToken2 := constants.OPEN_TOKEN_STR
	literalToken := "value"
	closeToken := constants.CLOSE_TOKEN_STR
	closeToken2 := constants.CLOSE_TOKEN_STR

	tokens := []string{
		openToken,
		openToken2,
		literalToken,
		closeToken,
		closeToken2,
	}

	nodeParser := NewNodeParser()
	actual := nodeParser.ParseIntoNode(tokens)

	expressionListActual := actual.(datamodels.ExpressionListNode)

	assert.Equal(t, 2, len(expressionListActual.GetChildren()))
	expressionListChild := expressionListActual.GetChildren()[0].(datamodels.ExpressionListNode)
	assert.Equal(t, 2, len(expressionListChild.GetChildren()))
	grandchild1 := expressionListChild.GetChildren()[0].(datamodels.AtomNode)
	assert.Equal(t, literalToken, grandchild1.GetValue())
	grandchild2 := expressionListChild.GetChildren()[1].(datamodels.AtomNode)
	assert.Equal(t, constants.NIL, grandchild2.GetValue())
	child2 := expressionListActual.GetChildren()[1].(datamodels.AtomNode)
	assert.Equal(t, constants.NIL, child2.GetValue())
}

func TestParseAtomNode(t *testing.T) {
	literal := "literal"
	tokens := [] string {literal}
	nodeParser := NewNodeParser()
	actual := nodeParser.ParseIntoNode(tokens)
	actualAtomNode := actual.(datamodels.AtomNode)
	assert.Equal(t, literal, actualAtomNode.GetValue())
}
