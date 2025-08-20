package main

import "fmt"
import "unsafe"

type Good struct {
	B int64
	A int32
}

func main() {
	fmt.Println("Size of Good struct:", unsafe.Sizeof(Good{}))
}
