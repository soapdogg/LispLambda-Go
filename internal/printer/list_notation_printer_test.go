package printer

import (
	"github.com/stretchr/testify/assert"
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"testing"
)

func TestPrintNonEmptyListOfNodes(t *testing.T){
	listNotationPrinter := NewListNotationPrinter()

	value := "value"
	atomNode := datamodels.NewAtomNode(value)
	nodes := []datamodels.Node{atomNode}
	actual := listNotationPrinter.PrintAllInListNotation(nodes)
	assert.Equal(t, value, actual)
}

func TestPrintExpressionListOfSizeTwo(t *testing.T) {
	listNotationPrinter := NewListNotationPrinter()
	
	c0Value := "c0"
	c0 := datamodels.NewAtomNode(c0Value)
	
	c1Value := "c1"
	c1 := datamodels.NewAtomNode(c1Value)
	
	children := []datamodels.Node{c0, c1}
	node := datamodels.NewExpressionListNode(children)
	
	expected := constants.OPEN_TOKEN_STR + c0Value + constants.ListDelimiter + c1Value + constants.CLOSE_TOKEN_STR
	
	actual := listNotationPrinter.PrintInListNotation(node)
	
	assert.Equal(t, expected, actual)
}

func TestExpressionListOfSizeTwoLastElementIsNil(t *testing.T){
	listNotationPrinter := NewListNotationPrinter()

	c0Value := "c0"
	c0 := datamodels.NewAtomNode(c0Value)

	c1Value := constants.NIL
	c1 := datamodels.NewAtomNode(c1Value)

	children := []datamodels.Node{c0, c1}
	node := datamodels.NewExpressionListNode(children)

	expected := constants.OPEN_TOKEN_STR + c0Value + constants.CLOSE_TOKEN_STR

	actual := listNotationPrinter.PrintInListNotation(node)

	assert.Equal(t, expected, actual)
}