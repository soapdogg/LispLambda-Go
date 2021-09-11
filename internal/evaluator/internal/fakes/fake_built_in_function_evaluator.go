// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/evaluator/internal"
	"sync"
)

type FakeBuiltInFunctionEvaluator struct {
	EvaluateBuiltInFunctionStub        func(string, datamodels.NodeStack, datamodels.ProgramItem, datamodels.NodeStack, datamodels.ProgramItemStack) error
	evaluateBuiltInFunctionMutex       sync.RWMutex
	evaluateBuiltInFunctionArgsForCall []struct {
		arg1 string
		arg2 datamodels.NodeStack
		arg3 datamodels.ProgramItem
		arg4 datamodels.NodeStack
		arg5 datamodels.ProgramItemStack
	}
	evaluateBuiltInFunctionReturns struct {
		result1 error
	}
	evaluateBuiltInFunctionReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBuiltInFunctionEvaluator) EvaluateBuiltInFunction(arg1 string, arg2 datamodels.NodeStack, arg3 datamodels.ProgramItem, arg4 datamodels.NodeStack, arg5 datamodels.ProgramItemStack) error {
	fake.evaluateBuiltInFunctionMutex.Lock()
	ret, specificReturn := fake.evaluateBuiltInFunctionReturnsOnCall[len(fake.evaluateBuiltInFunctionArgsForCall)]
	fake.evaluateBuiltInFunctionArgsForCall = append(fake.evaluateBuiltInFunctionArgsForCall, struct {
		arg1 string
		arg2 datamodels.NodeStack
		arg3 datamodels.ProgramItem
		arg4 datamodels.NodeStack
		arg5 datamodels.ProgramItemStack
	}{arg1, arg2, arg3, arg4, arg5})
	stub := fake.EvaluateBuiltInFunctionStub
	fakeReturns := fake.evaluateBuiltInFunctionReturns
	fake.recordInvocation("EvaluateBuiltInFunction", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.evaluateBuiltInFunctionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBuiltInFunctionEvaluator) EvaluateBuiltInFunctionCallCount() int {
	fake.evaluateBuiltInFunctionMutex.RLock()
	defer fake.evaluateBuiltInFunctionMutex.RUnlock()
	return len(fake.evaluateBuiltInFunctionArgsForCall)
}

func (fake *FakeBuiltInFunctionEvaluator) EvaluateBuiltInFunctionCalls(stub func(string, datamodels.NodeStack, datamodels.ProgramItem, datamodels.NodeStack, datamodels.ProgramItemStack) error) {
	fake.evaluateBuiltInFunctionMutex.Lock()
	defer fake.evaluateBuiltInFunctionMutex.Unlock()
	fake.EvaluateBuiltInFunctionStub = stub
}

func (fake *FakeBuiltInFunctionEvaluator) EvaluateBuiltInFunctionArgsForCall(i int) (string, datamodels.NodeStack, datamodels.ProgramItem, datamodels.NodeStack, datamodels.ProgramItemStack) {
	fake.evaluateBuiltInFunctionMutex.RLock()
	defer fake.evaluateBuiltInFunctionMutex.RUnlock()
	argsForCall := fake.evaluateBuiltInFunctionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeBuiltInFunctionEvaluator) EvaluateBuiltInFunctionReturns(result1 error) {
	fake.evaluateBuiltInFunctionMutex.Lock()
	defer fake.evaluateBuiltInFunctionMutex.Unlock()
	fake.EvaluateBuiltInFunctionStub = nil
	fake.evaluateBuiltInFunctionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuiltInFunctionEvaluator) EvaluateBuiltInFunctionReturnsOnCall(i int, result1 error) {
	fake.evaluateBuiltInFunctionMutex.Lock()
	defer fake.evaluateBuiltInFunctionMutex.Unlock()
	fake.EvaluateBuiltInFunctionStub = nil
	if fake.evaluateBuiltInFunctionReturnsOnCall == nil {
		fake.evaluateBuiltInFunctionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.evaluateBuiltInFunctionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuiltInFunctionEvaluator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.evaluateBuiltInFunctionMutex.RLock()
	defer fake.evaluateBuiltInFunctionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBuiltInFunctionEvaluator) recordInvocation(key string, args []interface{}) {
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

var _ internal.BuiltInFunctionEvaluator = new(FakeBuiltInFunctionEvaluator)
