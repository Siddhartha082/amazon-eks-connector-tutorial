// Code generated by mockery 2.9.0. DO NOT EDIT.

package ssm

import mock "github.com/stretchr/testify/mock"

// MockClient is an autogenerated mock type for the Client type
type MockClient struct {
	mock.Mock
}

// Region provides a mock function with given fields:
func (_m *MockClient) Region() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// RegisterManagedInstance provides a mock function with given fields: activationID, activationCode, publicKey, publicKeyType, fingerprint
func (_m *MockClient) RegisterManagedInstance(activationID string, activationCode string, publicKey string, publicKeyType string, fingerprint string) (string, error) {
	ret := _m.Called(activationID, activationCode, publicKey, publicKeyType, fingerprint)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string, string, string, string) string); ok {
		r0 = rf(activationID, activationCode, publicKey, publicKeyType, fingerprint)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, string, string) error); ok {
		r1 = rf(activationID, activationCode, publicKey, publicKeyType, fingerprint)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
