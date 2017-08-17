package rtm

// test without abort
// test with abort
// test _xtest behaviour

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXTestOutsideTransaction(t *testing.T) {
	actualResult := TxTest()
	assert.Equal(t, uint8(0), actualResult)
}

func TestXTestInsideTransaction(t *testing.T) {
	TxBegin()
	actualResult := TxTest()
	TxEnd()
	assert.Equal(t, uint8(1), actualResult)
}

func TestTransactionAbort(t *testing.T) {
	a := 0
	if status := TxBegin(); status == TxBeginStarted {
		a += 6
		TxAbort()
		TxEnd()
		t.Fail() // shouldn't get here
	} else {
		assert.Equal(t, uint32(0xf0000001), status)
		assert.Equal(t, uint32(0x1), status&TxAbortExplicit)
		assert.Equal(t, uint8(0xf0), GetImm(status))
		assert.Equal(t, 0, a)
	}
}

func TestTransactionCommit(t *testing.T) {
	a := 0
	if status := TxBegin(); status == TxBeginStarted {
		a += 6
		TxEnd()
	} else {
		t.Fail() // shouldn't get here
	}
	assert.Equal(t, 6, a)
}
