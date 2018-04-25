// Code generated by counterfeiter. DO NOT EDIT.
package canaryfakes

import (
	"replication-canary/canary"
	"replication-canary/models"
	"sync"
)

type FakeConnectionFactory struct {
	ConnsStub        func() ([]*models.NamedConnection, error)
	connsMutex       sync.RWMutex
	connsArgsForCall []struct{}
	connsReturns     struct {
		result1 []*models.NamedConnection
		result2 error
	}
	connsReturnsOnCall map[int]struct {
		result1 []*models.NamedConnection
		result2 error
	}
	WriteConnStub        func() (*models.NamedConnection, error)
	writeConnMutex       sync.RWMutex
	writeConnArgsForCall []struct{}
	writeConnReturns     struct {
		result1 *models.NamedConnection
		result2 error
	}
	writeConnReturnsOnCall map[int]struct {
		result1 *models.NamedConnection
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConnectionFactory) Conns() ([]*models.NamedConnection, error) {
	fake.connsMutex.Lock()
	ret, specificReturn := fake.connsReturnsOnCall[len(fake.connsArgsForCall)]
	fake.connsArgsForCall = append(fake.connsArgsForCall, struct{}{})
	fake.recordInvocation("Conns", []interface{}{})
	fake.connsMutex.Unlock()
	if fake.ConnsStub != nil {
		return fake.ConnsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.connsReturns.result1, fake.connsReturns.result2
}

func (fake *FakeConnectionFactory) ConnsCallCount() int {
	fake.connsMutex.RLock()
	defer fake.connsMutex.RUnlock()
	return len(fake.connsArgsForCall)
}

func (fake *FakeConnectionFactory) ConnsReturns(result1 []*models.NamedConnection, result2 error) {
	fake.ConnsStub = nil
	fake.connsReturns = struct {
		result1 []*models.NamedConnection
		result2 error
	}{result1, result2}
}

func (fake *FakeConnectionFactory) ConnsReturnsOnCall(i int, result1 []*models.NamedConnection, result2 error) {
	fake.ConnsStub = nil
	if fake.connsReturnsOnCall == nil {
		fake.connsReturnsOnCall = make(map[int]struct {
			result1 []*models.NamedConnection
			result2 error
		})
	}
	fake.connsReturnsOnCall[i] = struct {
		result1 []*models.NamedConnection
		result2 error
	}{result1, result2}
}

func (fake *FakeConnectionFactory) WriteConn() (*models.NamedConnection, error) {
	fake.writeConnMutex.Lock()
	ret, specificReturn := fake.writeConnReturnsOnCall[len(fake.writeConnArgsForCall)]
	fake.writeConnArgsForCall = append(fake.writeConnArgsForCall, struct{}{})
	fake.recordInvocation("WriteConn", []interface{}{})
	fake.writeConnMutex.Unlock()
	if fake.WriteConnStub != nil {
		return fake.WriteConnStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.writeConnReturns.result1, fake.writeConnReturns.result2
}

func (fake *FakeConnectionFactory) WriteConnCallCount() int {
	fake.writeConnMutex.RLock()
	defer fake.writeConnMutex.RUnlock()
	return len(fake.writeConnArgsForCall)
}

func (fake *FakeConnectionFactory) WriteConnReturns(result1 *models.NamedConnection, result2 error) {
	fake.WriteConnStub = nil
	fake.writeConnReturns = struct {
		result1 *models.NamedConnection
		result2 error
	}{result1, result2}
}

func (fake *FakeConnectionFactory) WriteConnReturnsOnCall(i int, result1 *models.NamedConnection, result2 error) {
	fake.WriteConnStub = nil
	if fake.writeConnReturnsOnCall == nil {
		fake.writeConnReturnsOnCall = make(map[int]struct {
			result1 *models.NamedConnection
			result2 error
		})
	}
	fake.writeConnReturnsOnCall[i] = struct {
		result1 *models.NamedConnection
		result2 error
	}{result1, result2}
}

func (fake *FakeConnectionFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.connsMutex.RLock()
	defer fake.connsMutex.RUnlock()
	fake.writeConnMutex.RLock()
	defer fake.writeConnMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeConnectionFactory) recordInvocation(key string, args []interface{}) {
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

var _ canary.ConnectionFactory = new(FakeConnectionFactory)