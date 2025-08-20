package main

import "fmt"

func escapeExample() *int {
	x := 7
	return &x // escapes to heap
}

func main() {
	p := escapeExample()
	fmt.Println("Escaped variable:", *p)
}
