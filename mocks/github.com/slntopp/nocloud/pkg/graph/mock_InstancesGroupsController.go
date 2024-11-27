// Code generated by mockery v2.43.2. DO NOT EDIT.

package graph_mocks

import (
	context "context"

	driver "github.com/arangodb/go-driver"
	graph "github.com/slntopp/nocloud/pkg/graph"

	instances "github.com/slntopp/nocloud-proto/instances"

	mock "github.com/stretchr/testify/mock"

	statuses "github.com/slntopp/nocloud-proto/statuses"
)

// MockInstancesGroupsController is an autogenerated mock type for the InstancesGroupsController type
type MockInstancesGroupsController struct {
	mock.Mock
}

type MockInstancesGroupsController_Expecter struct {
	mock *mock.Mock
}

func (_m *MockInstancesGroupsController) EXPECT() *MockInstancesGroupsController_Expecter {
	return &MockInstancesGroupsController_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, service, g
func (_m *MockInstancesGroupsController) Create(ctx context.Context, service driver.DocumentID, g *instances.InstancesGroup) error {
	ret := _m.Called(ctx, service, g)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, driver.DocumentID, *instances.InstancesGroup) error); ok {
		r0 = rf(ctx, service, g)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockInstancesGroupsController_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockInstancesGroupsController_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - service driver.DocumentID
//   - g *instances.InstancesGroup
func (_e *MockInstancesGroupsController_Expecter) Create(ctx interface{}, service interface{}, g interface{}) *MockInstancesGroupsController_Create_Call {
	return &MockInstancesGroupsController_Create_Call{Call: _e.mock.On("Create", ctx, service, g)}
}

func (_c *MockInstancesGroupsController_Create_Call) Run(run func(ctx context.Context, service driver.DocumentID, g *instances.InstancesGroup)) *MockInstancesGroupsController_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(driver.DocumentID), args[2].(*instances.InstancesGroup))
	})
	return _c
}

func (_c *MockInstancesGroupsController_Create_Call) Return(_a0 error) *MockInstancesGroupsController_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInstancesGroupsController_Create_Call) RunAndReturn(run func(context.Context, driver.DocumentID, *instances.InstancesGroup) error) *MockInstancesGroupsController_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, service, g
func (_m *MockInstancesGroupsController) Delete(ctx context.Context, service string, g *instances.InstancesGroup) error {
	ret := _m.Called(ctx, service, g)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *instances.InstancesGroup) error); ok {
		r0 = rf(ctx, service, g)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockInstancesGroupsController_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockInstancesGroupsController_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - service string
//   - g *instances.InstancesGroup
func (_e *MockInstancesGroupsController_Expecter) Delete(ctx interface{}, service interface{}, g interface{}) *MockInstancesGroupsController_Delete_Call {
	return &MockInstancesGroupsController_Delete_Call{Call: _e.mock.On("Delete", ctx, service, g)}
}

func (_c *MockInstancesGroupsController_Delete_Call) Run(run func(ctx context.Context, service string, g *instances.InstancesGroup)) *MockInstancesGroupsController_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*instances.InstancesGroup))
	})
	return _c
}

func (_c *MockInstancesGroupsController_Delete_Call) Return(_a0 error) *MockInstancesGroupsController_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInstancesGroupsController_Delete_Call) RunAndReturn(run func(context.Context, string, *instances.InstancesGroup) error) *MockInstancesGroupsController_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// GetEdge provides a mock function with given fields: ctx, inboundNode, collection
func (_m *MockInstancesGroupsController) GetEdge(ctx context.Context, inboundNode string, collection string) (string, error) {
	ret := _m.Called(ctx, inboundNode, collection)

	if len(ret) == 0 {
		panic("no return value specified for GetEdge")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (string, error)); ok {
		return rf(ctx, inboundNode, collection)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) string); ok {
		r0 = rf(ctx, inboundNode, collection)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, inboundNode, collection)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInstancesGroupsController_GetEdge_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEdge'
type MockInstancesGroupsController_GetEdge_Call struct {
	*mock.Call
}

// GetEdge is a helper method to define mock.On call
//   - ctx context.Context
//   - inboundNode string
//   - collection string
func (_e *MockInstancesGroupsController_Expecter) GetEdge(ctx interface{}, inboundNode interface{}, collection interface{}) *MockInstancesGroupsController_GetEdge_Call {
	return &MockInstancesGroupsController_GetEdge_Call{Call: _e.mock.On("GetEdge", ctx, inboundNode, collection)}
}

func (_c *MockInstancesGroupsController_GetEdge_Call) Run(run func(ctx context.Context, inboundNode string, collection string)) *MockInstancesGroupsController_GetEdge_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockInstancesGroupsController_GetEdge_Call) Return(_a0 string, _a1 error) *MockInstancesGroupsController_GetEdge_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInstancesGroupsController_GetEdge_Call) RunAndReturn(run func(context.Context, string, string) (string, error)) *MockInstancesGroupsController_GetEdge_Call {
	_c.Call.Return(run)
	return _c
}

// GetWithAccess provides a mock function with given fields: ctx, from, id
func (_m *MockInstancesGroupsController) GetWithAccess(ctx context.Context, from driver.DocumentID, id string) (graph.InstancesGroup, error) {
	ret := _m.Called(ctx, from, id)

	if len(ret) == 0 {
		panic("no return value specified for GetWithAccess")
	}

	var r0 graph.InstancesGroup
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, driver.DocumentID, string) (graph.InstancesGroup, error)); ok {
		return rf(ctx, from, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, driver.DocumentID, string) graph.InstancesGroup); ok {
		r0 = rf(ctx, from, id)
	} else {
		r0 = ret.Get(0).(graph.InstancesGroup)
	}

	if rf, ok := ret.Get(1).(func(context.Context, driver.DocumentID, string) error); ok {
		r1 = rf(ctx, from, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInstancesGroupsController_GetWithAccess_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetWithAccess'
type MockInstancesGroupsController_GetWithAccess_Call struct {
	*mock.Call
}

// GetWithAccess is a helper method to define mock.On call
//   - ctx context.Context
//   - from driver.DocumentID
//   - id string
func (_e *MockInstancesGroupsController_Expecter) GetWithAccess(ctx interface{}, from interface{}, id interface{}) *MockInstancesGroupsController_GetWithAccess_Call {
	return &MockInstancesGroupsController_GetWithAccess_Call{Call: _e.mock.On("GetWithAccess", ctx, from, id)}
}

func (_c *MockInstancesGroupsController_GetWithAccess_Call) Run(run func(ctx context.Context, from driver.DocumentID, id string)) *MockInstancesGroupsController_GetWithAccess_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(driver.DocumentID), args[2].(string))
	})
	return _c
}

func (_c *MockInstancesGroupsController_GetWithAccess_Call) Return(_a0 graph.InstancesGroup, _a1 error) *MockInstancesGroupsController_GetWithAccess_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInstancesGroupsController_GetWithAccess_Call) RunAndReturn(run func(context.Context, driver.DocumentID, string) (graph.InstancesGroup, error)) *MockInstancesGroupsController_GetWithAccess_Call {
	_c.Call.Return(run)
	return _c
}

// Instances provides a mock function with given fields:
func (_m *MockInstancesGroupsController) Instances() graph.InstancesController {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Instances")
	}

	var r0 graph.InstancesController
	if rf, ok := ret.Get(0).(func() graph.InstancesController); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(graph.InstancesController)
		}
	}

	return r0
}

// MockInstancesGroupsController_Instances_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Instances'
type MockInstancesGroupsController_Instances_Call struct {
	*mock.Call
}

// Instances is a helper method to define mock.On call
func (_e *MockInstancesGroupsController_Expecter) Instances() *MockInstancesGroupsController_Instances_Call {
	return &MockInstancesGroupsController_Instances_Call{Call: _e.mock.On("Instances")}
}

func (_c *MockInstancesGroupsController_Instances_Call) Run(run func()) *MockInstancesGroupsController_Instances_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockInstancesGroupsController_Instances_Call) Return(_a0 graph.InstancesController) *MockInstancesGroupsController_Instances_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInstancesGroupsController_Instances_Call) RunAndReturn(run func() graph.InstancesController) *MockInstancesGroupsController_Instances_Call {
	_c.Call.Return(run)
	return _c
}

// Provide provides a mock function with given fields: ctx, group, sp
func (_m *MockInstancesGroupsController) Provide(ctx context.Context, group string, sp string) error {
	ret := _m.Called(ctx, group, sp)

	if len(ret) == 0 {
		panic("no return value specified for Provide")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, group, sp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockInstancesGroupsController_Provide_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Provide'
type MockInstancesGroupsController_Provide_Call struct {
	*mock.Call
}

// Provide is a helper method to define mock.On call
//   - ctx context.Context
//   - group string
//   - sp string
func (_e *MockInstancesGroupsController_Expecter) Provide(ctx interface{}, group interface{}, sp interface{}) *MockInstancesGroupsController_Provide_Call {
	return &MockInstancesGroupsController_Provide_Call{Call: _e.mock.On("Provide", ctx, group, sp)}
}

func (_c *MockInstancesGroupsController_Provide_Call) Run(run func(ctx context.Context, group string, sp string)) *MockInstancesGroupsController_Provide_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockInstancesGroupsController_Provide_Call) Return(_a0 error) *MockInstancesGroupsController_Provide_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInstancesGroupsController_Provide_Call) RunAndReturn(run func(context.Context, string, string) error) *MockInstancesGroupsController_Provide_Call {
	_c.Call.Return(run)
	return _c
}

// SetStatus provides a mock function with given fields: ctx, ig, status
func (_m *MockInstancesGroupsController) SetStatus(ctx context.Context, ig *instances.InstancesGroup, status statuses.NoCloudStatus) error {
	ret := _m.Called(ctx, ig, status)

	if len(ret) == 0 {
		panic("no return value specified for SetStatus")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *instances.InstancesGroup, statuses.NoCloudStatus) error); ok {
		r0 = rf(ctx, ig, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockInstancesGroupsController_SetStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetStatus'
type MockInstancesGroupsController_SetStatus_Call struct {
	*mock.Call
}

// SetStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - ig *instances.InstancesGroup
//   - status statuses.NoCloudStatus
func (_e *MockInstancesGroupsController_Expecter) SetStatus(ctx interface{}, ig interface{}, status interface{}) *MockInstancesGroupsController_SetStatus_Call {
	return &MockInstancesGroupsController_SetStatus_Call{Call: _e.mock.On("SetStatus", ctx, ig, status)}
}

func (_c *MockInstancesGroupsController_SetStatus_Call) Run(run func(ctx context.Context, ig *instances.InstancesGroup, status statuses.NoCloudStatus)) *MockInstancesGroupsController_SetStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*instances.InstancesGroup), args[2].(statuses.NoCloudStatus))
	})
	return _c
}

func (_c *MockInstancesGroupsController_SetStatus_Call) Return(err error) *MockInstancesGroupsController_SetStatus_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockInstancesGroupsController_SetStatus_Call) RunAndReturn(run func(context.Context, *instances.InstancesGroup, statuses.NoCloudStatus) error) *MockInstancesGroupsController_SetStatus_Call {
	_c.Call.Return(run)
	return _c
}

// TransferIG provides a mock function with given fields: ctx, oldSrvEdge, newSrv, ig
func (_m *MockInstancesGroupsController) TransferIG(ctx context.Context, oldSrvEdge string, newSrv driver.DocumentID, ig driver.DocumentID) error {
	ret := _m.Called(ctx, oldSrvEdge, newSrv, ig)

	if len(ret) == 0 {
		panic("no return value specified for TransferIG")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, driver.DocumentID, driver.DocumentID) error); ok {
		r0 = rf(ctx, oldSrvEdge, newSrv, ig)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockInstancesGroupsController_TransferIG_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TransferIG'
type MockInstancesGroupsController_TransferIG_Call struct {
	*mock.Call
}

// TransferIG is a helper method to define mock.On call
//   - ctx context.Context
//   - oldSrvEdge string
//   - newSrv driver.DocumentID
//   - ig driver.DocumentID
func (_e *MockInstancesGroupsController_Expecter) TransferIG(ctx interface{}, oldSrvEdge interface{}, newSrv interface{}, ig interface{}) *MockInstancesGroupsController_TransferIG_Call {
	return &MockInstancesGroupsController_TransferIG_Call{Call: _e.mock.On("TransferIG", ctx, oldSrvEdge, newSrv, ig)}
}

func (_c *MockInstancesGroupsController_TransferIG_Call) Run(run func(ctx context.Context, oldSrvEdge string, newSrv driver.DocumentID, ig driver.DocumentID)) *MockInstancesGroupsController_TransferIG_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(driver.DocumentID), args[3].(driver.DocumentID))
	})
	return _c
}

func (_c *MockInstancesGroupsController_TransferIG_Call) Return(_a0 error) *MockInstancesGroupsController_TransferIG_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInstancesGroupsController_TransferIG_Call) RunAndReturn(run func(context.Context, string, driver.DocumentID, driver.DocumentID) error) *MockInstancesGroupsController_TransferIG_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, ig, oldIg
func (_m *MockInstancesGroupsController) Update(ctx context.Context, ig *instances.InstancesGroup, oldIg *instances.InstancesGroup) error {
	ret := _m.Called(ctx, ig, oldIg)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *instances.InstancesGroup, *instances.InstancesGroup) error); ok {
		r0 = rf(ctx, ig, oldIg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockInstancesGroupsController_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockInstancesGroupsController_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - ig *instances.InstancesGroup
//   - oldIg *instances.InstancesGroup
func (_e *MockInstancesGroupsController_Expecter) Update(ctx interface{}, ig interface{}, oldIg interface{}) *MockInstancesGroupsController_Update_Call {
	return &MockInstancesGroupsController_Update_Call{Call: _e.mock.On("Update", ctx, ig, oldIg)}
}

func (_c *MockInstancesGroupsController_Update_Call) Run(run func(ctx context.Context, ig *instances.InstancesGroup, oldIg *instances.InstancesGroup)) *MockInstancesGroupsController_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*instances.InstancesGroup), args[2].(*instances.InstancesGroup))
	})
	return _c
}

func (_c *MockInstancesGroupsController_Update_Call) Return(_a0 error) *MockInstancesGroupsController_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInstancesGroupsController_Update_Call) RunAndReturn(run func(context.Context, *instances.InstancesGroup, *instances.InstancesGroup) error) *MockInstancesGroupsController_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockInstancesGroupsController creates a new instance of MockInstancesGroupsController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockInstancesGroupsController(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockInstancesGroupsController {
	mock := &MockInstancesGroupsController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
