package impl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValueIsKeyword(t *testing.T) {
	keyword := "keyword"
	keywords := map[string]bool {keyword: true}
	functionNameAsserter := NewFunctionNameAsserter(keywords)

	actual := functionNameAsserter.AssertFunctionNameIsValid(keyword)
	assert.NotNil(t, actual)
}

func TestValueIsNumeric(t *testing.T) {
	keyword := "keyword"
	keywords := map[string]bool {keyword: true}
	functionNameAsserter := NewFunctionNameAsserter(keywords)

	actual := functionNameAsserter.AssertFunctionNameIsValid("1234")
	assert.NotNil(t, actual)
}

func TestValueIsValid(t *testing.T) {
	keyword := "keyword"
	keywords := map[string]bool {keyword: true}
	functionNameAsserter := NewFunctionNameAsserter(keywords)

	actual := functionNameAsserter.AssertFunctionNameIsValid("different")
	assert.Nil(t, actual)
}
