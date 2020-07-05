package shardkv

//
// Sharded key/value server.
// Lots of replica groups, each running op-at-a-time paxos.
// Shardmaster decides which group serves each shard.
// Shardmaster may change shard assignment from time to time.
//
// You will have to modify these definitions.
//

// Errors.
const (
	OK             = "OK"
	ErrNoKey       = "ErrNoKey"
	ErrWrongGroup  = "ErrWrongGroup"
	ErrWrongLeader = "ErrWrongLeader"
)

// Err represents Error string.
type Err string

// PutAppendArgs represents Put or Append args.
type PutAppendArgs struct {
	// You'll have to add definitions here.
	Key   string
	Value string
	Op    string // "Put" or "Append"
	// You'll have to add definitions here.
	// Field names must start with capital letters,
	// otherwise RPC will break.
}

// PutAppendReply represents Put or Append reply.
type PutAppendReply struct {
	Err Err
}

// GetArgs represents args.
type GetArgs struct {
	Key string
	// You'll have to add definitions here.
}

// GetReply represents reply.
type GetReply struct {
	Err   Err
	Value string
}
