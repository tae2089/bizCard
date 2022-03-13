// Code generated by mockery v2.10.0. DO NOT EDIT.

package application

import (
	domain "bizCard/domain"

	mock "github.com/stretchr/testify/mock"
)

// MockBizCardService is an autogenerated mock type for the BizCardService type
type MockBizCardService struct {
	mock.Mock
}

// DeleteBizCard provides a mock function with given fields: uid
func (_m *MockBizCardService) DeleteBizCard(uid int) string {
	ret := _m.Called(uid)

	var r0 string
	if rf, ok := ret.Get(0).(func(int) string); ok {
		r0 = rf(uid)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// FindBizCard provides a mock function with given fields: uid
func (_m *MockBizCardService) FindBizCard(uid int) *domain.BizCardInfo {
	ret := _m.Called(uid)

	var r0 *domain.BizCardInfo
	if rf, ok := ret.Get(0).(func(int) *domain.BizCardInfo); ok {
		r0 = rf(uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.BizCardInfo)
		}
	}

	return r0
}

// RegisterBizCard provides a mock function with given fields: bizCardDto
func (_m *MockBizCardService) RegisterBizCard(bizCardDto *domain.BizCardRegister) *domain.BizCardInfo {
	ret := _m.Called(bizCardDto)

	var r0 *domain.BizCardInfo
	if rf, ok := ret.Get(0).(func(*domain.BizCardRegister) *domain.BizCardInfo); ok {
		r0 = rf(bizCardDto)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.BizCardInfo)
		}
	}

	return r0
}

// UpdateBizCard provides a mock function with given fields: uid, bizCardUpdate
func (_m *MockBizCardService) UpdateBizCard(uid int, bizCardUpdate *domain.BizCardUpdate) *domain.BizCardInfo {
	ret := _m.Called(uid, bizCardUpdate)

	var r0 *domain.BizCardInfo
	if rf, ok := ret.Get(0).(func(int, *domain.BizCardUpdate) *domain.BizCardInfo); ok {
		r0 = rf(uid, bizCardUpdate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.BizCardInfo)
		}
	}

	return r0
}
