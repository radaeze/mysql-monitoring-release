// Code generated by counterfeiter. DO NOT EDIT.
package canaryfakes

import (
	"replication-canary/canary"
	"replication-canary/models"
	"sync"
	"time"
)

type FakeChirper struct {
	ChirpStub        func(conns []*models.NamedConnection, writeConn *models.NamedConnection, timestamp time.Time) (bool, error)
	chirpMutex       sync.RWMutex
	chirpArgsForCall []struct {
		conns     []*models.NamedConnection
		writeConn *models.NamedConnection
		timestamp time.Time
	}
	chirpReturns struct {
		result1 bool
		result2 error
	}
	chirpReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeChirper) Chirp(conns []*models.NamedConnection, writeConn *models.NamedConnection, timestamp time.Time) (bool, error) {
	var connsCopy []*models.NamedConnection
	if conns != nil {
		connsCopy = make([]*models.NamedConnection, len(conns))
		copy(connsCopy, conns)
	}
	fake.chirpMutex.Lock()
	ret, specificReturn := fake.chirpReturnsOnCall[len(fake.chirpArgsForCall)]
	fake.chirpArgsForCall = append(fake.chirpArgsForCall, struct {
		conns     []*models.NamedConnection
		writeConn *models.NamedConnection
		timestamp time.Time
	}{connsCopy, writeConn, timestamp})
	fake.recordInvocation("Chirp", []interface{}{connsCopy, writeConn, timestamp})
	fake.chirpMutex.Unlock()
	if fake.ChirpStub != nil {
		return fake.ChirpStub(conns, writeConn, timestamp)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.chirpReturns.result1, fake.chirpReturns.result2
}

func (fake *FakeChirper) ChirpCallCount() int {
	fake.chirpMutex.RLock()
	defer fake.chirpMutex.RUnlock()
	return len(fake.chirpArgsForCall)
}

func (fake *FakeChirper) ChirpArgsForCall(i int) ([]*models.NamedConnection, *models.NamedConnection, time.Time) {
	fake.chirpMutex.RLock()
	defer fake.chirpMutex.RUnlock()
	return fake.chirpArgsForCall[i].conns, fake.chirpArgsForCall[i].writeConn, fake.chirpArgsForCall[i].timestamp
}

func (fake *FakeChirper) ChirpReturns(result1 bool, result2 error) {
	fake.ChirpStub = nil
	fake.chirpReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeChirper) ChirpReturnsOnCall(i int, result1 bool, result2 error) {
	fake.ChirpStub = nil
	if fake.chirpReturnsOnCall == nil {
		fake.chirpReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.chirpReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeChirper) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.chirpMutex.RLock()
	defer fake.chirpMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeChirper) recordInvocation(key string, args []interface{}) {
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

var _ canary.Chirper = new(FakeChirper)
