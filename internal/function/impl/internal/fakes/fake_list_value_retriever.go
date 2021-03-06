// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/function/impl/internal"
	"sync"
)

type FakeListValueRetriever struct {
	RetrieveListValueStub        func(datamodels.Node, string) (datamodels.ExpressionListNode, error)
	retrieveListValueMutex       sync.RWMutex
	retrieveListValueArgsForCall []struct {
		arg1 datamodels.Node
		arg2 string
	}
	retrieveListValueReturns struct {
		result1 datamodels.ExpressionListNode
		result2 error
	}
	retrieveListValueReturnsOnCall map[int]struct {
		result1 datamodels.ExpressionListNode
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeListValueRetriever) RetrieveListValue(arg1 datamodels.Node, arg2 string) (datamodels.ExpressionListNode, error) {
	fake.retrieveListValueMutex.Lock()
	ret, specificReturn := fake.retrieveListValueReturnsOnCall[len(fake.retrieveListValueArgsForCall)]
	fake.retrieveListValueArgsForCall = append(fake.retrieveListValueArgsForCall, struct {
		arg1 datamodels.Node
		arg2 string
	}{arg1, arg2})
	stub := fake.RetrieveListValueStub
	fakeReturns := fake.retrieveListValueReturns
	fake.recordInvocation("RetrieveListValue", []interface{}{arg1, arg2})
	fake.retrieveListValueMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeListValueRetriever) RetrieveListValueCallCount() int {
	fake.retrieveListValueMutex.RLock()
	defer fake.retrieveListValueMutex.RUnlock()
	return len(fake.retrieveListValueArgsForCall)
}

func (fake *FakeListValueRetriever) RetrieveListValueCalls(stub func(datamodels.Node, string) (datamodels.ExpressionListNode, error)) {
	fake.retrieveListValueMutex.Lock()
	defer fake.retrieveListValueMutex.Unlock()
	fake.RetrieveListValueStub = stub
}

func (fake *FakeListValueRetriever) RetrieveListValueArgsForCall(i int) (datamodels.Node, string) {
	fake.retrieveListValueMutex.RLock()
	defer fake.retrieveListValueMutex.RUnlock()
	argsForCall := fake.retrieveListValueArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeListValueRetriever) RetrieveListValueReturns(result1 datamodels.ExpressionListNode, result2 error) {
	fake.retrieveListValueMutex.Lock()
	defer fake.retrieveListValueMutex.Unlock()
	fake.RetrieveListValueStub = nil
	fake.retrieveListValueReturns = struct {
		result1 datamodels.ExpressionListNode
		result2 error
	}{result1, result2}
}

func (fake *FakeListValueRetriever) RetrieveListValueReturnsOnCall(i int, result1 datamodels.ExpressionListNode, result2 error) {
	fake.retrieveListValueMutex.Lock()
	defer fake.retrieveListValueMutex.Unlock()
	fake.RetrieveListValueStub = nil
	if fake.retrieveListValueReturnsOnCall == nil {
		fake.retrieveListValueReturnsOnCall = make(map[int]struct {
			result1 datamodels.ExpressionListNode
			result2 error
		})
	}
	fake.retrieveListValueReturnsOnCall[i] = struct {
		result1 datamodels.ExpressionListNode
		result2 error
	}{result1, result2}
}

func (fake *FakeListValueRetriever) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.retrieveListValueMutex.RLock()
	defer fake.retrieveListValueMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeListValueRetriever) recordInvocation(key string, args []interface{}) {
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

var _ internal.ListValueRetriever = new(FakeListValueRetriever)
