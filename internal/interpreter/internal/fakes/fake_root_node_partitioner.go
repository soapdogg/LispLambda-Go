// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/interpreter/internal"
	"sync"
)

type FakeRootNodePartitioner struct {
	PartitionRootNodesStub        func([]datamodels.Node) datamodels.PartitionedRootNodes
	partitionRootNodesMutex       sync.RWMutex
	partitionRootNodesArgsForCall []struct {
		arg1 []datamodels.Node
	}
	partitionRootNodesReturns struct {
		result1 datamodels.PartitionedRootNodes
	}
	partitionRootNodesReturnsOnCall map[int]struct {
		result1 datamodels.PartitionedRootNodes
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRootNodePartitioner) PartitionRootNodes(arg1 []datamodels.Node) datamodels.PartitionedRootNodes {
	var arg1Copy []datamodels.Node
	if arg1 != nil {
		arg1Copy = make([]datamodels.Node, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.partitionRootNodesMutex.Lock()
	ret, specificReturn := fake.partitionRootNodesReturnsOnCall[len(fake.partitionRootNodesArgsForCall)]
	fake.partitionRootNodesArgsForCall = append(fake.partitionRootNodesArgsForCall, struct {
		arg1 []datamodels.Node
	}{arg1Copy})
	stub := fake.PartitionRootNodesStub
	fakeReturns := fake.partitionRootNodesReturns
	fake.recordInvocation("PartitionRootNodes", []interface{}{arg1Copy})
	fake.partitionRootNodesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRootNodePartitioner) PartitionRootNodesCallCount() int {
	fake.partitionRootNodesMutex.RLock()
	defer fake.partitionRootNodesMutex.RUnlock()
	return len(fake.partitionRootNodesArgsForCall)
}

func (fake *FakeRootNodePartitioner) PartitionRootNodesCalls(stub func([]datamodels.Node) datamodels.PartitionedRootNodes) {
	fake.partitionRootNodesMutex.Lock()
	defer fake.partitionRootNodesMutex.Unlock()
	fake.PartitionRootNodesStub = stub
}

func (fake *FakeRootNodePartitioner) PartitionRootNodesArgsForCall(i int) []datamodels.Node {
	fake.partitionRootNodesMutex.RLock()
	defer fake.partitionRootNodesMutex.RUnlock()
	argsForCall := fake.partitionRootNodesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRootNodePartitioner) PartitionRootNodesReturns(result1 datamodels.PartitionedRootNodes) {
	fake.partitionRootNodesMutex.Lock()
	defer fake.partitionRootNodesMutex.Unlock()
	fake.PartitionRootNodesStub = nil
	fake.partitionRootNodesReturns = struct {
		result1 datamodels.PartitionedRootNodes
	}{result1}
}

func (fake *FakeRootNodePartitioner) PartitionRootNodesReturnsOnCall(i int, result1 datamodels.PartitionedRootNodes) {
	fake.partitionRootNodesMutex.Lock()
	defer fake.partitionRootNodesMutex.Unlock()
	fake.PartitionRootNodesStub = nil
	if fake.partitionRootNodesReturnsOnCall == nil {
		fake.partitionRootNodesReturnsOnCall = make(map[int]struct {
			result1 datamodels.PartitionedRootNodes
		})
	}
	fake.partitionRootNodesReturnsOnCall[i] = struct {
		result1 datamodels.PartitionedRootNodes
	}{result1}
}

func (fake *FakeRootNodePartitioner) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.partitionRootNodesMutex.RLock()
	defer fake.partitionRootNodesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRootNodePartitioner) recordInvocation(key string, args []interface{}) {
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

var _ internal.RootNodePartitioner = new(FakeRootNodePartitioner)