package impl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDuplicateFormalParameter(t *testing.T) {
	duplicate := "duplicate"
	formalParameters := []string{duplicate, duplicate}

	invalidName := "invalidName"
	invalidNameSet := map[string] bool {invalidName: true}
	userDefinedFormalParametersAsserter := NewUserDefinedFormalParametersAsserter(
		invalidNameSet,
	)

	actual := userDefinedFormalParametersAsserter.AssertFormalParameters(formalParameters)

	assert.NotNil(t, actual)
}

func TestInvalidName(t *testing.T) {
	invalidName := "invalidName"

	formalParameters := []string{invalidName}

	invalidNameSet := map[string] bool {invalidName: true}
	userDefinedFormalParametersAsserter := NewUserDefinedFormalParametersAsserter(
		invalidNameSet,
	)

	actual := userDefinedFormalParametersAsserter.AssertFormalParameters(formalParameters)

	assert.NotNil(t, actual)
}

func TestValidName(t *testing.T){
	validName := "validName"
	formalParameters := []string{validName}

	invalidName := "invalidName"
	invalidNameSet := map[string] bool {invalidName: true}
	userDefinedFormalParametersAsserter := NewUserDefinedFormalParametersAsserter(
		invalidNameSet,
	)

	actual := userDefinedFormalParametersAsserter.AssertFormalParameters(formalParameters)

	assert.Nil(t, actual)
}