package internal

type InvalidNameDeterminer interface {
	IsInvalidName(value string) bool
}

type UserDefinedFormalParametersAsserter interface{
	AssertFormalParameters(formalParameters []string) error
}

type UserDefinedFunctionNameAsserter interface {
	AssertFunctionNameIsValid(functionName string) error
}