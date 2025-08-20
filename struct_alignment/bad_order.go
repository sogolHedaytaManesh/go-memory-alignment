package main

import (
	"fmt"
	"unsafe"
)

type Bad struct {
	A int32
	B int64
}

func main() {
	fmt.Println("Size of Bad struct:", unsafe.Sizeof(Bad{}))
}
