// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"lisp_lambda-go/internal/asserter"
	"lisp_lambda-go/internal/core/datamodels"
	"sync"
)

type FakeExpressionListLengthAsserter struct {
	AssertLengthIsAsExpectedStub        func([]datamodels.Node, map[string]datamodels.DefunFunction) error
	assertLengthIsAsExpectedMutex       sync.RWMutex
	assertLengthIsAsExpectedArgsForCall []struct {
		arg1 []datamodels.Node
		arg2 map[string]datamodels.DefunFunction
	}
	assertLengthIsAsExpectedReturns struct {
		result1 error
	}
	assertLengthIsAsExpectedReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeExpressionListLengthAsserter) AssertLengthIsAsExpected(arg1 []datamodels.Node, arg2 map[string]datamodels.DefunFunction) error {
	var arg1Copy []datamodels.Node
	if arg1 != nil {
		arg1Copy = make([]datamodels.Node, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.assertLengthIsAsExpectedMutex.Lock()
	ret, specificReturn := fake.assertLengthIsAsExpectedReturnsOnCall[len(fake.assertLengthIsAsExpectedArgsForCall)]
	fake.assertLengthIsAsExpectedArgsForCall = append(fake.assertLengthIsAsExpectedArgsForCall, struct {
		arg1 []datamodels.Node
		arg2 map[string]datamodels.DefunFunction
	}{arg1Copy, arg2})
	stub := fake.AssertLengthIsAsExpectedStub
	fakeReturns := fake.assertLengthIsAsExpectedReturns
	fake.recordInvocation("AssertLengthIsAsExpected", []interface{}{arg1Copy, arg2})
	fake.assertLengthIsAsExpectedMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeExpressionListLengthAsserter) AssertLengthIsAsExpectedCallCount() int {
	fake.assertLengthIsAsExpectedMutex.RLock()
	defer fake.assertLengthIsAsExpectedMutex.RUnlock()
	return len(fake.assertLengthIsAsExpectedArgsForCall)
}

func (fake *FakeExpressionListLengthAsserter) AssertLengthIsAsExpectedCalls(stub func([]datamodels.Node, map[string]datamodels.DefunFunction) error) {
	fake.assertLengthIsAsExpectedMutex.Lock()
	defer fake.assertLengthIsAsExpectedMutex.Unlock()
	fake.AssertLengthIsAsExpectedStub = stub
}

func (fake *FakeExpressionListLengthAsserter) AssertLengthIsAsExpectedArgsForCall(i int) ([]datamodels.Node, map[string]datamodels.DefunFunction) {
	fake.assertLengthIsAsExpectedMutex.RLock()
	defer fake.assertLengthIsAsExpectedMutex.RUnlock()
	argsForCall := fake.assertLengthIsAsExpectedArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeExpressionListLengthAsserter) AssertLengthIsAsExpectedReturns(result1 error) {
	fake.assertLengthIsAsExpectedMutex.Lock()
	defer fake.assertLengthIsAsExpectedMutex.Unlock()
	fake.AssertLengthIsAsExpectedStub = nil
	fake.assertLengthIsAsExpectedReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeExpressionListLengthAsserter) AssertLengthIsAsExpectedReturnsOnCall(i int, result1 error) {
	fake.assertLengthIsAsExpectedMutex.Lock()
	defer fake.assertLengthIsAsExpectedMutex.Unlock()
	fake.AssertLengthIsAsExpectedStub = nil
	if fake.assertLengthIsAsExpectedReturnsOnCall == nil {
		fake.assertLengthIsAsExpectedReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.assertLengthIsAsExpectedReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeExpressionListLengthAsserter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.assertLengthIsAsExpectedMutex.RLock()
	defer fake.assertLengthIsAsExpectedMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeExpressionListLengthAsserter) recordInvocation(key string, args []interface{}) {
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

var _ asserter.ExpressionListLengthAsserter = new(FakeExpressionListLengthAsserter)
