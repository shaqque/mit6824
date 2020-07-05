package shardmaster

//
// Master shard server: assigns shards to replication groups.
//
// RPC interface:
// Join(servers) -- add a set of groups (gid -> server-list mapping).
// Leave(gids) -- delete a set of groups.
// Move(shard, gid) -- hand off one shard from current owner to gid.
// Query(num) -> fetch Config # num, or latest config if num==-1.
//
// A Config (configuration) describes a set of replica groups, and the
// replica group responsible for each shard. Configs are numbered. Config
// #0 is the initial configuration, with no groups and all shards
// assigned to group 0 (the invalid group).
//
// You will need to add fields to the RPC argument structs.
//

// NShards represents the number of shards.
const NShards = 10

// Config is an assignment of shards to groups.
// Please don't change this.
type Config struct {
	Num    int              // config number
	Shards [NShards]int     // shard -> gid
	Groups map[int][]string // gid -> servers[]
}

// Ok means ok.
const (
	OK = "OK"
)

// Err is an error string.
type Err string

// JoinArgs struct.
type JoinArgs struct {
	Servers map[int][]string // new GID -> servers mappings
}

// JoinReply struct.
type JoinReply struct {
	WrongLeader bool
	Err         Err
}

// LeaveArgs struct.
type LeaveArgs struct {
	GIDs []int
}

// LeaveReply struct.
type LeaveReply struct {
	WrongLeader bool
	Err         Err
}

// MoveArgs struct.
type MoveArgs struct {
	Shard int
	GID   int
}

// MoveReply struct.
type MoveReply struct {
	WrongLeader bool
	Err         Err
}

// QueryArgs struct.
type QueryArgs struct {
	Num int // desired config number
}

// QueryReply struct.
type QueryReply struct {
	WrongLeader bool
	Err         Err
	Config      Config
}
