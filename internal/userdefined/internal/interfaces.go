package internal

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o fakes/fake_formal_parameters_asserter.go . FormalParametersAsserter
type FormalParametersAsserter interface{
	AssertFormalParameters(formalParameters []string) error
}

//counterfeiter:generate -o fakes/fake_function_name_asserter.go . FunctionNameAsserter
type FunctionNameAsserter interface {
	AssertFunctionNameIsValid(functionName string) error
}