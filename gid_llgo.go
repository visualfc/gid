//go:build llgo
// +build llgo

package gid

import _ "unsafe"

//go:linkname goid github.com/xgo-dev/llgo/runtime/internal/runtime.goid
func goid() uint64

func Get() uint64 {
	return goid()
}
