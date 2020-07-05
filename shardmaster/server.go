package shardmaster

import (
	"sync"

	"6.824/labgob"
	"6.824/labrpc"
	"6.824/raft"
)

// ShardMaster represents a shard master.
type ShardMaster struct {
	mu      sync.Mutex
	me      int
	rf      *raft.Raft
	applyCh chan raft.ApplyMsg

	// Your data here.

	configs []Config // indexed by config num
}

// Op represents an op.
type Op struct {
	// Your data here.
}

// Join func.
func (sm *ShardMaster) Join(args *JoinArgs, reply *JoinReply) {
	// Your code here.
}

// Leave func.
func (sm *ShardMaster) Leave(args *LeaveArgs, reply *LeaveReply) {
	// Your code here.
}

// Move func.
func (sm *ShardMaster) Move(args *MoveArgs, reply *MoveReply) {
	// Your code here.
}

// Query func.
func (sm *ShardMaster) Query(args *QueryArgs, reply *QueryReply) {
	// Your code here.
}

// Kill is called by the tester when a ShardMaster instance won't
// be needed again. you are not required to do anything
// in Kill(), but it might be convenient to (for example)
// turn off debug output from this instance.
func (sm *ShardMaster) Kill() {
	sm.rf.Kill()
	// Your code here, if desired.
}

// Raft is needed by shardkv tester
func (sm *ShardMaster) Raft() *raft.Raft {
	return sm.rf
}

// StartServer starts the server.
// servers[] contains the ports of the set of
// servers that will cooperate via Paxos to
// form the fault-tolerant shardmaster service.
// me is the index of the current server in servers[].
func StartServer(servers []*labrpc.ClientEnd, me int, persister *raft.Persister) *ShardMaster {
	sm := new(ShardMaster)
	sm.me = me

	sm.configs = make([]Config, 1)
	sm.configs[0].Groups = map[int][]string{}

	labgob.Register(Op{})
	sm.applyCh = make(chan raft.ApplyMsg)
	sm.rf = raft.Make(servers, me, persister, sm.applyCh)

	// Your code here.

	return sm
}
