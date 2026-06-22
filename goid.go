//go:build (!js || (js && wasm)) && !llgo
// +build !js js,wasm
// +build !llgo

package gid

import "github.com/timandy/routine"

func Get() uint64 {
	return routine.Goid()
}
