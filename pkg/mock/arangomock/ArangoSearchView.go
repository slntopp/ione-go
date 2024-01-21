// Code generated by mockery v1.0.0. DO NOT EDIT.

package arangomock

import context "context"
import driver "github.com/arangodb/go-driver"
import mock "github.com/stretchr/testify/mock"

// ArangoSearchView is an autogenerated mock type for the ArangoSearchView type
type ArangoSearchView struct {
	mock.Mock
}

// ArangoSearchView provides a mock function with given fields:
func (_m *ArangoSearchView) ArangoSearchView() (driver.ArangoSearchView, error) {
	ret := _m.Called()

	var r0 driver.ArangoSearchView
	if rf, ok := ret.Get(0).(func() driver.ArangoSearchView); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.ArangoSearchView)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Database provides a mock function with given fields:
func (_m *ArangoSearchView) Database() driver.Database {
	ret := _m.Called()

	var r0 driver.Database
	if rf, ok := ret.Get(0).(func() driver.Database); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.Database)
		}
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *ArangoSearchView) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Properties provides a mock function with given fields: ctx
func (_m *ArangoSearchView) Properties(ctx context.Context) (driver.ArangoSearchViewProperties, error) {
	ret := _m.Called(ctx)

	var r0 driver.ArangoSearchViewProperties
	if rf, ok := ret.Get(0).(func(context.Context) driver.ArangoSearchViewProperties); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(driver.ArangoSearchViewProperties)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Remove provides a mock function with given fields: ctx
func (_m *ArangoSearchView) Remove(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetProperties provides a mock function with given fields: ctx, options
func (_m *ArangoSearchView) SetProperties(ctx context.Context, options driver.ArangoSearchViewProperties) error {
	ret := _m.Called(ctx, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, driver.ArangoSearchViewProperties) error); ok {
		r0 = rf(ctx, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Type provides a mock function with given fields:
func (_m *ArangoSearchView) Type() driver.ViewType {
	ret := _m.Called()

	var r0 driver.ViewType
	if rf, ok := ret.Get(0).(func() driver.ViewType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(driver.ViewType)
	}

	return r0
}
