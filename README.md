[![Go Report Card](https://goreportcard.com/badge/github.com/0xmjk/go-tsx-rtm)](https://goreportcard.com/report/github.com/0xmjk/go-tsx-rtm) 
[![GoDoc](https://godoc.org/github.com/0xmjk/go-tsx-rtm?status.svg)](https://godoc.org/github.com/0xmjk/go-tsx-rtm) 
[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/0xmjk/go-tsx-rtm/blob/master/LICENSE) 

*Transactional Synchronization Extensions (TSX)* is an extension to the Intel x86 ISA adding hardware transactional memory support.

*Restricted Transactional Memory (RTM)* uses new `XBEGIN` and `XEND` instructions to mark start and end of a critical section.
The processor would treat this section as an atomic transaction.

This package exposes these primitives to the developer.

This code doesn't check if the CPU supports RTM at all, so it is necessary to do it first, e.g.:

```go
import (
  "github.com/intel-go/cpuid"
)

func CpuHasRTM() bool {
	return cpuid.HasExtendedFeature(cpuid.RTM)
}
```

Caveat: Golang will not inline assembly at the moment so using this might be slower than similar GCC implementation.

Have a look at Intel's ["Intrinsics for Restricted Transactional Memory Operations"](https://software.intel.com/en-us/node/524024) for more background.
