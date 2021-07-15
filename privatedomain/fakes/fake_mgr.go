// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	cfclient "github.com/cloudfoundry-community/go-cfclient"
	"github.com/vmwarepivotallabs/cf-mgmt/privatedomain"
)

type FakeManager struct {
	CreatePrivateDomainsStub        func() error
	createPrivateDomainsMutex       sync.RWMutex
	createPrivateDomainsArgsForCall []struct {
	}
	createPrivateDomainsReturns struct {
		result1 error
	}
	createPrivateDomainsReturnsOnCall map[int]struct {
		result1 error
	}
	ListOrgOwnedPrivateDomainsStub        func(string) (map[string]cfclient.Domain, error)
	listOrgOwnedPrivateDomainsMutex       sync.RWMutex
	listOrgOwnedPrivateDomainsArgsForCall []struct {
		arg1 string
	}
	listOrgOwnedPrivateDomainsReturns struct {
		result1 map[string]cfclient.Domain
		result2 error
	}
	listOrgOwnedPrivateDomainsReturnsOnCall map[int]struct {
		result1 map[string]cfclient.Domain
		result2 error
	}
	ListOrgSharedPrivateDomainsStub        func(string) (map[string]cfclient.Domain, error)
	listOrgSharedPrivateDomainsMutex       sync.RWMutex
	listOrgSharedPrivateDomainsArgsForCall []struct {
		arg1 string
	}
	listOrgSharedPrivateDomainsReturns struct {
		result1 map[string]cfclient.Domain
		result2 error
	}
	listOrgSharedPrivateDomainsReturnsOnCall map[int]struct {
		result1 map[string]cfclient.Domain
		result2 error
	}
	SharePrivateDomainsStub        func() error
	sharePrivateDomainsMutex       sync.RWMutex
	sharePrivateDomainsArgsForCall []struct {
	}
	sharePrivateDomainsReturns struct {
		result1 error
	}
	sharePrivateDomainsReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeManager) CreatePrivateDomains() error {
	fake.createPrivateDomainsMutex.Lock()
	ret, specificReturn := fake.createPrivateDomainsReturnsOnCall[len(fake.createPrivateDomainsArgsForCall)]
	fake.createPrivateDomainsArgsForCall = append(fake.createPrivateDomainsArgsForCall, struct {
	}{})
	stub := fake.CreatePrivateDomainsStub
	fakeReturns := fake.createPrivateDomainsReturns
	fake.recordInvocation("CreatePrivateDomains", []interface{}{})
	fake.createPrivateDomainsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeManager) CreatePrivateDomainsCallCount() int {
	fake.createPrivateDomainsMutex.RLock()
	defer fake.createPrivateDomainsMutex.RUnlock()
	return len(fake.createPrivateDomainsArgsForCall)
}

func (fake *FakeManager) CreatePrivateDomainsCalls(stub func() error) {
	fake.createPrivateDomainsMutex.Lock()
	defer fake.createPrivateDomainsMutex.Unlock()
	fake.CreatePrivateDomainsStub = stub
}

func (fake *FakeManager) CreatePrivateDomainsReturns(result1 error) {
	fake.createPrivateDomainsMutex.Lock()
	defer fake.createPrivateDomainsMutex.Unlock()
	fake.CreatePrivateDomainsStub = nil
	fake.createPrivateDomainsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) CreatePrivateDomainsReturnsOnCall(i int, result1 error) {
	fake.createPrivateDomainsMutex.Lock()
	defer fake.createPrivateDomainsMutex.Unlock()
	fake.CreatePrivateDomainsStub = nil
	if fake.createPrivateDomainsReturnsOnCall == nil {
		fake.createPrivateDomainsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createPrivateDomainsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) ListOrgOwnedPrivateDomains(arg1 string) (map[string]cfclient.Domain, error) {
	fake.listOrgOwnedPrivateDomainsMutex.Lock()
	ret, specificReturn := fake.listOrgOwnedPrivateDomainsReturnsOnCall[len(fake.listOrgOwnedPrivateDomainsArgsForCall)]
	fake.listOrgOwnedPrivateDomainsArgsForCall = append(fake.listOrgOwnedPrivateDomainsArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ListOrgOwnedPrivateDomainsStub
	fakeReturns := fake.listOrgOwnedPrivateDomainsReturns
	fake.recordInvocation("ListOrgOwnedPrivateDomains", []interface{}{arg1})
	fake.listOrgOwnedPrivateDomainsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeManager) ListOrgOwnedPrivateDomainsCallCount() int {
	fake.listOrgOwnedPrivateDomainsMutex.RLock()
	defer fake.listOrgOwnedPrivateDomainsMutex.RUnlock()
	return len(fake.listOrgOwnedPrivateDomainsArgsForCall)
}

func (fake *FakeManager) ListOrgOwnedPrivateDomainsCalls(stub func(string) (map[string]cfclient.Domain, error)) {
	fake.listOrgOwnedPrivateDomainsMutex.Lock()
	defer fake.listOrgOwnedPrivateDomainsMutex.Unlock()
	fake.ListOrgOwnedPrivateDomainsStub = stub
}

func (fake *FakeManager) ListOrgOwnedPrivateDomainsArgsForCall(i int) string {
	fake.listOrgOwnedPrivateDomainsMutex.RLock()
	defer fake.listOrgOwnedPrivateDomainsMutex.RUnlock()
	argsForCall := fake.listOrgOwnedPrivateDomainsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeManager) ListOrgOwnedPrivateDomainsReturns(result1 map[string]cfclient.Domain, result2 error) {
	fake.listOrgOwnedPrivateDomainsMutex.Lock()
	defer fake.listOrgOwnedPrivateDomainsMutex.Unlock()
	fake.ListOrgOwnedPrivateDomainsStub = nil
	fake.listOrgOwnedPrivateDomainsReturns = struct {
		result1 map[string]cfclient.Domain
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) ListOrgOwnedPrivateDomainsReturnsOnCall(i int, result1 map[string]cfclient.Domain, result2 error) {
	fake.listOrgOwnedPrivateDomainsMutex.Lock()
	defer fake.listOrgOwnedPrivateDomainsMutex.Unlock()
	fake.ListOrgOwnedPrivateDomainsStub = nil
	if fake.listOrgOwnedPrivateDomainsReturnsOnCall == nil {
		fake.listOrgOwnedPrivateDomainsReturnsOnCall = make(map[int]struct {
			result1 map[string]cfclient.Domain
			result2 error
		})
	}
	fake.listOrgOwnedPrivateDomainsReturnsOnCall[i] = struct {
		result1 map[string]cfclient.Domain
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) ListOrgSharedPrivateDomains(arg1 string) (map[string]cfclient.Domain, error) {
	fake.listOrgSharedPrivateDomainsMutex.Lock()
	ret, specificReturn := fake.listOrgSharedPrivateDomainsReturnsOnCall[len(fake.listOrgSharedPrivateDomainsArgsForCall)]
	fake.listOrgSharedPrivateDomainsArgsForCall = append(fake.listOrgSharedPrivateDomainsArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ListOrgSharedPrivateDomainsStub
	fakeReturns := fake.listOrgSharedPrivateDomainsReturns
	fake.recordInvocation("ListOrgSharedPrivateDomains", []interface{}{arg1})
	fake.listOrgSharedPrivateDomainsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeManager) ListOrgSharedPrivateDomainsCallCount() int {
	fake.listOrgSharedPrivateDomainsMutex.RLock()
	defer fake.listOrgSharedPrivateDomainsMutex.RUnlock()
	return len(fake.listOrgSharedPrivateDomainsArgsForCall)
}

func (fake *FakeManager) ListOrgSharedPrivateDomainsCalls(stub func(string) (map[string]cfclient.Domain, error)) {
	fake.listOrgSharedPrivateDomainsMutex.Lock()
	defer fake.listOrgSharedPrivateDomainsMutex.Unlock()
	fake.ListOrgSharedPrivateDomainsStub = stub
}

func (fake *FakeManager) ListOrgSharedPrivateDomainsArgsForCall(i int) string {
	fake.listOrgSharedPrivateDomainsMutex.RLock()
	defer fake.listOrgSharedPrivateDomainsMutex.RUnlock()
	argsForCall := fake.listOrgSharedPrivateDomainsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeManager) ListOrgSharedPrivateDomainsReturns(result1 map[string]cfclient.Domain, result2 error) {
	fake.listOrgSharedPrivateDomainsMutex.Lock()
	defer fake.listOrgSharedPrivateDomainsMutex.Unlock()
	fake.ListOrgSharedPrivateDomainsStub = nil
	fake.listOrgSharedPrivateDomainsReturns = struct {
		result1 map[string]cfclient.Domain
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) ListOrgSharedPrivateDomainsReturnsOnCall(i int, result1 map[string]cfclient.Domain, result2 error) {
	fake.listOrgSharedPrivateDomainsMutex.Lock()
	defer fake.listOrgSharedPrivateDomainsMutex.Unlock()
	fake.ListOrgSharedPrivateDomainsStub = nil
	if fake.listOrgSharedPrivateDomainsReturnsOnCall == nil {
		fake.listOrgSharedPrivateDomainsReturnsOnCall = make(map[int]struct {
			result1 map[string]cfclient.Domain
			result2 error
		})
	}
	fake.listOrgSharedPrivateDomainsReturnsOnCall[i] = struct {
		result1 map[string]cfclient.Domain
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) SharePrivateDomains() error {
	fake.sharePrivateDomainsMutex.Lock()
	ret, specificReturn := fake.sharePrivateDomainsReturnsOnCall[len(fake.sharePrivateDomainsArgsForCall)]
	fake.sharePrivateDomainsArgsForCall = append(fake.sharePrivateDomainsArgsForCall, struct {
	}{})
	stub := fake.SharePrivateDomainsStub
	fakeReturns := fake.sharePrivateDomainsReturns
	fake.recordInvocation("SharePrivateDomains", []interface{}{})
	fake.sharePrivateDomainsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeManager) SharePrivateDomainsCallCount() int {
	fake.sharePrivateDomainsMutex.RLock()
	defer fake.sharePrivateDomainsMutex.RUnlock()
	return len(fake.sharePrivateDomainsArgsForCall)
}

func (fake *FakeManager) SharePrivateDomainsCalls(stub func() error) {
	fake.sharePrivateDomainsMutex.Lock()
	defer fake.sharePrivateDomainsMutex.Unlock()
	fake.SharePrivateDomainsStub = stub
}

func (fake *FakeManager) SharePrivateDomainsReturns(result1 error) {
	fake.sharePrivateDomainsMutex.Lock()
	defer fake.sharePrivateDomainsMutex.Unlock()
	fake.SharePrivateDomainsStub = nil
	fake.sharePrivateDomainsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) SharePrivateDomainsReturnsOnCall(i int, result1 error) {
	fake.sharePrivateDomainsMutex.Lock()
	defer fake.sharePrivateDomainsMutex.Unlock()
	fake.SharePrivateDomainsStub = nil
	if fake.sharePrivateDomainsReturnsOnCall == nil {
		fake.sharePrivateDomainsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sharePrivateDomainsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createPrivateDomainsMutex.RLock()
	defer fake.createPrivateDomainsMutex.RUnlock()
	fake.listOrgOwnedPrivateDomainsMutex.RLock()
	defer fake.listOrgOwnedPrivateDomainsMutex.RUnlock()
	fake.listOrgSharedPrivateDomainsMutex.RLock()
	defer fake.listOrgSharedPrivateDomainsMutex.RUnlock()
	fake.sharePrivateDomainsMutex.RLock()
	defer fake.sharePrivateDomainsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeManager) recordInvocation(key string, args []interface{}) {
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

var _ privatedomain.Manager = new(FakeManager)
