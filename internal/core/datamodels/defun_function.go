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

func (d *DefunFunction) GetFunctionName() string {
	return d.name
}

func (d *DefunFunction) GetFormalParameters() []string {
	return d.formalParameters
}

func (d *DefunFunction) GetBody() Node {
	return d.body
}