// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/evaluator/internal"
	"sync"
)

type FakeUnfinishedProgramStackItemEvaluator struct {
	EvaluateUnfinishedProgramItemStub        func(datamodels.ProgramItem, datamodels.NodeStack, datamodels.ProgramItemStack)
	evaluateUnfinishedProgramItemMutex       sync.RWMutex
	evaluateUnfinishedProgramItemArgsForCall []struct {
		arg1 datamodels.ProgramItem
		arg2 datamodels.NodeStack
		arg3 datamodels.ProgramItemStack
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUnfinishedProgramStackItemEvaluator) EvaluateUnfinishedProgramItem(arg1 datamodels.ProgramItem, arg2 datamodels.NodeStack, arg3 datamodels.ProgramItemStack) {
	fake.evaluateUnfinishedProgramItemMutex.Lock()
	fake.evaluateUnfinishedProgramItemArgsForCall = append(fake.evaluateUnfinishedProgramItemArgsForCall, struct {
		arg1 datamodels.ProgramItem
		arg2 datamodels.NodeStack
		arg3 datamodels.ProgramItemStack
	}{arg1, arg2, arg3})
	stub := fake.EvaluateUnfinishedProgramItemStub
	fake.recordInvocation("EvaluateUnfinishedProgramItem", []interface{}{arg1, arg2, arg3})
	fake.evaluateUnfinishedProgramItemMutex.Unlock()
	if stub != nil {
		fake.EvaluateUnfinishedProgramItemStub(arg1, arg2, arg3)
	}
}

func (fake *FakeUnfinishedProgramStackItemEvaluator) EvaluateUnfinishedProgramItemCallCount() int {
	fake.evaluateUnfinishedProgramItemMutex.RLock()
	defer fake.evaluateUnfinishedProgramItemMutex.RUnlock()
	return len(fake.evaluateUnfinishedProgramItemArgsForCall)
}

func (fake *FakeUnfinishedProgramStackItemEvaluator) EvaluateUnfinishedProgramItemCalls(stub func(datamodels.ProgramItem, datamodels.NodeStack, datamodels.ProgramItemStack)) {
	fake.evaluateUnfinishedProgramItemMutex.Lock()
	defer fake.evaluateUnfinishedProgramItemMutex.Unlock()
	fake.EvaluateUnfinishedProgramItemStub = stub
}

func (fake *FakeUnfinishedProgramStackItemEvaluator) EvaluateUnfinishedProgramItemArgsForCall(i int) (datamodels.ProgramItem, datamodels.NodeStack, datamodels.ProgramItemStack) {
	fake.evaluateUnfinishedProgramItemMutex.RLock()
	defer fake.evaluateUnfinishedProgramItemMutex.RUnlock()
	argsForCall := fake.evaluateUnfinishedProgramItemArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeUnfinishedProgramStackItemEvaluator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.evaluateUnfinishedProgramItemMutex.RLock()
	defer fake.evaluateUnfinishedProgramItemMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUnfinishedProgramStackItemEvaluator) recordInvocation(key string, args []interface{}) {
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

var _ internal.UnfinishedProgramStackItemEvaluator = new(FakeUnfinishedProgramStackItemEvaluator)
