package parser

import (
	top "lisp_lambda-go/internal"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/parser/internal"
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


var _ top.Parser = &parser{}