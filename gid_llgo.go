//go:build llgo
// +build llgo

package gid

import (
	"unsafe"
)

//go:linkname getg github.com/goplus/llgo/runtime/internal/runtime.getg
func getg() *g

func Get() uint64 {
	return getg().goid
}

//llgo:type C
type goroutineFunc func(unsafe.Pointer) unsafe.Pointer

type g struct {
	defer_ unsafe.Pointer
	panic_ unsafe.Pointer
	m      unsafe.Pointer

	atomicstatus uint32
	goid         uint64
	parentGoid   uint64

	startfn  goroutineFunc
	startarg unsafe.Pointer

	context *unsafe.Pointer

	goexit       bool
	isMain       bool
	paniconfault bool
}
