package datamodels

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o fakes/fake_defun_function.go . DefunFunction
type DefunFunction interface {
	GetFunctionName() string
	GetFormalParameters() []string
	GetBody() Node
}

type defunFunction struct {
	name string
	formalParameters []string
	body Node
}

func NewDefunFunction(
	name string,
	formalParameters []string,
	body Node,
) *defunFunction {
	return &defunFunction{
		name,
		formalParameters,
		body,
	}
}

func (d *defunFunction) GetFunctionName() string {
	return d.name
}

func (d *defunFunction) GetFormalParameters() []string {
	return d.formalParameters
}

func (d *defunFunction) GetBody() Node {
	return d.body
}

var _ DefunFunction = &defunFunction{}