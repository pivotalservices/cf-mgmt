// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	cfclient "github.com/cloudfoundry-community/go-cfclient"
	"github.com/pivotalservices/cf-mgmt/space"
)

type FakeManager struct {
	CreateSpacesStub        func() error
	createSpacesMutex       sync.RWMutex
	createSpacesArgsForCall []struct {
	}
	createSpacesReturns struct {
		result1 error
	}
	createSpacesReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteSpacesStub        func() error
	deleteSpacesMutex       sync.RWMutex
	deleteSpacesArgsForCall []struct {
	}
	deleteSpacesReturns struct {
		result1 error
	}
	deleteSpacesReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteSpacesForOrgStub        func(string, string) error
	deleteSpacesForOrgMutex       sync.RWMutex
	deleteSpacesForOrgArgsForCall []struct {
		arg1 string
		arg2 string
	}
	deleteSpacesForOrgReturns struct {
		result1 error
	}
	deleteSpacesForOrgReturnsOnCall map[int]struct {
		result1 error
	}
	FindSpaceStub        func(string, string) (cfclient.Space, error)
	findSpaceMutex       sync.RWMutex
	findSpaceArgsForCall []struct {
		arg1 string
		arg2 string
	}
	findSpaceReturns struct {
		result1 cfclient.Space
		result2 error
	}
	findSpaceReturnsOnCall map[int]struct {
		result1 cfclient.Space
		result2 error
	}
	ListSpacesStub        func(string) ([]cfclient.Space, error)
	listSpacesMutex       sync.RWMutex
	listSpacesArgsForCall []struct {
		arg1 string
	}
	listSpacesReturns struct {
		result1 []cfclient.Space
		result2 error
	}
	listSpacesReturnsOnCall map[int]struct {
		result1 []cfclient.Space
		result2 error
	}
	UpdateSpacesStub        func() error
	updateSpacesMutex       sync.RWMutex
	updateSpacesArgsForCall []struct {
	}
	updateSpacesReturns struct {
		result1 error
	}
	updateSpacesReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateSpacesMetadataStub        func() error
	updateSpacesMetadataMutex       sync.RWMutex
	updateSpacesMetadataArgsForCall []struct {
	}
	updateSpacesMetadataReturns struct {
		result1 error
	}
	updateSpacesMetadataReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeManager) CreateSpaces() error {
	fake.createSpacesMutex.Lock()
	ret, specificReturn := fake.createSpacesReturnsOnCall[len(fake.createSpacesArgsForCall)]
	fake.createSpacesArgsForCall = append(fake.createSpacesArgsForCall, struct {
	}{})
	fake.recordInvocation("CreateSpaces", []interface{}{})
	fake.createSpacesMutex.Unlock()
	if fake.CreateSpacesStub != nil {
		return fake.CreateSpacesStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createSpacesReturns
	return fakeReturns.result1
}

func (fake *FakeManager) CreateSpacesCallCount() int {
	fake.createSpacesMutex.RLock()
	defer fake.createSpacesMutex.RUnlock()
	return len(fake.createSpacesArgsForCall)
}

func (fake *FakeManager) CreateSpacesCalls(stub func() error) {
	fake.createSpacesMutex.Lock()
	defer fake.createSpacesMutex.Unlock()
	fake.CreateSpacesStub = stub
}

func (fake *FakeManager) CreateSpacesReturns(result1 error) {
	fake.createSpacesMutex.Lock()
	defer fake.createSpacesMutex.Unlock()
	fake.CreateSpacesStub = nil
	fake.createSpacesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) CreateSpacesReturnsOnCall(i int, result1 error) {
	fake.createSpacesMutex.Lock()
	defer fake.createSpacesMutex.Unlock()
	fake.CreateSpacesStub = nil
	if fake.createSpacesReturnsOnCall == nil {
		fake.createSpacesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createSpacesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) DeleteSpaces() error {
	fake.deleteSpacesMutex.Lock()
	ret, specificReturn := fake.deleteSpacesReturnsOnCall[len(fake.deleteSpacesArgsForCall)]
	fake.deleteSpacesArgsForCall = append(fake.deleteSpacesArgsForCall, struct {
	}{})
	fake.recordInvocation("DeleteSpaces", []interface{}{})
	fake.deleteSpacesMutex.Unlock()
	if fake.DeleteSpacesStub != nil {
		return fake.DeleteSpacesStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteSpacesReturns
	return fakeReturns.result1
}

func (fake *FakeManager) DeleteSpacesCallCount() int {
	fake.deleteSpacesMutex.RLock()
	defer fake.deleteSpacesMutex.RUnlock()
	return len(fake.deleteSpacesArgsForCall)
}

func (fake *FakeManager) DeleteSpacesCalls(stub func() error) {
	fake.deleteSpacesMutex.Lock()
	defer fake.deleteSpacesMutex.Unlock()
	fake.DeleteSpacesStub = stub
}

func (fake *FakeManager) DeleteSpacesReturns(result1 error) {
	fake.deleteSpacesMutex.Lock()
	defer fake.deleteSpacesMutex.Unlock()
	fake.DeleteSpacesStub = nil
	fake.deleteSpacesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) DeleteSpacesReturnsOnCall(i int, result1 error) {
	fake.deleteSpacesMutex.Lock()
	defer fake.deleteSpacesMutex.Unlock()
	fake.DeleteSpacesStub = nil
	if fake.deleteSpacesReturnsOnCall == nil {
		fake.deleteSpacesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteSpacesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) DeleteSpacesForOrg(arg1 string, arg2 string) error {
	fake.deleteSpacesForOrgMutex.Lock()
	ret, specificReturn := fake.deleteSpacesForOrgReturnsOnCall[len(fake.deleteSpacesForOrgArgsForCall)]
	fake.deleteSpacesForOrgArgsForCall = append(fake.deleteSpacesForOrgArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("DeleteSpacesForOrg", []interface{}{arg1, arg2})
	fake.deleteSpacesForOrgMutex.Unlock()
	if fake.DeleteSpacesForOrgStub != nil {
		return fake.DeleteSpacesForOrgStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteSpacesForOrgReturns
	return fakeReturns.result1
}

func (fake *FakeManager) DeleteSpacesForOrgCallCount() int {
	fake.deleteSpacesForOrgMutex.RLock()
	defer fake.deleteSpacesForOrgMutex.RUnlock()
	return len(fake.deleteSpacesForOrgArgsForCall)
}

func (fake *FakeManager) DeleteSpacesForOrgCalls(stub func(string, string) error) {
	fake.deleteSpacesForOrgMutex.Lock()
	defer fake.deleteSpacesForOrgMutex.Unlock()
	fake.DeleteSpacesForOrgStub = stub
}

func (fake *FakeManager) DeleteSpacesForOrgArgsForCall(i int) (string, string) {
	fake.deleteSpacesForOrgMutex.RLock()
	defer fake.deleteSpacesForOrgMutex.RUnlock()
	argsForCall := fake.deleteSpacesForOrgArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeManager) DeleteSpacesForOrgReturns(result1 error) {
	fake.deleteSpacesForOrgMutex.Lock()
	defer fake.deleteSpacesForOrgMutex.Unlock()
	fake.DeleteSpacesForOrgStub = nil
	fake.deleteSpacesForOrgReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) DeleteSpacesForOrgReturnsOnCall(i int, result1 error) {
	fake.deleteSpacesForOrgMutex.Lock()
	defer fake.deleteSpacesForOrgMutex.Unlock()
	fake.DeleteSpacesForOrgStub = nil
	if fake.deleteSpacesForOrgReturnsOnCall == nil {
		fake.deleteSpacesForOrgReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteSpacesForOrgReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) FindSpace(arg1 string, arg2 string) (cfclient.Space, error) {
	fake.findSpaceMutex.Lock()
	ret, specificReturn := fake.findSpaceReturnsOnCall[len(fake.findSpaceArgsForCall)]
	fake.findSpaceArgsForCall = append(fake.findSpaceArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("FindSpace", []interface{}{arg1, arg2})
	fake.findSpaceMutex.Unlock()
	if fake.FindSpaceStub != nil {
		return fake.FindSpaceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.findSpaceReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeManager) FindSpaceCallCount() int {
	fake.findSpaceMutex.RLock()
	defer fake.findSpaceMutex.RUnlock()
	return len(fake.findSpaceArgsForCall)
}

func (fake *FakeManager) FindSpaceCalls(stub func(string, string) (cfclient.Space, error)) {
	fake.findSpaceMutex.Lock()
	defer fake.findSpaceMutex.Unlock()
	fake.FindSpaceStub = stub
}

func (fake *FakeManager) FindSpaceArgsForCall(i int) (string, string) {
	fake.findSpaceMutex.RLock()
	defer fake.findSpaceMutex.RUnlock()
	argsForCall := fake.findSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeManager) FindSpaceReturns(result1 cfclient.Space, result2 error) {
	fake.findSpaceMutex.Lock()
	defer fake.findSpaceMutex.Unlock()
	fake.FindSpaceStub = nil
	fake.findSpaceReturns = struct {
		result1 cfclient.Space
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) FindSpaceReturnsOnCall(i int, result1 cfclient.Space, result2 error) {
	fake.findSpaceMutex.Lock()
	defer fake.findSpaceMutex.Unlock()
	fake.FindSpaceStub = nil
	if fake.findSpaceReturnsOnCall == nil {
		fake.findSpaceReturnsOnCall = make(map[int]struct {
			result1 cfclient.Space
			result2 error
		})
	}
	fake.findSpaceReturnsOnCall[i] = struct {
		result1 cfclient.Space
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) ListSpaces(arg1 string) ([]cfclient.Space, error) {
	fake.listSpacesMutex.Lock()
	ret, specificReturn := fake.listSpacesReturnsOnCall[len(fake.listSpacesArgsForCall)]
	fake.listSpacesArgsForCall = append(fake.listSpacesArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ListSpaces", []interface{}{arg1})
	fake.listSpacesMutex.Unlock()
	if fake.ListSpacesStub != nil {
		return fake.ListSpacesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listSpacesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeManager) ListSpacesCallCount() int {
	fake.listSpacesMutex.RLock()
	defer fake.listSpacesMutex.RUnlock()
	return len(fake.listSpacesArgsForCall)
}

func (fake *FakeManager) ListSpacesCalls(stub func(string) ([]cfclient.Space, error)) {
	fake.listSpacesMutex.Lock()
	defer fake.listSpacesMutex.Unlock()
	fake.ListSpacesStub = stub
}

func (fake *FakeManager) ListSpacesArgsForCall(i int) string {
	fake.listSpacesMutex.RLock()
	defer fake.listSpacesMutex.RUnlock()
	argsForCall := fake.listSpacesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeManager) ListSpacesReturns(result1 []cfclient.Space, result2 error) {
	fake.listSpacesMutex.Lock()
	defer fake.listSpacesMutex.Unlock()
	fake.ListSpacesStub = nil
	fake.listSpacesReturns = struct {
		result1 []cfclient.Space
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) ListSpacesReturnsOnCall(i int, result1 []cfclient.Space, result2 error) {
	fake.listSpacesMutex.Lock()
	defer fake.listSpacesMutex.Unlock()
	fake.ListSpacesStub = nil
	if fake.listSpacesReturnsOnCall == nil {
		fake.listSpacesReturnsOnCall = make(map[int]struct {
			result1 []cfclient.Space
			result2 error
		})
	}
	fake.listSpacesReturnsOnCall[i] = struct {
		result1 []cfclient.Space
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) UpdateSpaces() error {
	fake.updateSpacesMutex.Lock()
	ret, specificReturn := fake.updateSpacesReturnsOnCall[len(fake.updateSpacesArgsForCall)]
	fake.updateSpacesArgsForCall = append(fake.updateSpacesArgsForCall, struct {
	}{})
	fake.recordInvocation("UpdateSpaces", []interface{}{})
	fake.updateSpacesMutex.Unlock()
	if fake.UpdateSpacesStub != nil {
		return fake.UpdateSpacesStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.updateSpacesReturns
	return fakeReturns.result1
}

func (fake *FakeManager) UpdateSpacesCallCount() int {
	fake.updateSpacesMutex.RLock()
	defer fake.updateSpacesMutex.RUnlock()
	return len(fake.updateSpacesArgsForCall)
}

func (fake *FakeManager) UpdateSpacesCalls(stub func() error) {
	fake.updateSpacesMutex.Lock()
	defer fake.updateSpacesMutex.Unlock()
	fake.UpdateSpacesStub = stub
}

func (fake *FakeManager) UpdateSpacesReturns(result1 error) {
	fake.updateSpacesMutex.Lock()
	defer fake.updateSpacesMutex.Unlock()
	fake.UpdateSpacesStub = nil
	fake.updateSpacesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) UpdateSpacesReturnsOnCall(i int, result1 error) {
	fake.updateSpacesMutex.Lock()
	defer fake.updateSpacesMutex.Unlock()
	fake.UpdateSpacesStub = nil
	if fake.updateSpacesReturnsOnCall == nil {
		fake.updateSpacesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateSpacesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) UpdateSpacesMetadata() error {
	fake.updateSpacesMetadataMutex.Lock()
	ret, specificReturn := fake.updateSpacesMetadataReturnsOnCall[len(fake.updateSpacesMetadataArgsForCall)]
	fake.updateSpacesMetadataArgsForCall = append(fake.updateSpacesMetadataArgsForCall, struct {
	}{})
	fake.recordInvocation("UpdateSpacesMetadata", []interface{}{})
	fake.updateSpacesMetadataMutex.Unlock()
	if fake.UpdateSpacesMetadataStub != nil {
		return fake.UpdateSpacesMetadataStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.updateSpacesMetadataReturns
	return fakeReturns.result1
}

func (fake *FakeManager) UpdateSpacesMetadataCallCount() int {
	fake.updateSpacesMetadataMutex.RLock()
	defer fake.updateSpacesMetadataMutex.RUnlock()
	return len(fake.updateSpacesMetadataArgsForCall)
}

func (fake *FakeManager) UpdateSpacesMetadataCalls(stub func() error) {
	fake.updateSpacesMetadataMutex.Lock()
	defer fake.updateSpacesMetadataMutex.Unlock()
	fake.UpdateSpacesMetadataStub = stub
}

func (fake *FakeManager) UpdateSpacesMetadataReturns(result1 error) {
	fake.updateSpacesMetadataMutex.Lock()
	defer fake.updateSpacesMetadataMutex.Unlock()
	fake.UpdateSpacesMetadataStub = nil
	fake.updateSpacesMetadataReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) UpdateSpacesMetadataReturnsOnCall(i int, result1 error) {
	fake.updateSpacesMetadataMutex.Lock()
	defer fake.updateSpacesMetadataMutex.Unlock()
	fake.UpdateSpacesMetadataStub = nil
	if fake.updateSpacesMetadataReturnsOnCall == nil {
		fake.updateSpacesMetadataReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateSpacesMetadataReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createSpacesMutex.RLock()
	defer fake.createSpacesMutex.RUnlock()
	fake.deleteSpacesMutex.RLock()
	defer fake.deleteSpacesMutex.RUnlock()
	fake.deleteSpacesForOrgMutex.RLock()
	defer fake.deleteSpacesForOrgMutex.RUnlock()
	fake.findSpaceMutex.RLock()
	defer fake.findSpaceMutex.RUnlock()
	fake.listSpacesMutex.RLock()
	defer fake.listSpacesMutex.RUnlock()
	fake.updateSpacesMutex.RLock()
	defer fake.updateSpacesMutex.RUnlock()
	fake.updateSpacesMetadataMutex.RLock()
	defer fake.updateSpacesMetadataMutex.RUnlock()
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

var _ space.Manager = new(FakeManager)
