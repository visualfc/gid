package gid

import (
	"bytes"
	"runtime"
	"strconv"
)

// GetSlow returns the current goroutine ID using the portable runtime.Stack
// fallback. It is slower and allocates, but does not depend on runtime g
// structure layout.
func GetSlow() uint64 {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	if n == 0 {
		return 0
	}
	line := buf[:n]
	const prefix = "goroutine "
	if !bytes.HasPrefix(line, []byte(prefix)) {
		return 0
	}
	line = line[len(prefix):]
	end := bytes.IndexByte(line, ' ')
	if end < 0 {
		return 0
	}
	id, _ := strconv.ParseUint(string(line[:end]), 10, 64)
	return id
}
