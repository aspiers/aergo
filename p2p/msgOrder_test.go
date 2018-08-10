// Code generated by mockery v1.0.0. DO NOT EDIT.
package p2p

import (
	net "github.com/libp2p/go-libp2p-net"

	protocol "github.com/libp2p/go-libp2p-protocol"
	mock "github.com/stretchr/testify/mock"
)

// MockMsgOrder is an autogenerated mock type for the msgOrder type
type MockMsgOrder struct {
	mock.Mock
}

// GetProtocolID provides a mock function with given fields:
func (_m *MockMsgOrder) GetProtocolID() protocol.ID {
	ret := _m.Called()

	var r0 protocol.ID
	if rf, ok := ret.Get(0).(func() protocol.ID); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(protocol.ID)
	}

	return r0
}

// GetRequestID provides a mock function with given fields:
func (_m *MockMsgOrder) GetRequestID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// IsNeedSign provides a mock function with given fields:
func (_m *MockMsgOrder) Timestamp() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// IsGossip provides a mock function with given fields:
func (_m *MockMsgOrder) IsGossip() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsNeedSign provides a mock function with given fields:
func (_m *MockMsgOrder) IsNeedSign() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsRequest provides a mock function with given fields:
func (_m *MockMsgOrder) IsRequest() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ResponseExpected provides a mock function with given fields:
func (_m *MockMsgOrder) ResponseExpected() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// SendOver provides a mock function with given fields: s
func (_m *MockMsgOrder) SendOver(s net.Stream) error {
	ret := _m.Called(s)

	var r0 error
	if rf, ok := ret.Get(0).(func(net.Stream) error); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SignWith provides a mock function with given fields: ps
func (_m *MockMsgOrder) SignWith(ps PeerManager) error {
	ret := _m.Called(ps)

	var r0 error
	if rf, ok := ret.Get(0).(func(PeerManager) error); ok {
		r0 = rf(ps)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
