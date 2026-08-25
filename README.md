# gid
golang routine id

- Go uses assembly fast paths on supported architectures.
- LLGo reads `goid` from the runtime goroutine returned by LLGo's `getg`.

Supported Go versions: 1.23.x, 1.24.x, 1.25.x, 1.26.x, and 1.27.x.
Supported architectures: 386, amd64, arm, arm64, loong64, mips, mipsle,
mips64, mips64le, ppc64, ppc64le, riscv64, s390x, and wasm.

Go 1.23 and Go 1.24 use the `go123` assembly fast paths. Go 1.25, Go 1.26,
and Go 1.27 use the `go125` assembly fast paths. Earlier Go versions use the
portable `GetSlow` fallback.

```
import "github.com/visualfc/gid"

var id uint64 = gid.Get()

// Portable fallback when runtime layout compatibility is more important
// than performance.
var slowID uint64 = gid.GetSlow()
```
