package porcupine

import "time"

// CheckOperations checks the operations.
func CheckOperations(model Model, history []Operation) bool {
	res, _ := checkOperations(model, history, false, 0)
	return res == Ok
}

// CheckOperationsTimeout checks the operations with timeout.
// timeout = 0 means no timeout
// if this operation times out, then a false positive is possible
func CheckOperationsTimeout(model Model, history []Operation, timeout time.Duration) CheckResult {
	res, _ := checkOperations(model, history, false, timeout)
	return res
}

// CheckOperationsVerbose checks the operations.
// timeout = 0 means no timeout
// if this operation times out, then a false positive is possible
func CheckOperationsVerbose(model Model, history []Operation, timeout time.Duration) (CheckResult, linearizationInfo) {
	return checkOperations(model, history, true, timeout)
}

// CheckEvents checks the events.
func CheckEvents(model Model, history []Event) bool {
	res, _ := checkEvents(model, history, false, 0)
	return res == Ok
}

// CheckEventsTimeout checks the events with timeout.
// timeout = 0 means no timeout
// if this operation times out, then a false positive is possible
func CheckEventsTimeout(model Model, history []Event, timeout time.Duration) CheckResult {
	res, _ := checkEvents(model, history, false, timeout)
	return res
}

// CheckEventsVerbose checks the events.
// timeout = 0 means no timeout
// if this operation times out, then a false positive is possible
func CheckEventsVerbose(model Model, history []Event, timeout time.Duration) (CheckResult, linearizationInfo) {
	return checkEvents(model, history, true, timeout)
}
