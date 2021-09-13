package impl

import (
	"errors"
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	parser2 "lisp_lambda-go/internal/parser"
	"lisp_lambda-go/internal/parser/impl/internal"
)

type parser struct {
	nodeParser internal.NodeParser
}

func NewParser(
	nodeParser internal.NodeParser,
) *parser {
	return &parser{
		nodeParser,
	}
}

func (p *parser) Parse(tokens []string) ([]datamodels.Node, error) {
	openClose := 0
	var listOfLists [][]string

	currentList := []string{}
	for _, token := range tokens {
		currentList = append(currentList, token)
		if token == constants.OPEN_TOKEN_STR {
			openClose++
		} else if token == constants.CLOSE_TOKEN_STR {
			openClose--
		}
		if openClose < 0 {
			return nil, errors.New("Expected either an ATOM or OPEN token.   Actual Value: )")
		}
		if openClose == 0 {
			listOfLists = append(listOfLists, currentList)
			currentList = []string{}
		}
	}

	if openClose > 0 {
		return nil, errors.New("Expected a token. Actual: nil")
	}

	var nodes []datamodels.Node
	for _, list := range listOfLists {
		node := p.nodeParser.ParseIntoNode(list)
		nodes = append(nodes, node)
	}
	return nodes, nil
}


var _ parser2.Parser = &parser{}