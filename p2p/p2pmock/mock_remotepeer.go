// Code generated by MockGen. DO NOT EDIT.
// Source: remotepeer.go

// Package p2pmock is a generated GoMock package.
package p2pmock

import (
	p2pcommon "github.com/aergoio/aergo/p2p/p2pcommon"
	types "github.com/aergoio/aergo/types"
	gomock "github.com/golang/mock/gomock"
	go_libp2p_core "github.com/libp2p/go-libp2p-core"
	reflect "reflect"
	time "time"
)

// MockRemotePeer is a mock of RemotePeer interface
type MockRemotePeer struct {
	ctrl     *gomock.Controller
	recorder *MockRemotePeerMockRecorder
}

// MockRemotePeerMockRecorder is the mock recorder for MockRemotePeer
type MockRemotePeerMockRecorder struct {
	mock *MockRemotePeer
}

// NewMockRemotePeer creates a new mock instance
func NewMockRemotePeer(ctrl *gomock.Controller) *MockRemotePeer {
	mock := &MockRemotePeer{ctrl: ctrl}
	mock.recorder = &MockRemotePeerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRemotePeer) EXPECT() *MockRemotePeerMockRecorder {
	return m.recorder
}

// ID mocks base method
func (m *MockRemotePeer) ID() go_libp2p_core.PeerID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(go_libp2p_core.PeerID)
	return ret0
}

// ID indicates an expected call of ID
func (mr *MockRemotePeerMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockRemotePeer)(nil).ID))
}

// Meta mocks base method
func (m *MockRemotePeer) Meta() p2pcommon.PeerMeta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Meta")
	ret0, _ := ret[0].(p2pcommon.PeerMeta)
	return ret0
}

// Meta indicates an expected call of Meta
func (mr *MockRemotePeerMockRecorder) Meta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Meta", reflect.TypeOf((*MockRemotePeer)(nil).Meta))
}

// ManageNumber mocks base method
func (m *MockRemotePeer) ManageNumber() uint32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ManageNumber")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// ManageNumber indicates an expected call of ManageNumber
func (mr *MockRemotePeerMockRecorder) ManageNumber() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ManageNumber", reflect.TypeOf((*MockRemotePeer)(nil).ManageNumber))
}

// Name mocks base method
func (m *MockRemotePeer) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockRemotePeerMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockRemotePeer)(nil).Name))
}

// Version mocks base method
func (m *MockRemotePeer) Version() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(string)
	return ret0
}

// Version indicates an expected call of Version
func (mr *MockRemotePeerMockRecorder) Version() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockRemotePeer)(nil).Version))
}

// AddMessageHandler mocks base method
func (m *MockRemotePeer) AddMessageHandler(subProtocol p2pcommon.SubProtocol, handler p2pcommon.MessageHandler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddMessageHandler", subProtocol, handler)
}

// AddMessageHandler indicates an expected call of AddMessageHandler
func (mr *MockRemotePeerMockRecorder) AddMessageHandler(subProtocol, handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMessageHandler", reflect.TypeOf((*MockRemotePeer)(nil).AddMessageHandler), subProtocol, handler)
}

// State mocks base method
func (m *MockRemotePeer) State() types.PeerState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "State")
	ret0, _ := ret[0].(types.PeerState)
	return ret0
}

// State indicates an expected call of State
func (mr *MockRemotePeerMockRecorder) State() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*MockRemotePeer)(nil).State))
}

// LastStatus mocks base method
func (m *MockRemotePeer) LastStatus() *types.LastBlockStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastStatus")
	ret0, _ := ret[0].(*types.LastBlockStatus)
	return ret0
}

// LastStatus indicates an expected call of LastStatus
func (mr *MockRemotePeerMockRecorder) LastStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastStatus", reflect.TypeOf((*MockRemotePeer)(nil).LastStatus))
}

// RunPeer mocks base method
func (m *MockRemotePeer) RunPeer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RunPeer")
}

// RunPeer indicates an expected call of RunPeer
func (mr *MockRemotePeerMockRecorder) RunPeer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunPeer", reflect.TypeOf((*MockRemotePeer)(nil).RunPeer))
}

// Stop mocks base method
func (m *MockRemotePeer) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockRemotePeerMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockRemotePeer)(nil).Stop))
}

// SendMessage mocks base method
func (m *MockRemotePeer) SendMessage(msg p2pcommon.MsgOrder) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendMessage", msg)
}

// SendMessage indicates an expected call of SendMessage
func (mr *MockRemotePeerMockRecorder) SendMessage(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockRemotePeer)(nil).SendMessage), msg)
}

// SendAndWaitMessage mocks base method
func (m *MockRemotePeer) SendAndWaitMessage(msg p2pcommon.MsgOrder, ttl time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendAndWaitMessage", msg, ttl)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendAndWaitMessage indicates an expected call of SendAndWaitMessage
func (mr *MockRemotePeerMockRecorder) SendAndWaitMessage(msg, ttl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAndWaitMessage", reflect.TypeOf((*MockRemotePeer)(nil).SendAndWaitMessage), msg, ttl)
}

// PushTxsNotice mocks base method
func (m *MockRemotePeer) PushTxsNotice(txHashes []types.TxID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PushTxsNotice", txHashes)
}

// PushTxsNotice indicates an expected call of PushTxsNotice
func (mr *MockRemotePeerMockRecorder) PushTxsNotice(txHashes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushTxsNotice", reflect.TypeOf((*MockRemotePeer)(nil).PushTxsNotice), txHashes)
}

// ConsumeRequest mocks base method
func (m *MockRemotePeer) ConsumeRequest(msgID p2pcommon.MsgID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ConsumeRequest", msgID)
}

// ConsumeRequest indicates an expected call of ConsumeRequest
func (mr *MockRemotePeerMockRecorder) ConsumeRequest(msgID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsumeRequest", reflect.TypeOf((*MockRemotePeer)(nil).ConsumeRequest), msgID)
}

// GetReceiver mocks base method
func (m *MockRemotePeer) GetReceiver(id p2pcommon.MsgID) p2pcommon.ResponseReceiver {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReceiver", id)
	ret0, _ := ret[0].(p2pcommon.ResponseReceiver)
	return ret0
}

// GetReceiver indicates an expected call of GetReceiver
func (mr *MockRemotePeerMockRecorder) GetReceiver(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReceiver", reflect.TypeOf((*MockRemotePeer)(nil).GetReceiver), id)
}

// UpdateBlkCache mocks base method
func (m *MockRemotePeer) UpdateBlkCache(blkHash []byte, blkNumber uint64) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBlkCache", blkHash, blkNumber)
	ret0, _ := ret[0].(bool)
	return ret0
}

// UpdateBlkCache indicates an expected call of UpdateBlkCache
func (mr *MockRemotePeerMockRecorder) UpdateBlkCache(blkHash, blkNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBlkCache", reflect.TypeOf((*MockRemotePeer)(nil).UpdateBlkCache), blkHash, blkNumber)
}

// UpdateTxCache mocks base method
func (m *MockRemotePeer) UpdateTxCache(hashes []types.TxID) []types.TxID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTxCache", hashes)
	ret0, _ := ret[0].([]types.TxID)
	return ret0
}

// UpdateTxCache indicates an expected call of UpdateTxCache
func (mr *MockRemotePeerMockRecorder) UpdateTxCache(hashes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTxCache", reflect.TypeOf((*MockRemotePeer)(nil).UpdateTxCache), hashes)
}

// UpdateLastNotice mocks base method
func (m *MockRemotePeer) UpdateLastNotice(blkHash []byte, blkNumber uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateLastNotice", blkHash, blkNumber)
}

// UpdateLastNotice indicates an expected call of UpdateLastNotice
func (mr *MockRemotePeerMockRecorder) UpdateLastNotice(blkHash, blkNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLastNotice", reflect.TypeOf((*MockRemotePeer)(nil).UpdateLastNotice), blkHash, blkNumber)
}

// MF mocks base method
func (m *MockRemotePeer) MF() p2pcommon.MoFactory {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MF")
	ret0, _ := ret[0].(p2pcommon.MoFactory)
	return ret0
}

// MF indicates an expected call of MF
func (mr *MockRemotePeerMockRecorder) MF() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MF", reflect.TypeOf((*MockRemotePeer)(nil).MF))
}
