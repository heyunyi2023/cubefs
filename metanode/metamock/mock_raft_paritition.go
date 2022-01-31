package metamock

import (
	"context"
	"github.com/chubaofs/chubaofs/raftstore"
	"github.com/tiglabs/raft/proto"
)

type ApplyFunc func(mp interface{},command []byte, index uint64) (resp interface{}, err error)

type MockPartition struct {
	Id      uint64
	walPath string
	applyId uint64
	Buff    []byte
	Mp  	interface{}
	Apply   ApplyFunc
}

func (m MockPartition) Submit(cmd []byte) (resp interface{}, err error) {
	m.applyId++
	m.Buff = make([]byte, len(cmd))
	copy(m.Buff, cmd)

	return m.Apply(m.Mp, cmd, m.applyId)
}

func (m MockPartition) SubmitWithCtx(ctx context.Context, cmd []byte) (resp interface{}, err error) {
	m.applyId++
	m.Buff = make([]byte, len(cmd))
	copy(m.Buff, cmd)
	return m.Apply(m.Mp, cmd, m.applyId)
}

func (m MockPartition) ChangeMember(changeType proto.ConfChangeType, peer proto.Peer, context []byte) (resp interface{}, err error) {
	panic("implement me")
}

func (m MockPartition) ResetMember(peers []proto.Peer, context []byte) (err error) {
	panic("implement me")
}

func (m MockPartition) Stop() error {
	panic("implement me")
}

func (m MockPartition) Delete() error {
	panic("implement me")
}

func (m MockPartition) Expired() error {
	panic("implement me")
}

func (m MockPartition) Status() (status *raftstore.PartitionStatus) {
	panic("implement me")
}

func (m MockPartition) LeaderTerm() (leaderID, term uint64) {
	return m.Id, 0
}

func (m MockPartition) IsRaftLeader() bool {
	return true
}

func (m MockPartition) AppliedIndex() uint64 {
	panic("implement me")
}

func (m MockPartition) CommittedIndex() uint64 {
	panic("implement me")
}

func (m MockPartition) Truncate(index uint64) {
	panic("implement me")
}

func (m MockPartition) TryToLeader(nodeID uint64) error {
	panic("implement me")
}

func (m MockPartition) IsOfflinePeer() bool {
	panic("implement me")
}
