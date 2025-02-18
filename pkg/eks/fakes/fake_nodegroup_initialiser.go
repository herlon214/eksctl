// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/weaveworks/eksctl/pkg/apis/eksctl.io/v1alpha5"
	"github.com/weaveworks/eksctl/pkg/cfn/manager"
	"github.com/weaveworks/eksctl/pkg/eks"
	"github.com/weaveworks/eksctl/pkg/utils/tasks"
	"k8s.io/client-go/kubernetes"
)

type FakeNodeGroupInitialiser struct {
	DoAllNodegroupStackTasksStub        func(*tasks.TaskTree, string, string) error
	doAllNodegroupStackTasksMutex       sync.RWMutex
	doAllNodegroupStackTasksArgsForCall []struct {
		arg1 *tasks.TaskTree
		arg2 string
		arg3 string
	}
	doAllNodegroupStackTasksReturns struct {
		result1 error
	}
	doAllNodegroupStackTasksReturnsOnCall map[int]struct {
		result1 error
	}
	DoesAWSNodeUseIRSAStub        func(v1alpha5.ClusterProvider, kubernetes.Interface) (bool, error)
	doesAWSNodeUseIRSAMutex       sync.RWMutex
	doesAWSNodeUseIRSAArgsForCall []struct {
		arg1 v1alpha5.ClusterProvider
		arg2 kubernetes.Interface
	}
	doesAWSNodeUseIRSAReturns struct {
		result1 bool
		result2 error
	}
	doesAWSNodeUseIRSAReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	ExpandInstanceSelectorOptionsStub        func([]v1alpha5.NodePool, []string) error
	expandInstanceSelectorOptionsMutex       sync.RWMutex
	expandInstanceSelectorOptionsArgsForCall []struct {
		arg1 []v1alpha5.NodePool
		arg2 []string
	}
	expandInstanceSelectorOptionsReturns struct {
		result1 error
	}
	expandInstanceSelectorOptionsReturnsOnCall map[int]struct {
		result1 error
	}
	NewAWSSelectorSessionStub        func(v1alpha5.ClusterProvider)
	newAWSSelectorSessionMutex       sync.RWMutex
	newAWSSelectorSessionArgsForCall []struct {
		arg1 v1alpha5.ClusterProvider
	}
	NormalizeStub        func([]v1alpha5.NodePool, *v1alpha5.ClusterMeta) error
	normalizeMutex       sync.RWMutex
	normalizeArgsForCall []struct {
		arg1 []v1alpha5.NodePool
		arg2 *v1alpha5.ClusterMeta
	}
	normalizeReturns struct {
		result1 error
	}
	normalizeReturnsOnCall map[int]struct {
		result1 error
	}
	ValidateExistingNodeGroupsForCompatibilityStub        func(*v1alpha5.ClusterConfig, manager.StackManager) error
	validateExistingNodeGroupsForCompatibilityMutex       sync.RWMutex
	validateExistingNodeGroupsForCompatibilityArgsForCall []struct {
		arg1 *v1alpha5.ClusterConfig
		arg2 manager.StackManager
	}
	validateExistingNodeGroupsForCompatibilityReturns struct {
		result1 error
	}
	validateExistingNodeGroupsForCompatibilityReturnsOnCall map[int]struct {
		result1 error
	}
	ValidateLegacySubnetsForNodeGroupsStub        func(*v1alpha5.ClusterConfig, v1alpha5.ClusterProvider) error
	validateLegacySubnetsForNodeGroupsMutex       sync.RWMutex
	validateLegacySubnetsForNodeGroupsArgsForCall []struct {
		arg1 *v1alpha5.ClusterConfig
		arg2 v1alpha5.ClusterProvider
	}
	validateLegacySubnetsForNodeGroupsReturns struct {
		result1 error
	}
	validateLegacySubnetsForNodeGroupsReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeNodeGroupInitialiser) DoAllNodegroupStackTasks(arg1 *tasks.TaskTree, arg2 string, arg3 string) error {
	fake.doAllNodegroupStackTasksMutex.Lock()
	ret, specificReturn := fake.doAllNodegroupStackTasksReturnsOnCall[len(fake.doAllNodegroupStackTasksArgsForCall)]
	fake.doAllNodegroupStackTasksArgsForCall = append(fake.doAllNodegroupStackTasksArgsForCall, struct {
		arg1 *tasks.TaskTree
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.DoAllNodegroupStackTasksStub
	fakeReturns := fake.doAllNodegroupStackTasksReturns
	fake.recordInvocation("DoAllNodegroupStackTasks", []interface{}{arg1, arg2, arg3})
	fake.doAllNodegroupStackTasksMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeNodeGroupInitialiser) DoAllNodegroupStackTasksCallCount() int {
	fake.doAllNodegroupStackTasksMutex.RLock()
	defer fake.doAllNodegroupStackTasksMutex.RUnlock()
	return len(fake.doAllNodegroupStackTasksArgsForCall)
}

func (fake *FakeNodeGroupInitialiser) DoAllNodegroupStackTasksCalls(stub func(*tasks.TaskTree, string, string) error) {
	fake.doAllNodegroupStackTasksMutex.Lock()
	defer fake.doAllNodegroupStackTasksMutex.Unlock()
	fake.DoAllNodegroupStackTasksStub = stub
}

func (fake *FakeNodeGroupInitialiser) DoAllNodegroupStackTasksArgsForCall(i int) (*tasks.TaskTree, string, string) {
	fake.doAllNodegroupStackTasksMutex.RLock()
	defer fake.doAllNodegroupStackTasksMutex.RUnlock()
	argsForCall := fake.doAllNodegroupStackTasksArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeNodeGroupInitialiser) DoAllNodegroupStackTasksReturns(result1 error) {
	fake.doAllNodegroupStackTasksMutex.Lock()
	defer fake.doAllNodegroupStackTasksMutex.Unlock()
	fake.DoAllNodegroupStackTasksStub = nil
	fake.doAllNodegroupStackTasksReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNodeGroupInitialiser) DoAllNodegroupStackTasksReturnsOnCall(i int, result1 error) {
	fake.doAllNodegroupStackTasksMutex.Lock()
	defer fake.doAllNodegroupStackTasksMutex.Unlock()
	fake.DoAllNodegroupStackTasksStub = nil
	if fake.doAllNodegroupStackTasksReturnsOnCall == nil {
		fake.doAllNodegroupStackTasksReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.doAllNodegroupStackTasksReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNodeGroupInitialiser) DoesAWSNodeUseIRSA(arg1 v1alpha5.ClusterProvider, arg2 kubernetes.Interface) (bool, error) {
	fake.doesAWSNodeUseIRSAMutex.Lock()
	ret, specificReturn := fake.doesAWSNodeUseIRSAReturnsOnCall[len(fake.doesAWSNodeUseIRSAArgsForCall)]
	fake.doesAWSNodeUseIRSAArgsForCall = append(fake.doesAWSNodeUseIRSAArgsForCall, struct {
		arg1 v1alpha5.ClusterProvider
		arg2 kubernetes.Interface
	}{arg1, arg2})
	stub := fake.DoesAWSNodeUseIRSAStub
	fakeReturns := fake.doesAWSNodeUseIRSAReturns
	fake.recordInvocation("DoesAWSNodeUseIRSA", []interface{}{arg1, arg2})
	fake.doesAWSNodeUseIRSAMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeNodeGroupInitialiser) DoesAWSNodeUseIRSACallCount() int {
	fake.doesAWSNodeUseIRSAMutex.RLock()
	defer fake.doesAWSNodeUseIRSAMutex.RUnlock()
	return len(fake.doesAWSNodeUseIRSAArgsForCall)
}

func (fake *FakeNodeGroupInitialiser) DoesAWSNodeUseIRSACalls(stub func(v1alpha5.ClusterProvider, kubernetes.Interface) (bool, error)) {
	fake.doesAWSNodeUseIRSAMutex.Lock()
	defer fake.doesAWSNodeUseIRSAMutex.Unlock()
	fake.DoesAWSNodeUseIRSAStub = stub
}

func (fake *FakeNodeGroupInitialiser) DoesAWSNodeUseIRSAArgsForCall(i int) (v1alpha5.ClusterProvider, kubernetes.Interface) {
	fake.doesAWSNodeUseIRSAMutex.RLock()
	defer fake.doesAWSNodeUseIRSAMutex.RUnlock()
	argsForCall := fake.doesAWSNodeUseIRSAArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeNodeGroupInitialiser) DoesAWSNodeUseIRSAReturns(result1 bool, result2 error) {
	fake.doesAWSNodeUseIRSAMutex.Lock()
	defer fake.doesAWSNodeUseIRSAMutex.Unlock()
	fake.DoesAWSNodeUseIRSAStub = nil
	fake.doesAWSNodeUseIRSAReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeNodeGroupInitialiser) DoesAWSNodeUseIRSAReturnsOnCall(i int, result1 bool, result2 error) {
	fake.doesAWSNodeUseIRSAMutex.Lock()
	defer fake.doesAWSNodeUseIRSAMutex.Unlock()
	fake.DoesAWSNodeUseIRSAStub = nil
	if fake.doesAWSNodeUseIRSAReturnsOnCall == nil {
		fake.doesAWSNodeUseIRSAReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.doesAWSNodeUseIRSAReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeNodeGroupInitialiser) ExpandInstanceSelectorOptions(arg1 []v1alpha5.NodePool, arg2 []string) error {
	var arg1Copy []v1alpha5.NodePool
	if arg1 != nil {
		arg1Copy = make([]v1alpha5.NodePool, len(arg1))
		copy(arg1Copy, arg1)
	}
	var arg2Copy []string
	if arg2 != nil {
		arg2Copy = make([]string, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.expandInstanceSelectorOptionsMutex.Lock()
	ret, specificReturn := fake.expandInstanceSelectorOptionsReturnsOnCall[len(fake.expandInstanceSelectorOptionsArgsForCall)]
	fake.expandInstanceSelectorOptionsArgsForCall = append(fake.expandInstanceSelectorOptionsArgsForCall, struct {
		arg1 []v1alpha5.NodePool
		arg2 []string
	}{arg1Copy, arg2Copy})
	stub := fake.ExpandInstanceSelectorOptionsStub
	fakeReturns := fake.expandInstanceSelectorOptionsReturns
	fake.recordInvocation("ExpandInstanceSelectorOptions", []interface{}{arg1Copy, arg2Copy})
	fake.expandInstanceSelectorOptionsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeNodeGroupInitialiser) ExpandInstanceSelectorOptionsCallCount() int {
	fake.expandInstanceSelectorOptionsMutex.RLock()
	defer fake.expandInstanceSelectorOptionsMutex.RUnlock()
	return len(fake.expandInstanceSelectorOptionsArgsForCall)
}

func (fake *FakeNodeGroupInitialiser) ExpandInstanceSelectorOptionsCalls(stub func([]v1alpha5.NodePool, []string) error) {
	fake.expandInstanceSelectorOptionsMutex.Lock()
	defer fake.expandInstanceSelectorOptionsMutex.Unlock()
	fake.ExpandInstanceSelectorOptionsStub = stub
}

func (fake *FakeNodeGroupInitialiser) ExpandInstanceSelectorOptionsArgsForCall(i int) ([]v1alpha5.NodePool, []string) {
	fake.expandInstanceSelectorOptionsMutex.RLock()
	defer fake.expandInstanceSelectorOptionsMutex.RUnlock()
	argsForCall := fake.expandInstanceSelectorOptionsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeNodeGroupInitialiser) ExpandInstanceSelectorOptionsReturns(result1 error) {
	fake.expandInstanceSelectorOptionsMutex.Lock()
	defer fake.expandInstanceSelectorOptionsMutex.Unlock()
	fake.ExpandInstanceSelectorOptionsStub = nil
	fake.expandInstanceSelectorOptionsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNodeGroupInitialiser) ExpandInstanceSelectorOptionsReturnsOnCall(i int, result1 error) {
	fake.expandInstanceSelectorOptionsMutex.Lock()
	defer fake.expandInstanceSelectorOptionsMutex.Unlock()
	fake.ExpandInstanceSelectorOptionsStub = nil
	if fake.expandInstanceSelectorOptionsReturnsOnCall == nil {
		fake.expandInstanceSelectorOptionsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.expandInstanceSelectorOptionsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNodeGroupInitialiser) NewAWSSelectorSession(arg1 v1alpha5.ClusterProvider) {
	fake.newAWSSelectorSessionMutex.Lock()
	fake.newAWSSelectorSessionArgsForCall = append(fake.newAWSSelectorSessionArgsForCall, struct {
		arg1 v1alpha5.ClusterProvider
	}{arg1})
	stub := fake.NewAWSSelectorSessionStub
	fake.recordInvocation("NewAWSSelectorSession", []interface{}{arg1})
	fake.newAWSSelectorSessionMutex.Unlock()
	if stub != nil {
		fake.NewAWSSelectorSessionStub(arg1)
	}
}

func (fake *FakeNodeGroupInitialiser) NewAWSSelectorSessionCallCount() int {
	fake.newAWSSelectorSessionMutex.RLock()
	defer fake.newAWSSelectorSessionMutex.RUnlock()
	return len(fake.newAWSSelectorSessionArgsForCall)
}

func (fake *FakeNodeGroupInitialiser) NewAWSSelectorSessionCalls(stub func(v1alpha5.ClusterProvider)) {
	fake.newAWSSelectorSessionMutex.Lock()
	defer fake.newAWSSelectorSessionMutex.Unlock()
	fake.NewAWSSelectorSessionStub = stub
}

func (fake *FakeNodeGroupInitialiser) NewAWSSelectorSessionArgsForCall(i int) v1alpha5.ClusterProvider {
	fake.newAWSSelectorSessionMutex.RLock()
	defer fake.newAWSSelectorSessionMutex.RUnlock()
	argsForCall := fake.newAWSSelectorSessionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeNodeGroupInitialiser) Normalize(arg1 []v1alpha5.NodePool, arg2 *v1alpha5.ClusterMeta) error {
	var arg1Copy []v1alpha5.NodePool
	if arg1 != nil {
		arg1Copy = make([]v1alpha5.NodePool, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.normalizeMutex.Lock()
	ret, specificReturn := fake.normalizeReturnsOnCall[len(fake.normalizeArgsForCall)]
	fake.normalizeArgsForCall = append(fake.normalizeArgsForCall, struct {
		arg1 []v1alpha5.NodePool
		arg2 *v1alpha5.ClusterMeta
	}{arg1Copy, arg2})
	stub := fake.NormalizeStub
	fakeReturns := fake.normalizeReturns
	fake.recordInvocation("Normalize", []interface{}{arg1Copy, arg2})
	fake.normalizeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeNodeGroupInitialiser) NormalizeCallCount() int {
	fake.normalizeMutex.RLock()
	defer fake.normalizeMutex.RUnlock()
	return len(fake.normalizeArgsForCall)
}

func (fake *FakeNodeGroupInitialiser) NormalizeCalls(stub func([]v1alpha5.NodePool, *v1alpha5.ClusterMeta) error) {
	fake.normalizeMutex.Lock()
	defer fake.normalizeMutex.Unlock()
	fake.NormalizeStub = stub
}

func (fake *FakeNodeGroupInitialiser) NormalizeArgsForCall(i int) ([]v1alpha5.NodePool, *v1alpha5.ClusterMeta) {
	fake.normalizeMutex.RLock()
	defer fake.normalizeMutex.RUnlock()
	argsForCall := fake.normalizeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeNodeGroupInitialiser) NormalizeReturns(result1 error) {
	fake.normalizeMutex.Lock()
	defer fake.normalizeMutex.Unlock()
	fake.NormalizeStub = nil
	fake.normalizeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNodeGroupInitialiser) NormalizeReturnsOnCall(i int, result1 error) {
	fake.normalizeMutex.Lock()
	defer fake.normalizeMutex.Unlock()
	fake.NormalizeStub = nil
	if fake.normalizeReturnsOnCall == nil {
		fake.normalizeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.normalizeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNodeGroupInitialiser) ValidateExistingNodeGroupsForCompatibility(arg1 *v1alpha5.ClusterConfig, arg2 manager.StackManager) error {
	fake.validateExistingNodeGroupsForCompatibilityMutex.Lock()
	ret, specificReturn := fake.validateExistingNodeGroupsForCompatibilityReturnsOnCall[len(fake.validateExistingNodeGroupsForCompatibilityArgsForCall)]
	fake.validateExistingNodeGroupsForCompatibilityArgsForCall = append(fake.validateExistingNodeGroupsForCompatibilityArgsForCall, struct {
		arg1 *v1alpha5.ClusterConfig
		arg2 manager.StackManager
	}{arg1, arg2})
	stub := fake.ValidateExistingNodeGroupsForCompatibilityStub
	fakeReturns := fake.validateExistingNodeGroupsForCompatibilityReturns
	fake.recordInvocation("ValidateExistingNodeGroupsForCompatibility", []interface{}{arg1, arg2})
	fake.validateExistingNodeGroupsForCompatibilityMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeNodeGroupInitialiser) ValidateExistingNodeGroupsForCompatibilityCallCount() int {
	fake.validateExistingNodeGroupsForCompatibilityMutex.RLock()
	defer fake.validateExistingNodeGroupsForCompatibilityMutex.RUnlock()
	return len(fake.validateExistingNodeGroupsForCompatibilityArgsForCall)
}

func (fake *FakeNodeGroupInitialiser) ValidateExistingNodeGroupsForCompatibilityCalls(stub func(*v1alpha5.ClusterConfig, manager.StackManager) error) {
	fake.validateExistingNodeGroupsForCompatibilityMutex.Lock()
	defer fake.validateExistingNodeGroupsForCompatibilityMutex.Unlock()
	fake.ValidateExistingNodeGroupsForCompatibilityStub = stub
}

func (fake *FakeNodeGroupInitialiser) ValidateExistingNodeGroupsForCompatibilityArgsForCall(i int) (*v1alpha5.ClusterConfig, manager.StackManager) {
	fake.validateExistingNodeGroupsForCompatibilityMutex.RLock()
	defer fake.validateExistingNodeGroupsForCompatibilityMutex.RUnlock()
	argsForCall := fake.validateExistingNodeGroupsForCompatibilityArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeNodeGroupInitialiser) ValidateExistingNodeGroupsForCompatibilityReturns(result1 error) {
	fake.validateExistingNodeGroupsForCompatibilityMutex.Lock()
	defer fake.validateExistingNodeGroupsForCompatibilityMutex.Unlock()
	fake.ValidateExistingNodeGroupsForCompatibilityStub = nil
	fake.validateExistingNodeGroupsForCompatibilityReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNodeGroupInitialiser) ValidateExistingNodeGroupsForCompatibilityReturnsOnCall(i int, result1 error) {
	fake.validateExistingNodeGroupsForCompatibilityMutex.Lock()
	defer fake.validateExistingNodeGroupsForCompatibilityMutex.Unlock()
	fake.ValidateExistingNodeGroupsForCompatibilityStub = nil
	if fake.validateExistingNodeGroupsForCompatibilityReturnsOnCall == nil {
		fake.validateExistingNodeGroupsForCompatibilityReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validateExistingNodeGroupsForCompatibilityReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNodeGroupInitialiser) ValidateLegacySubnetsForNodeGroups(arg1 *v1alpha5.ClusterConfig, arg2 v1alpha5.ClusterProvider) error {
	fake.validateLegacySubnetsForNodeGroupsMutex.Lock()
	ret, specificReturn := fake.validateLegacySubnetsForNodeGroupsReturnsOnCall[len(fake.validateLegacySubnetsForNodeGroupsArgsForCall)]
	fake.validateLegacySubnetsForNodeGroupsArgsForCall = append(fake.validateLegacySubnetsForNodeGroupsArgsForCall, struct {
		arg1 *v1alpha5.ClusterConfig
		arg2 v1alpha5.ClusterProvider
	}{arg1, arg2})
	stub := fake.ValidateLegacySubnetsForNodeGroupsStub
	fakeReturns := fake.validateLegacySubnetsForNodeGroupsReturns
	fake.recordInvocation("ValidateLegacySubnetsForNodeGroups", []interface{}{arg1, arg2})
	fake.validateLegacySubnetsForNodeGroupsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeNodeGroupInitialiser) ValidateLegacySubnetsForNodeGroupsCallCount() int {
	fake.validateLegacySubnetsForNodeGroupsMutex.RLock()
	defer fake.validateLegacySubnetsForNodeGroupsMutex.RUnlock()
	return len(fake.validateLegacySubnetsForNodeGroupsArgsForCall)
}

func (fake *FakeNodeGroupInitialiser) ValidateLegacySubnetsForNodeGroupsCalls(stub func(*v1alpha5.ClusterConfig, v1alpha5.ClusterProvider) error) {
	fake.validateLegacySubnetsForNodeGroupsMutex.Lock()
	defer fake.validateLegacySubnetsForNodeGroupsMutex.Unlock()
	fake.ValidateLegacySubnetsForNodeGroupsStub = stub
}

func (fake *FakeNodeGroupInitialiser) ValidateLegacySubnetsForNodeGroupsArgsForCall(i int) (*v1alpha5.ClusterConfig, v1alpha5.ClusterProvider) {
	fake.validateLegacySubnetsForNodeGroupsMutex.RLock()
	defer fake.validateLegacySubnetsForNodeGroupsMutex.RUnlock()
	argsForCall := fake.validateLegacySubnetsForNodeGroupsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeNodeGroupInitialiser) ValidateLegacySubnetsForNodeGroupsReturns(result1 error) {
	fake.validateLegacySubnetsForNodeGroupsMutex.Lock()
	defer fake.validateLegacySubnetsForNodeGroupsMutex.Unlock()
	fake.ValidateLegacySubnetsForNodeGroupsStub = nil
	fake.validateLegacySubnetsForNodeGroupsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNodeGroupInitialiser) ValidateLegacySubnetsForNodeGroupsReturnsOnCall(i int, result1 error) {
	fake.validateLegacySubnetsForNodeGroupsMutex.Lock()
	defer fake.validateLegacySubnetsForNodeGroupsMutex.Unlock()
	fake.ValidateLegacySubnetsForNodeGroupsStub = nil
	if fake.validateLegacySubnetsForNodeGroupsReturnsOnCall == nil {
		fake.validateLegacySubnetsForNodeGroupsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validateLegacySubnetsForNodeGroupsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNodeGroupInitialiser) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.doAllNodegroupStackTasksMutex.RLock()
	defer fake.doAllNodegroupStackTasksMutex.RUnlock()
	fake.doesAWSNodeUseIRSAMutex.RLock()
	defer fake.doesAWSNodeUseIRSAMutex.RUnlock()
	fake.expandInstanceSelectorOptionsMutex.RLock()
	defer fake.expandInstanceSelectorOptionsMutex.RUnlock()
	fake.newAWSSelectorSessionMutex.RLock()
	defer fake.newAWSSelectorSessionMutex.RUnlock()
	fake.normalizeMutex.RLock()
	defer fake.normalizeMutex.RUnlock()
	fake.validateExistingNodeGroupsForCompatibilityMutex.RLock()
	defer fake.validateExistingNodeGroupsForCompatibilityMutex.RUnlock()
	fake.validateLegacySubnetsForNodeGroupsMutex.RLock()
	defer fake.validateLegacySubnetsForNodeGroupsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeNodeGroupInitialiser) recordInvocation(key string, args []interface{}) {
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

var _ eks.NodeGroupInitialiser = new(FakeNodeGroupInitialiser)
