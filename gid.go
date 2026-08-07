//go:build !wasm && !llgo
// +build !wasm,!llgo

package gid

import "github.com/timandy/routine"

func Get() uint64 {
	return routine.Goid()
}
