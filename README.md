# gid
golang routine id

- Go use <https://github.com/timandy/routine>
- LLGo use C.pthread_self

```
import "github.com/visualfc/gid"

var id uint64 = gid.Get()
```
