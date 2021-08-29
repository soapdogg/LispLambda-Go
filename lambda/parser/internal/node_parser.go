package internal

import (
	"lisp_lambda-go/lambda/core/constants"
	"lisp_lambda-go/lambda/core/datamodels"
)

type NodeParser interface {
	ParseIntoNode(tokens []string) datamodels.Node
}

type nodeParser struct {}

func NewNodeParser() *nodeParser {
	return &nodeParser{}
}

func (n *nodeParser) ParseIntoNode(tokens []string) datamodels.Node{
	if len(tokens) == 1 {
		return datamodels.NewAtomNode(
			tokens[0],
		)
	}

	stack := datamodels.NewNodeListStack()
	result := datamodels.NewExpressionListNode(
		[]datamodels.Node{},
	)

	for _, token := range tokens {
		if token == constants.OPEN_TOKEN_STR {
			stack.Push([]datamodels.Node{})
		} else if token == constants.CLOSE_TOKEN_STR {
			node := datamodels.NewAtomNode(constants.NIL)
			top := stack.Pop()
			top = append(top, node)
			result = datamodels.NewExpressionListNode(top)
			if stack.IsNotEmpty() {
				top = stack.Pop()
				top = append(top, result)
				stack.Push(top)
			}
		} else {
			node := datamodels.NewAtomNode(token)
			top := stack.Pop()
			top = append(top, node)
			stack.Push(top)
		}
	}

	return result
}

var _ NodeParser = &nodeParser{}