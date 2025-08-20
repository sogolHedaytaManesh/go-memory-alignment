package main

import (
	"fmt"
	"unsafe"
)

type Example struct {
	A int32
	B int64
}

func main() {
	fmt.Println("Size:", unsafe.Sizeof(Example{}))
	fmt.Println("Alignment:", unsafe.Alignof(Example{}))
}
