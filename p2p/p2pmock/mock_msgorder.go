// Code generated by MockGen. DO NOT EDIT.
// Source: internalmsg.go

// Package p2pmock is a generated GoMock package.
package p2pmock

import (
	p2pcommon "github.com/aergoio/aergo/p2p/p2pcommon"
	types "github.com/aergoio/aergo/types"
	raftpb "github.com/aergoio/etcd/raft/raftpb"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockMsgOrder is a mock of MsgOrder interface
type MockMsgOrder struct {
	ctrl     *gomock.Controller
	recorder *MockMsgOrderMockRecorder
}

// MockMsgOrderMockRecorder is the mock recorder for MockMsgOrder
type MockMsgOrderMockRecorder struct {
	mock *MockMsgOrder
}

// NewMockMsgOrder creates a new mock instance
func NewMockMsgOrder(ctrl *gomock.Controller) *MockMsgOrder {
	mock := &MockMsgOrder{ctrl: ctrl}
	mock.recorder = &MockMsgOrderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMsgOrder) EXPECT() *MockMsgOrderMockRecorder {
	return m.recorder
}

// GetMsgID mocks base method
func (m *MockMsgOrder) GetMsgID() p2pcommon.MsgID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMsgID")
	ret0, _ := ret[0].(p2pcommon.MsgID)
	return ret0
}

// GetMsgID indicates an expected call of GetMsgID
func (mr *MockMsgOrderMockRecorder) GetMsgID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMsgID", reflect.TypeOf((*MockMsgOrder)(nil).GetMsgID))
}

// Timestamp mocks base method
func (m *MockMsgOrder) Timestamp() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Timestamp")
	ret0, _ := ret[0].(int64)
	return ret0
}

// Timestamp indicates an expected call of Timestamp
func (mr *MockMsgOrderMockRecorder) Timestamp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Timestamp", reflect.TypeOf((*MockMsgOrder)(nil).Timestamp))
}

// IsRequest mocks base method
func (m *MockMsgOrder) IsRequest() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRequest")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsRequest indicates an expected call of IsRequest
func (mr *MockMsgOrderMockRecorder) IsRequest() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRequest", reflect.TypeOf((*MockMsgOrder)(nil).IsRequest))
}

// IsNeedSign mocks base method
func (m *MockMsgOrder) IsNeedSign() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsNeedSign")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsNeedSign indicates an expected call of IsNeedSign
func (mr *MockMsgOrderMockRecorder) IsNeedSign() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsNeedSign", reflect.TypeOf((*MockMsgOrder)(nil).IsNeedSign))
}

// GetProtocolID mocks base method
func (m *MockMsgOrder) GetProtocolID() p2pcommon.SubProtocol {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProtocolID")
	ret0, _ := ret[0].(p2pcommon.SubProtocol)
	return ret0
}

// GetProtocolID indicates an expected call of GetProtocolID
func (mr *MockMsgOrderMockRecorder) GetProtocolID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProtocolID", reflect.TypeOf((*MockMsgOrder)(nil).GetProtocolID))
}

// SendTo mocks base method
func (m *MockMsgOrder) SendTo(p p2pcommon.RemotePeer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendTo", p)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendTo indicates an expected call of SendTo
func (mr *MockMsgOrderMockRecorder) SendTo(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTo", reflect.TypeOf((*MockMsgOrder)(nil).SendTo), p)
}

// CancelSend mocks base method
func (m *MockMsgOrder) CancelSend(pi p2pcommon.RemotePeer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CancelSend", pi)
}

// CancelSend indicates an expected call of CancelSend
func (mr *MockMsgOrderMockRecorder) CancelSend(pi interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelSend", reflect.TypeOf((*MockMsgOrder)(nil).CancelSend), pi)
}

// MockMoFactory is a mock of MoFactory interface
type MockMoFactory struct {
	ctrl     *gomock.Controller
	recorder *MockMoFactoryMockRecorder
}

// MockMoFactoryMockRecorder is the mock recorder for MockMoFactory
type MockMoFactoryMockRecorder struct {
	mock *MockMoFactory
}

// NewMockMoFactory creates a new mock instance
func NewMockMoFactory(ctrl *gomock.Controller) *MockMoFactory {
	mock := &MockMoFactory{ctrl: ctrl}
	mock.recorder = &MockMoFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMoFactory) EXPECT() *MockMoFactoryMockRecorder {
	return m.recorder
}

// NewMsgRequestOrder mocks base method
func (m *MockMoFactory) NewMsgRequestOrder(expectResponse bool, protocolID p2pcommon.SubProtocol, message p2pcommon.MessageBody) p2pcommon.MsgOrder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewMsgRequestOrder", expectResponse, protocolID, message)
	ret0, _ := ret[0].(p2pcommon.MsgOrder)
	return ret0
}

// NewMsgRequestOrder indicates an expected call of NewMsgRequestOrder
func (mr *MockMoFactoryMockRecorder) NewMsgRequestOrder(expectResponse, protocolID, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewMsgRequestOrder", reflect.TypeOf((*MockMoFactory)(nil).NewMsgRequestOrder), expectResponse, protocolID, message)
}

// NewMsgRequestOrderWithReceiver mocks base method
func (m *MockMoFactory) NewMsgRequestOrderWithReceiver(respReceiver p2pcommon.ResponseReceiver, protocolID p2pcommon.SubProtocol, message p2pcommon.MessageBody) p2pcommon.MsgOrder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewMsgRequestOrderWithReceiver", respReceiver, protocolID, message)
	ret0, _ := ret[0].(p2pcommon.MsgOrder)
	return ret0
}

// NewMsgRequestOrderWithReceiver indicates an expected call of NewMsgRequestOrderWithReceiver
func (mr *MockMoFactoryMockRecorder) NewMsgRequestOrderWithReceiver(respReceiver, protocolID, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewMsgRequestOrderWithReceiver", reflect.TypeOf((*MockMoFactory)(nil).NewMsgRequestOrderWithReceiver), respReceiver, protocolID, message)
}

// NewMsgResponseOrder mocks base method
func (m *MockMoFactory) NewMsgResponseOrder(reqID p2pcommon.MsgID, protocolID p2pcommon.SubProtocol, message p2pcommon.MessageBody) p2pcommon.MsgOrder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewMsgResponseOrder", reqID, protocolID, message)
	ret0, _ := ret[0].(p2pcommon.MsgOrder)
	return ret0
}

// NewMsgResponseOrder indicates an expected call of NewMsgResponseOrder
func (mr *MockMoFactoryMockRecorder) NewMsgResponseOrder(reqID, protocolID, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewMsgResponseOrder", reflect.TypeOf((*MockMoFactory)(nil).NewMsgResponseOrder), reqID, protocolID, message)
}

// NewMsgBlkBroadcastOrder mocks base method
func (m *MockMoFactory) NewMsgBlkBroadcastOrder(noticeMsg *types.NewBlockNotice) p2pcommon.MsgOrder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewMsgBlkBroadcastOrder", noticeMsg)
	ret0, _ := ret[0].(p2pcommon.MsgOrder)
	return ret0
}

// NewMsgBlkBroadcastOrder indicates an expected call of NewMsgBlkBroadcastOrder
func (mr *MockMoFactoryMockRecorder) NewMsgBlkBroadcastOrder(noticeMsg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewMsgBlkBroadcastOrder", reflect.TypeOf((*MockMoFactory)(nil).NewMsgBlkBroadcastOrder), noticeMsg)
}

// NewMsgTxBroadcastOrder mocks base method
func (m *MockMoFactory) NewMsgTxBroadcastOrder(noticeMsg *types.NewTransactionsNotice) p2pcommon.MsgOrder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewMsgTxBroadcastOrder", noticeMsg)
	ret0, _ := ret[0].(p2pcommon.MsgOrder)
	return ret0
}

// NewMsgTxBroadcastOrder indicates an expected call of NewMsgTxBroadcastOrder
func (mr *MockMoFactoryMockRecorder) NewMsgTxBroadcastOrder(noticeMsg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewMsgTxBroadcastOrder", reflect.TypeOf((*MockMoFactory)(nil).NewMsgTxBroadcastOrder), noticeMsg)
}

// NewMsgBPBroadcastOrder mocks base method
func (m *MockMoFactory) NewMsgBPBroadcastOrder(noticeMsg *types.BlockProducedNotice) p2pcommon.MsgOrder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewMsgBPBroadcastOrder", noticeMsg)
	ret0, _ := ret[0].(p2pcommon.MsgOrder)
	return ret0
}

// NewMsgBPBroadcastOrder indicates an expected call of NewMsgBPBroadcastOrder
func (mr *MockMoFactoryMockRecorder) NewMsgBPBroadcastOrder(noticeMsg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewMsgBPBroadcastOrder", reflect.TypeOf((*MockMoFactory)(nil).NewMsgBPBroadcastOrder), noticeMsg)
}

// NewRaftMsgOrder mocks base method
func (m *MockMoFactory) NewRaftMsgOrder(msgType raftpb.MessageType, raftMsg *raftpb.Message) p2pcommon.MsgOrder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewRaftMsgOrder", msgType, raftMsg)
	ret0, _ := ret[0].(p2pcommon.MsgOrder)
	return ret0
}

// NewRaftMsgOrder indicates an expected call of NewRaftMsgOrder
func (mr *MockMoFactoryMockRecorder) NewRaftMsgOrder(msgType, raftMsg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewRaftMsgOrder", reflect.TypeOf((*MockMoFactory)(nil).NewRaftMsgOrder), msgType, raftMsg)
}

// NewTossMsgOrder mocks base method
func (m *MockMoFactory) NewTossMsgOrder(orgMsg p2pcommon.Message) p2pcommon.MsgOrder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewTossMsgOrder", orgMsg)
	ret0, _ := ret[0].(p2pcommon.MsgOrder)
	return ret0
}

// NewTossMsgOrder indicates an expected call of NewTossMsgOrder
func (mr *MockMoFactoryMockRecorder) NewTossMsgOrder(orgMsg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewTossMsgOrder", reflect.TypeOf((*MockMoFactory)(nil).NewTossMsgOrder), orgMsg)
}
