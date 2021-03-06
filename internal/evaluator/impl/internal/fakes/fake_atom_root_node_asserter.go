// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/evaluator/impl/internal"
	"sync"
)

type FakeAtomRootNodeAsserter struct {
	AssertAtomRootNodeStub        func(datamodels.AtomNode) error
	assertAtomRootNodeMutex       sync.RWMutex
	assertAtomRootNodeArgsForCall []struct {
		arg1 datamodels.AtomNode
	}
	assertAtomRootNodeReturns struct {
		result1 error
	}
	assertAtomRootNodeReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAtomRootNodeAsserter) AssertAtomRootNode(arg1 datamodels.AtomNode) error {
	fake.assertAtomRootNodeMutex.Lock()
	ret, specificReturn := fake.assertAtomRootNodeReturnsOnCall[len(fake.assertAtomRootNodeArgsForCall)]
	fake.assertAtomRootNodeArgsForCall = append(fake.assertAtomRootNodeArgsForCall, struct {
		arg1 datamodels.AtomNode
	}{arg1})
	stub := fake.AssertAtomRootNodeStub
	fakeReturns := fake.assertAtomRootNodeReturns
	fake.recordInvocation("AssertAtomRootNode", []interface{}{arg1})
	fake.assertAtomRootNodeMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeAtomRootNodeAsserter) AssertAtomRootNodeCallCount() int {
	fake.assertAtomRootNodeMutex.RLock()
	defer fake.assertAtomRootNodeMutex.RUnlock()
	return len(fake.assertAtomRootNodeArgsForCall)
}

func (fake *FakeAtomRootNodeAsserter) AssertAtomRootNodeCalls(stub func(datamodels.AtomNode) error) {
	fake.assertAtomRootNodeMutex.Lock()
	defer fake.assertAtomRootNodeMutex.Unlock()
	fake.AssertAtomRootNodeStub = stub
}

func (fake *FakeAtomRootNodeAsserter) AssertAtomRootNodeArgsForCall(i int) datamodels.AtomNode {
	fake.assertAtomRootNodeMutex.RLock()
	defer fake.assertAtomRootNodeMutex.RUnlock()
	argsForCall := fake.assertAtomRootNodeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAtomRootNodeAsserter) AssertAtomRootNodeReturns(result1 error) {
	fake.assertAtomRootNodeMutex.Lock()
	defer fake.assertAtomRootNodeMutex.Unlock()
	fake.AssertAtomRootNodeStub = nil
	fake.assertAtomRootNodeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAtomRootNodeAsserter) AssertAtomRootNodeReturnsOnCall(i int, result1 error) {
	fake.assertAtomRootNodeMutex.Lock()
	defer fake.assertAtomRootNodeMutex.Unlock()
	fake.AssertAtomRootNodeStub = nil
	if fake.assertAtomRootNodeReturnsOnCall == nil {
		fake.assertAtomRootNodeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.assertAtomRootNodeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAtomRootNodeAsserter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.assertAtomRootNodeMutex.RLock()
	defer fake.assertAtomRootNodeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAtomRootNodeAsserter) recordInvocation(key string, args []interface{}) {
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

var _ internal.AtomRootNodeAsserter = new(FakeAtomRootNodeAsserter)
