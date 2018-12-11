/*
 * @file
 * @copyright defined in aergo/LICENSE.txt
 */

package pmap

import (
	"fmt"
	"github.com/aergoio/aergo/config"
	"github.com/golang/mock/gomock"
	"reflect"
	"sync"
	"testing"

	"github.com/aergoio/aergo-actor/actor"
	"github.com/aergoio/aergo/p2p"
	"github.com/aergoio/aergo/p2p/mocks"
	"github.com/aergoio/aergo/pkg/component"
	"github.com/aergoio/aergo/types"
	"github.com/golang/protobuf/proto"
	"github.com/libp2p/go-libp2p-peer"
)

var (
	pmapDummyCfg = &config.P2PConfig{}
	pmapDummyP2PS = &p2p.P2P{}
)

func TestPeerMapService_BeforeStop(t *testing.T) {

	type fields struct {
		BaseComponent *component.BaseComponent
		listen        bool
		peerRegistry  map[peer.ID]p2p.PeerMeta
	}
	tests := []struct {
		name   string
		fields fields

		calledStreamHandler bool
	}{
		{"Tlisten", fields{listen: true}, true},
		{"TNot", fields{listen: false}, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			p2ps := &p2p.P2P{}
			cfg := &config.P2PConfig{}
			mockNT := mock_p2p.NewMockNetworkTransport(ctrl)
			pms := NewMapService(cfg, p2ps)
			pms.listen = tt.fields.listen
			pms.nt = mockNT

			if tt.calledStreamHandler {
				//mockNT.EXPECT().SetStreamHandler(p2p.AergoMapSub, pms.onConnect).Times(1)
				mockNT.EXPECT().RemoveStreamHandler(p2p.AergoMapSub).Times(1)
			}
			//pms.AfterStart()

			pms.BeforeStop()


			ctrl.Finish()
		})
	}
}

func TestPeerMapService_readRequest(t *testing.T) {

	dummyMeta := p2p.PeerMeta{ID:""}
	type args struct {
		meta p2p.PeerMeta
		msgStub *p2p.V030Message
		readErr error
	}
	tests := []struct {
		name    string
		args    args

		//want    p2p.Message
		//want1   *types.MapQuery
		wantErr bool
	}{
		{"TNormal", args{meta:dummyMeta, readErr:nil}, false },
		{"TError", args{meta:dummyMeta, readErr:fmt.Errorf("testerr")}, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockNT := mock_p2p.NewMockNetworkTransport(ctrl)
			pms := NewMapService(pmapDummyCfg, pmapDummyP2PS)
			pms.nt = mockNT

			msgStub := &p2p.V030Message{}
			mockRd := mock_p2p.NewMockMsgReader(ctrl)

			mockRd.EXPECT().ReadMsg().Times(1	).Return(msgStub, tt.args.readErr)

			got, got1, err := pms.readRequest(tt.args.meta, mockRd)
			if (err != nil) != tt.wantErr {
				t.Errorf("PeerMapService.readRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got == nil {
					t.Errorf("PeerMapService.readRequest() got = %v, want %v", got, "not nil")
				}
				if got1 == nil {
					t.Errorf("PeerMapService.readRequest() got = %v, want %v", got, "not nil")
				}
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("PeerMapService.readRequest() got = %v, want %v", got, tt.want)
			//}
			//if !reflect.DeepEqual(got1, tt.want1) {
			//	t.Errorf("PeerMapService.readRequest() got1 = %v, want %v", got1, tt.want1)
			//}
			ctrl.Finish()

		})
	}
}

func TestPeerMapService_handleQuery(t *testing.T) {
	type fields struct {
		BaseComponent *component.BaseComponent
		p2ps          *p2p.P2P
		listen        bool
		nt            p2p.NetworkTransport
		mutex         *sync.Mutex
		peerRegistry  map[peer.ID]p2p.PeerMeta
	}
	type args struct {
		container p2p.Message
		query     *types.MapQuery
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *types.MapResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pms := &PeerMapService{
				BaseComponent: tt.fields.BaseComponent,
				p2ps:          tt.fields.p2ps,
				listen:        tt.fields.listen,
				nt:            tt.fields.nt,
				mutex:         tt.fields.mutex,
				peerRegistry:  tt.fields.peerRegistry,
			}
			got, err := pms.handleQuery(tt.args.container, tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("PeerMapService.handleQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PeerMapService.handleQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPeerMapService_registerPeer(t *testing.T) {
	type args struct {
		receivedMeta p2p.PeerMeta
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"TSingle",args{}, false},
		{"TMulti",args{}, false},
		{"TDup",args{}, false},
		{"TConcurrent",args{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockNT := mock_p2p.NewMockNetworkTransport(ctrl)
			pms := NewMapService(pmapDummyCfg, pmapDummyP2PS)
			pms.nt = mockNT

			if err := pms.registerPeer(tt.args.receivedMeta); (err != nil) != tt.wantErr {
				t.Errorf("PeerMapService.registerPeer() error = %v, wantErr %v", err, tt.wantErr)
			}

			ctrl.Finish()

		})
	}
}

func TestPeerMapService_writeResponse(t *testing.T) {
	type fields struct {
		BaseComponent *component.BaseComponent
		p2ps          *p2p.P2P
		listen        bool
		nt            p2p.NetworkTransport
		mutex         *sync.Mutex
		peerRegistry  map[peer.ID]p2p.PeerMeta
	}
	type args struct {
		reqContainer p2p.Message
		meta         p2p.PeerMeta
		resp         *types.MapResponse
		wt           p2p.MsgWriter
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pms := &PeerMapService{
				BaseComponent: tt.fields.BaseComponent,
				p2ps:          tt.fields.p2ps,
				listen:        tt.fields.listen,
				nt:            tt.fields.nt,
				mutex:         tt.fields.mutex,
				peerRegistry:  tt.fields.peerRegistry,
			}
			if err := pms.writeResponse(tt.args.reqContainer, tt.args.meta, tt.args.resp, tt.args.wt); (err != nil) != tt.wantErr {
				t.Errorf("PeerMapService.writeResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPeerMapService_Receive(t *testing.T) {
	type fields struct {
		BaseComponent *component.BaseComponent
		p2ps          *p2p.P2P
		listen        bool
		nt            p2p.NetworkTransport
		mutex         *sync.Mutex
		peerRegistry  map[peer.ID]p2p.PeerMeta
	}
	type args struct {
		context actor.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pms := &PeerMapService{
				BaseComponent: tt.fields.BaseComponent,
				p2ps:          tt.fields.p2ps,
				listen:        tt.fields.listen,
				nt:            tt.fields.nt,
				mutex:         tt.fields.mutex,
				peerRegistry:  tt.fields.peerRegistry,
			}
			pms.Receive(tt.args.context)
		})
	}
}

func Test_createV030Message(t *testing.T) {
	type args struct {
		msgID       p2p.MsgID
		orgID       p2p.MsgID
		subProtocol p2p.SubProtocol
		innerMsg    proto.Message
	}
	tests := []struct {
		name    string
		args    args
		want    *p2p.V030Message
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createV030Message(tt.args.msgID, tt.args.orgID, tt.args.subProtocol, tt.args.innerMsg)
			if (err != nil) != tt.wantErr {
				t.Errorf("createV030Message() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createV030Message() = %v, want %v", got, tt.want)
			}
		})
	}
}
