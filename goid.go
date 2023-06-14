//go:build !js || (js && wasm)
// +build !js js,wasm

package gid

import "github.com/timandy/routine"

func Get() int64 {
	return routine.Goid()
}
