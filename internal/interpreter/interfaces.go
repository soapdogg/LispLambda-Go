package interpreter

type Interpreter interface {
	Interpret(input string) (string, error)
}