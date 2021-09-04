package datamodels

type DefunFunction struct {
	name string
	formalParameters []string
	body Node
}

func NewDefunFunction(
	name string,
	formalParameters []string,
	body Node,
) *DefunFunction {
	return &DefunFunction{
		name,
		formalParameters,
		body,
	}
}