// Code generated by mockery v1.0.0. DO NOT EDIT.

package arangomock

import context "context"
import driver "github.com/arangodb/go-driver"
import mock "github.com/stretchr/testify/mock"

// DatabaseCollections is an autogenerated mock type for the DatabaseCollections type
type DatabaseCollections struct {
	mock.Mock
}

// Collection provides a mock function with given fields: ctx, name
func (_m *DatabaseCollections) Collection(ctx context.Context, name string) (driver.Collection, error) {
	ret := _m.Called(ctx, name)

	var r0 driver.Collection
	if rf, ok := ret.Get(0).(func(context.Context, string) driver.Collection); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.Collection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CollectionExists provides a mock function with given fields: ctx, name
func (_m *DatabaseCollections) CollectionExists(ctx context.Context, name string) (bool, error) {
	ret := _m.Called(ctx, name)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Collections provides a mock function with given fields: ctx
func (_m *DatabaseCollections) Collections(ctx context.Context) ([]driver.Collection, error) {
	ret := _m.Called(ctx)

	var r0 []driver.Collection
	if rf, ok := ret.Get(0).(func(context.Context) []driver.Collection); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]driver.Collection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateCollection provides a mock function with given fields: ctx, name, options
func (_m *DatabaseCollections) CreateCollection(ctx context.Context, name string, options *driver.CreateCollectionOptions) (driver.Collection, error) {
	ret := _m.Called(ctx, name, options)

	var r0 driver.Collection
	if rf, ok := ret.Get(0).(func(context.Context, string, *driver.CreateCollectionOptions) driver.Collection); ok {
		r0 = rf(ctx, name, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.Collection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *driver.CreateCollectionOptions) error); ok {
		r1 = rf(ctx, name, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
