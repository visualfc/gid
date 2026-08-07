//go:build wasm && !llgo
// +build wasm,!llgo

package gid

import (
	_ "github.com/timandy/routine"
	"unsafe"
)

// offsetGoid is discovered by routine at startup, so the fast path does not
// depend on a fixed runtime.g layout.
//
//go:linkname offsetGoid github.com/timandy/routine.offsetGoid
var offsetGoid uintptr

func getg() unsafe.Pointer

// Get returns the id of the current goroutine.
func Get() uint64 {
	return *(*uint64)(unsafe.Pointer(uintptr(getg()) + offsetGoid))
}
