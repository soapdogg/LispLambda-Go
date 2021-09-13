package impl

import parser2 "lisp_lambda-go/internal/parser"

type singleton struct {
	parser parser2.Parser
}

func NewSingleton() *singleton {
	nodeParser := NewNodeParser()
	parser := NewParser(
		nodeParser,
	)

	return &singleton{
		parser,
	}
}

func (s *singleton) GetParser() parser2.Parser {
	return s.parser
}
