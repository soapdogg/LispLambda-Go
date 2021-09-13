package impl

import (
	"errors"
	"lisp_lambda-go/internal/userdefined/impl/internal"
)

type userDefinedFormalParametersAsserter struct {
	invalidNames map[string]bool
}

func NewUserDefinedFormalParametersAsserter(
	invalidNames map[string] bool,
) *userDefinedFormalParametersAsserter {
	return &userDefinedFormalParametersAsserter{invalidNames}
}

func (u userDefinedFormalParametersAsserter) AssertFormalParameters(formalParameters []string) error {
	formalParametersSet := map[string] bool{}
	for _, formalParameter := range formalParameters {
		formalParametersSet[formalParameter] = true
	}
	if len(formalParameters) != len(formalParametersSet) {
		return errors.New("Error! Duplicate formal parameter name!\n")
	}
	for invalidName, _ := range u.invalidNames {
		_, ok := formalParametersSet[invalidName]
		if ok {
			return errors.New("Error! Invalid formal parameter name!\n")
		}
	}
	return nil
}


var _ internal.FormalParametersAsserter = &userDefinedFormalParametersAsserter{}