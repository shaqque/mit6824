package raft

import "log"

// Debug is used for debugging.
const Debug = 0

// DPrintf prints a log for debugging.
func DPrintf(format string, a ...interface{}) (n int, err error) {
	if Debug > 0 {
		log.Printf(format, a...)
	}
	return
}
