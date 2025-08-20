package main

import "fmt"
import "unsafe"

type Inner struct {
	A int32
	B int32
}

type Outer struct {
	I Inner
	C int64
}

func main() {
	fmt.Println("Size of Outer struct:", unsafe.Sizeof(Outer{}))
}
