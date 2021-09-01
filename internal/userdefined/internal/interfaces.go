package internal

type FormalParametersAsserter interface{
	AssertFormalParameters(formalParameters []string) error
}

type FunctionNameAsserter interface {
	AssertFunctionNameIsValid(functionName string) error
}