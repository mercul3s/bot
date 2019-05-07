// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// DataStore is an autogenerated mock type for the DataStore type
type DataStore struct {
	mock.Mock
}

// AddItem provides a mock function with given fields: key, value, bucket
func (_m *DataStore) AddItem(key string, value string, bucket string) error {
	ret := _m.Called(key, value, bucket)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(key, value, bucket)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateBuckets provides a mock function with given fields: buckets
func (_m *DataStore) CreateBuckets(buckets string) error {
	ret := _m.Called(buckets)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(buckets)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListItems provides a mock function with given fields: bucket
func (_m *DataStore) ListItems(bucket string) ([]string, error) {
	ret := _m.Called(bucket)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(bucket)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(bucket)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveItem provides a mock function with given fields: key, bucket
func (_m *DataStore) RemoveItem(key string, bucket string) error {
	ret := _m.Called(key, bucket)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(key, bucket)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
