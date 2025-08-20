package main

import (
	"fmt"
	"sync/atomic"
)

type Counter struct {
	Value int64
}

func main() {
	var c Counter
	atomic.AddInt64(&c.Value, 1)
	fmt.Println("Atomic counter:", c.Value)
}
