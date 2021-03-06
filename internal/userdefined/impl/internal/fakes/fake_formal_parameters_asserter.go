// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"lisp_lambda-go/internal/userdefined/impl/internal"
	"sync"
)

type FakeFormalParametersAsserter struct {
	AssertFormalParametersStub        func([]string) error
	assertFormalParametersMutex       sync.RWMutex
	assertFormalParametersArgsForCall []struct {
		arg1 []string
	}
	assertFormalParametersReturns struct {
		result1 error
	}
	assertFormalParametersReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFormalParametersAsserter) AssertFormalParameters(arg1 []string) error {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.assertFormalParametersMutex.Lock()
	ret, specificReturn := fake.assertFormalParametersReturnsOnCall[len(fake.assertFormalParametersArgsForCall)]
	fake.assertFormalParametersArgsForCall = append(fake.assertFormalParametersArgsForCall, struct {
		arg1 []string
	}{arg1Copy})
	stub := fake.AssertFormalParametersStub
	fakeReturns := fake.assertFormalParametersReturns
	fake.recordInvocation("AssertFormalParameters", []interface{}{arg1Copy})
	fake.assertFormalParametersMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeFormalParametersAsserter) AssertFormalParametersCallCount() int {
	fake.assertFormalParametersMutex.RLock()
	defer fake.assertFormalParametersMutex.RUnlock()
	return len(fake.assertFormalParametersArgsForCall)
}

func (fake *FakeFormalParametersAsserter) AssertFormalParametersCalls(stub func([]string) error) {
	fake.assertFormalParametersMutex.Lock()
	defer fake.assertFormalParametersMutex.Unlock()
	fake.AssertFormalParametersStub = stub
}

func (fake *FakeFormalParametersAsserter) AssertFormalParametersArgsForCall(i int) []string {
	fake.assertFormalParametersMutex.RLock()
	defer fake.assertFormalParametersMutex.RUnlock()
	argsForCall := fake.assertFormalParametersArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeFormalParametersAsserter) AssertFormalParametersReturns(result1 error) {
	fake.assertFormalParametersMutex.Lock()
	defer fake.assertFormalParametersMutex.Unlock()
	fake.AssertFormalParametersStub = nil
	fake.assertFormalParametersReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFormalParametersAsserter) AssertFormalParametersReturnsOnCall(i int, result1 error) {
	fake.assertFormalParametersMutex.Lock()
	defer fake.assertFormalParametersMutex.Unlock()
	fake.AssertFormalParametersStub = nil
	if fake.assertFormalParametersReturnsOnCall == nil {
		fake.assertFormalParametersReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.assertFormalParametersReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeFormalParametersAsserter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.assertFormalParametersMutex.RLock()
	defer fake.assertFormalParametersMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeFormalParametersAsserter) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ internal.FormalParametersAsserter = new(FakeFormalParametersAsserter)
