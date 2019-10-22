/*
 * @file
 * @copyright defined in aergo/LICENSE.txt
 */

package p2p

import (
	"github.com/aergoio/aergo-lib/log"
	"github.com/aergoio/aergo/p2p/p2pcommon"
	"github.com/aergoio/aergo/p2p/p2putil"
	"github.com/aergoio/aergo/types"
	"strings"
	"sync"
)

type RaftRoleManager struct {
	p2ps      *P2P
	logger    *log.Logger
	raftBP    map[types.PeerID]bool
	raftMutex sync.Mutex
}

func (rm *RaftRoleManager) UpdateBP(toAdd []types.PeerID, toRemove []types.PeerID) {
	rm.raftMutex.Lock()
	defer rm.raftMutex.Unlock()
	changes := make([]p2pcommon.AttrModifier,0, len(toAdd)+len(toRemove))
	for _, pid := range toRemove {
		delete(rm.raftBP, pid)
		changes = append(changes, p2pcommon.AttrModifier{pid, types.PeerRole_Watcher})
		rm.logger.Debug().Str(p2putil.LogPeerID, p2putil.ShortForm(pid)).Msg("raftBP removed")
	}
	for _, pid := range toAdd {
		rm.raftBP[pid] = true
		changes = append(changes, p2pcommon.AttrModifier{pid, types.PeerRole_Producer})
		rm.logger.Debug().Str(p2putil.LogPeerID, p2putil.ShortForm(pid)).Msg("raftBP added")
	}
	rm.p2ps.pm.UpdatePeerRole(changes)
}

func (rm *RaftRoleManager) SelfRole() types.PeerRole {
	return rm.p2ps.selfRole
}

func (rm *RaftRoleManager) GetRole(pid types.PeerID) types.PeerRole {
	rm.raftMutex.Lock()
	defer rm.raftMutex.Unlock()
	if _, found := rm.raftBP[pid]; found {
		// TODO check if leader or follower
		return types.PeerRole_Producer
	} else {
		return types.PeerRole_Watcher
	}
}

func (rm *RaftRoleManager) NotifyNewBlockMsg(mo p2pcommon.MsgOrder, peers []p2pcommon.RemotePeer) (skipped, sent int) {
	// TODO filter to only contain bp and trusted node.
	for _, neighbor := range peers {
		if neighbor != nil && neighbor.State() == types.RUNNING &&
			neighbor.Role() == types.PeerRole_Watcher {
			sent++
			neighbor.SendMessage(mo)
		} else {
			skipped++
		}
	}
	return
}

type DefaultRoleManager struct {
	p2ps *P2P
}

func (rm *DefaultRoleManager) UpdateBP(toAdd []types.PeerID, toRemove []types.PeerID) {
	changes := make([]p2pcommon.AttrModifier,0, len(toAdd)+len(toRemove))
	for _, pid := range toRemove {
		changes = append(changes, p2pcommon.AttrModifier{pid, types.PeerRole_Watcher})
	}
	for _, pid := range toAdd {
		changes = append(changes, p2pcommon.AttrModifier{pid, types.PeerRole_Producer})
	}
	rm.p2ps.pm.UpdatePeerRole(changes)
}

func (rm *DefaultRoleManager) SelfRole() types.PeerRole {
	return rm.p2ps.selfRole
}

func (rm *DefaultRoleManager) GetRole(pid types.PeerID) types.PeerRole {
	prettyID := pid.Pretty()
	bps := rm.p2ps.consacc.ConsensusInfo().Bps
	for _, bp := range bps {
		if strings.Contains(bp, prettyID) {
			return types.PeerRole_Producer
		}
	}
	return types.PeerRole_Watcher
}

func (rm *DefaultRoleManager) NotifyNewBlockMsg(mo p2pcommon.MsgOrder, peers []p2pcommon.RemotePeer) (skipped, sent int) {
	// TODO filter to only contain bp and trusted node.
	for _, neighbor := range peers {
		if neighbor != nil && neighbor.State() == types.RUNNING {
			sent++
			neighbor.SendMessage(mo)
		} else {
			skipped++
		}
	}
	return
}

