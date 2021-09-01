package printer

import (
	"lisp_lambda-go/internal"
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"strings"
)

type listNotationPrinter struct {}

func NewListNotationPrinter() *listNotationPrinter {
	return &listNotationPrinter{}
}

func (l listNotationPrinter) PrintAllInListNotation(nodes []datamodels.Node) string {
	var strs []string
	for _, node := range nodes {
		str := l.PrintInListNotation(node)
		strs = append(strs, str)
	}
	return strings.Join(strs, constants.NewLine)
}

func (l listNotationPrinter) PrintInListNotation(node datamodels.Node) string {
	expressionNode, isExpressionNode := node.(*datamodels.ExpressionListNode)
	if isExpressionNode {
		result := constants.OPEN_TOKEN_STR
		for i := 0; i < len(expressionNode.GetChildren()) - 1; i++ {
			result += l.PrintInListNotation(expressionNode.GetChildren()[i])
			result += constants.Space
		}
		result = strings.TrimSpace(result)
		lastChild := expressionNode.GetChildren()[len(expressionNode.GetChildren()) - 1]
		atomNode, isAtomNode := lastChild.(*datamodels.AtomNode)
		if isAtomNode && atomNode.GetValue() == constants.NIL {
			result += constants.CLOSE_TOKEN_STR
		} else {
			result += constants.ListDelimiter
			result += l.PrintInListNotation(lastChild)
			result += constants.CLOSE_TOKEN_STR
		}
		return result
	} else {
		atomNode, _ := node.(*datamodels.AtomNode)
		return atomNode.GetValue()
	}
}

var _ internal.ListNotationPrinter = &listNotationPrinter{}