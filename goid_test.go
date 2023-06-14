package gid

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"testing"
)

func extractGID(s []byte) int64 {
	s = s[len("goroutine "):]
	s = s[:bytes.IndexByte(s, ' ')]
	gid, _ := strconv.ParseInt(string(s), 10, 64)
	return gid
}

func getGoid() int64 {
	if runtime.GOOS == "js" && runtime.GOARCH != "wasm" {
		return Get()
	}
	var buf [64]byte
	return extractGID(buf[:runtime.Stack(buf[:], false)])
}

func TestGet(t *testing.T) {
	ch := make(chan *string, 100)
	for i := 0; i < cap(ch); i++ {
		go func(i int) {
			goid := Get()
			expected := getGoid()
			if goid == expected {
				ch <- nil
				return
			}
			s := fmt.Sprintf("Expected %d, but got %d", expected, goid)
			ch <- &s
		}(i)
	}

	for i := 0; i < cap(ch); i++ {
		val := <-ch
		if val != nil {
			t.Fatal(*val)
		}
	}
}

func BenchmarkGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get()
	}
}
