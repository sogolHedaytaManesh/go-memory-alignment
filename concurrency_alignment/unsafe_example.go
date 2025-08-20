package main

import "fmt"
import "unsafe"

type Weird struct {
	A int8
	B int64
}

func main() {
	fmt.Println("Size of Weird struct:", unsafe.Sizeof(Weird{}))
	fmt.Println("Alignment of Weird struct:", unsafe.Alignof(Weird{}))
}
