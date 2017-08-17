// Package rtm provides primitives for Intel's Restricted Transactional Memory Operations
package rtm

// TxBegin marks the start of transaction. It will return a status code
// different to TxBeginStarted when the transaction fails.
func TxBegin() (status uint32)

// TxAbort aborts transaction
func TxAbort()

// TxEnd marks the end of transaction
func TxEnd()

// TxTest returns uint8(1) if the processor is executing a transactional region.
func TxTest() (status uint8)

// GetImm returns uint8 value from the uint32 status returned by TxBegin
func GetImm(status uint32) uint8 {
	return uint8(((status) >> 24) & 0xff)
}

const (
	// TxBeginStarted is returned by TxBegin() when transaction is started
	TxBeginStarted uint32 = ^uint32(0)
	// TxAbortExplicit bit is set if abort caused by explicit abort instruction.
	TxAbortExplicit uint32 = (1 << 0)
	// TxAbortRetry bit is set if the transaction may succeed on a retry
	TxAbortRetry uint32 = (1 << 1)
	// TxAbortConflict bit is set if another logical processor triggered a
	// conflict with a memory address that was part of the transaction
	TxAbortConflict uint32 = (1 << 2)
	// TxAbortCapacity bit is set if RTM buffer overflowed
	TxAbortCapacity uint32 = (1 << 3)
	// TxAbortDebug is set if debug breakpoint triggered
	TxAbortDebug uint32 = (1 << 4)
	// TxAbortNested is set if abort occurred in a nested transaction
	TxAbortNested uint32 = (1 << 5)
)
