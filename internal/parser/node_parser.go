package parser

import (
	constants2 "lisp_lambda-go/internal/core/constants"
	datamodels2 "lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/parser/internal"
)

type nodeParser struct {}

func NewNodeParser() *nodeParser {
	return &nodeParser{}
}

func (n *nodeParser) ParseIntoNode(tokens []string) datamodels2.Node {
	if len(tokens) == 1 {
		return datamodels2.NewAtomNode(
			tokens[0],
		)
	}

	stack := datamodels2.NewNodeListStack()
	result := datamodels2.NewExpressionListNode(
		[]datamodels2.Node{},
	)

	for _, token := range tokens {
		if token == constants2.OPEN_TOKEN_STR {
			stack.Push([]datamodels2.Node{})
		} else if token == constants2.CLOSE_TOKEN_STR {
			node := datamodels2.NewAtomNode(constants2.NIL)
			top := stack.Pop()
			top = append(top, node)
			result = datamodels2.NewExpressionListNode(top)
			if stack.IsNotEmpty() {
				top = stack.Pop()
				top = append(top, result)
				stack.Push(top)
			}
		} else {
			node := datamodels2.NewAtomNode(token)
			top := stack.Pop()
			top = append(top, node)
			stack.Push(top)
		}
	}

	return result
}

var _ internal.NodeParser = &nodeParser{}