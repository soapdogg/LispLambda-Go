// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"lisp_lambda-go/internal/core/datamodels"
	"sync"
)

type FakeAtomNode struct {
	GetValueStub        func() string
	getValueMutex       sync.RWMutex
	getValueArgsForCall []struct {
	}
	getValueReturns struct {
		result1 string
	}
	getValueReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAtomNode) GetValue() string {
	fake.getValueMutex.Lock()
	ret, specificReturn := fake.getValueReturnsOnCall[len(fake.getValueArgsForCall)]
	fake.getValueArgsForCall = append(fake.getValueArgsForCall, struct {
	}{})
	stub := fake.GetValueStub
	fakeReturns := fake.getValueReturns
	fake.recordInvocation("GetValue", []interface{}{})
	fake.getValueMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeAtomNode) GetValueCallCount() int {
	fake.getValueMutex.RLock()
	defer fake.getValueMutex.RUnlock()
	return len(fake.getValueArgsForCall)
}

func (fake *FakeAtomNode) GetValueCalls(stub func() string) {
	fake.getValueMutex.Lock()
	defer fake.getValueMutex.Unlock()
	fake.GetValueStub = stub
}

func (fake *FakeAtomNode) GetValueReturns(result1 string) {
	fake.getValueMutex.Lock()
	defer fake.getValueMutex.Unlock()
	fake.GetValueStub = nil
	fake.getValueReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeAtomNode) GetValueReturnsOnCall(i int, result1 string) {
	fake.getValueMutex.Lock()
	defer fake.getValueMutex.Unlock()
	fake.GetValueStub = nil
	if fake.getValueReturnsOnCall == nil {
		fake.getValueReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getValueReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeAtomNode) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getValueMutex.RLock()
	defer fake.getValueMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAtomNode) recordInvocation(key string, args []interface{}) {
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

var _ datamodels.AtomNode = new(FakeAtomNode)
