// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import domain "kom-service-template/domain"
import mock "github.com/stretchr/testify/mock"
import reflect "reflect"

// BookRepository is an autogenerated mock type for the BookRepository type
type BookRepository struct {
	mock.Mock
}

// CreateMany provides a mock function with given fields: databaseName, collectionName, books
func (_m *BookRepository) CreateMany(databaseName string, collectionName string, books []domain.Book) (interface{}, error) {
	ret := _m.Called(databaseName, collectionName, books)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string, string, []domain.Book) interface{}); ok {
		r0 = rf(databaseName, collectionName, books)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, []domain.Book) error); ok {
		r1 = rf(databaseName, collectionName, books)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: databaseName, collectionName, filter
func (_m *BookRepository) Delete(databaseName string, collectionName string, filter interface{}) (interface{}, error) {
	ret := _m.Called(databaseName, collectionName, filter)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string, string, interface{}) interface{}); ok {
		r0 = rf(databaseName, collectionName, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, interface{}) error); ok {
		r1 = rf(databaseName, collectionName, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Read provides a mock function with given fields: databaseName, collectionName, filter, limit, dataModel
func (_m *BookRepository) Read(databaseName string, collectionName string, filter interface{}, limit int64, dataModel reflect.Type) (interface{}, error) {
	ret := _m.Called(databaseName, collectionName, filter, limit, dataModel)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string, string, interface{}, int64, reflect.Type) interface{}); ok {
		r0 = rf(databaseName, collectionName, filter, limit, dataModel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, interface{}, int64, reflect.Type) error); ok {
		r1 = rf(databaseName, collectionName, filter, limit, dataModel)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: databaseName, collectionName, filter, update
func (_m *BookRepository) Update(databaseName string, collectionName string, filter interface{}, update interface{}) (interface{}, error) {
	ret := _m.Called(databaseName, collectionName, filter, update)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string, string, interface{}, interface{}) interface{}); ok {
		r0 = rf(databaseName, collectionName, filter, update)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, interface{}, interface{}) error); ok {
		r1 = rf(databaseName, collectionName, filter, update)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
