package kvraft

// Errors for Raft
const (
	OK             = "OK"
	ErrNoKey       = "ErrNoKey"
	ErrWrongLeader = "ErrWrongLeader"
)

// Err is an error string.
type Err string

// PutAppendArgs represents Put or Append
type PutAppendArgs struct {
	Key   string
	Value string
	Op    string // "Put" or "Append"
	// You'll have to add definitions here.
	// Field names must start with capital letters,
	// otherwise RPC will break.
}

// PutAppendReply represents a replied error.
type PutAppendReply struct {
	Err Err
}

// GetArgs represents args.
type GetArgs struct {
	Key string
	// You'll have to add definitions here.
}

// GetReply represents a reply.
type GetReply struct {
	Err   Err
	Value string
}
