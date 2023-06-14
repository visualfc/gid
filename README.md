# gid
golang routine id

- GopherJS use `$curGoroutine`
- Other GOOS/GOARCH use <https://github.com/timandy/routine>

```
import "github.com/visualfc/gid"

var id int64 = gid.Get()
```
