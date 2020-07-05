package raft

//
// support for Raft and kvraft to save persistent
// Raft state (log &c) and k/v server snapshots.
//
// we will use the original persister.go to test your code for grading.
// so, while you can modify this code to help you debug, please
// test with the original before submitting.
//

import "sync"

// Persister represents a persister.
type Persister struct {
	mu        sync.Mutex
	raftstate []byte
	snapshot  []byte
}

// MakePersister returns a new persister.
func MakePersister() *Persister {
	return &Persister{}
}

// Copy copies the persister.
func (ps *Persister) Copy() *Persister {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	np := MakePersister()
	np.raftstate = ps.raftstate
	np.snapshot = ps.snapshot
	return np
}

// SaveRaftState saves the state.
func (ps *Persister) SaveRaftState(state []byte) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ps.raftstate = state
}

// ReadRaftState returns the state.
func (ps *Persister) ReadRaftState() []byte {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	return ps.raftstate
}

// RaftStateSize returns the size of the state.
func (ps *Persister) RaftStateSize() int {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	return len(ps.raftstate)
}

// SaveStateAndSnapshot saves both Raft state and K/V snapshot as a single atomic action,
// to help avoid them getting out of sync.
func (ps *Persister) SaveStateAndSnapshot(state []byte, snapshot []byte) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ps.raftstate = state
	ps.snapshot = snapshot
}

// ReadSnapshot returns the snapshot.
func (ps *Persister) ReadSnapshot() []byte {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	return ps.snapshot
}

// SnapshotSize returns the size of the snapshot.
func (ps *Persister) SnapshotSize() int {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	return len(ps.snapshot)
}
