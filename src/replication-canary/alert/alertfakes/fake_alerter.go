// Code generated by counterfeiter. DO NOT EDIT.
package alertfakes

import (
	"replication-canary/alert"
	"sync"
	"time"
)

type FakeAlerter struct {
	NotUnhealthyStub        func(time.Time) error
	notUnhealthyMutex       sync.RWMutex
	notUnhealthyArgsForCall []struct {
		arg1 time.Time
	}
	notUnhealthyReturns struct {
		result1 error
	}
	notUnhealthyReturnsOnCall map[int]struct {
		result1 error
	}
	UnhealthyStub        func(time.Time) error
	unhealthyMutex       sync.RWMutex
	unhealthyArgsForCall []struct {
		arg1 time.Time
	}
	unhealthyReturns struct {
		result1 error
	}
	unhealthyReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAlerter) NotUnhealthy(arg1 time.Time) error {
	fake.notUnhealthyMutex.Lock()
	ret, specificReturn := fake.notUnhealthyReturnsOnCall[len(fake.notUnhealthyArgsForCall)]
	fake.notUnhealthyArgsForCall = append(fake.notUnhealthyArgsForCall, struct {
		arg1 time.Time
	}{arg1})
	fake.recordInvocation("NotUnhealthy", []interface{}{arg1})
	fake.notUnhealthyMutex.Unlock()
	if fake.NotUnhealthyStub != nil {
		return fake.NotUnhealthyStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.notUnhealthyReturns.result1
}

func (fake *FakeAlerter) NotUnhealthyCallCount() int {
	fake.notUnhealthyMutex.RLock()
	defer fake.notUnhealthyMutex.RUnlock()
	return len(fake.notUnhealthyArgsForCall)
}

func (fake *FakeAlerter) NotUnhealthyArgsForCall(i int) time.Time {
	fake.notUnhealthyMutex.RLock()
	defer fake.notUnhealthyMutex.RUnlock()
	return fake.notUnhealthyArgsForCall[i].arg1
}

func (fake *FakeAlerter) NotUnhealthyReturns(result1 error) {
	fake.NotUnhealthyStub = nil
	fake.notUnhealthyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAlerter) NotUnhealthyReturnsOnCall(i int, result1 error) {
	fake.NotUnhealthyStub = nil
	if fake.notUnhealthyReturnsOnCall == nil {
		fake.notUnhealthyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.notUnhealthyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAlerter) Unhealthy(arg1 time.Time) error {
	fake.unhealthyMutex.Lock()
	ret, specificReturn := fake.unhealthyReturnsOnCall[len(fake.unhealthyArgsForCall)]
	fake.unhealthyArgsForCall = append(fake.unhealthyArgsForCall, struct {
		arg1 time.Time
	}{arg1})
	fake.recordInvocation("Unhealthy", []interface{}{arg1})
	fake.unhealthyMutex.Unlock()
	if fake.UnhealthyStub != nil {
		return fake.UnhealthyStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.unhealthyReturns.result1
}

func (fake *FakeAlerter) UnhealthyCallCount() int {
	fake.unhealthyMutex.RLock()
	defer fake.unhealthyMutex.RUnlock()
	return len(fake.unhealthyArgsForCall)
}

func (fake *FakeAlerter) UnhealthyArgsForCall(i int) time.Time {
	fake.unhealthyMutex.RLock()
	defer fake.unhealthyMutex.RUnlock()
	return fake.unhealthyArgsForCall[i].arg1
}

func (fake *FakeAlerter) UnhealthyReturns(result1 error) {
	fake.UnhealthyStub = nil
	fake.unhealthyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAlerter) UnhealthyReturnsOnCall(i int, result1 error) {
	fake.UnhealthyStub = nil
	if fake.unhealthyReturnsOnCall == nil {
		fake.unhealthyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.unhealthyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAlerter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.notUnhealthyMutex.RLock()
	defer fake.notUnhealthyMutex.RUnlock()
	fake.unhealthyMutex.RLock()
	defer fake.unhealthyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAlerter) recordInvocation(key string, args []interface{}) {
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

var _ alert.Alerter = new(FakeAlerter)
