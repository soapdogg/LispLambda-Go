package impl

import "lisp_lambda-go/internal/printer"

type singleton struct {
	listNotationPrinter printer.ListNotationPrinter
}

func NewSingleton() *singleton {
	listNotationPrinter := NewListNotationPrinter()

	return &singleton{
		listNotationPrinter,
	}
}

func (s *singleton) GetPrinter() printer.ListNotationPrinter {
	return s.listNotationPrinter
}
