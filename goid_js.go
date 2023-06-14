//go:build js && !wasm
// +build js,!wasm

package gid

import (
	"sync/atomic"

	"github.com/gopherjs/gopherjs/js"
)

var (
	index int64 = 0
)

const (
	goid_key = "__goid__"
)

func Get() int64 {
	obj := js.Global.Get("$curGoroutine")
	id := obj.Get(goid_key)
	if id == js.Undefined {
		n := atomic.AddInt64(&index, 1)
		obj.Set(goid_key, n)
		return n
	}
	return id.Int64()
}
