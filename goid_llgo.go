//go:build llgo
// +build llgo

package gid

import (
	"unsafe"
)

//go:linkname self C.pthread_self
func self() unsafe.Pointer

func Get() uint64 {
	return uint64(uintptr(self()))
}
