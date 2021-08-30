package parser

import (
	"lisp_lambda-go/lambda/core/datamodels"
	"lisp_lambda-go/lambda/parser/internal"
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

func (p *parser) Parse(tokens []string) []datamodels.Node {
	panic("implement me")
}


var _ Parser = &parser{}