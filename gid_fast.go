//go:build !llgo && go1.23 && (386 || amd64 || arm || arm64 || loong64 || mips || mipsle || mips64 || mips64le || ppc64 || ppc64le || riscv64 || s390x || wasm)
package gid

// Get returns the current goroutine ID through the runtime g register.
func Get() uint64 { return fastGet() }

func fastGet() uint64
