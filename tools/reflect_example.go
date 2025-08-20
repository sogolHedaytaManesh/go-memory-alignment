package main

import (
	"fmt"
	"reflect"
)

type Demo struct {
	A int
	B string
}

func main() {
	t := reflect.TypeOf(Demo{})
	fmt.Println("Type:", t.Name())
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("Field %s: type=%s, offset=%d\n", f.Name, f.Type, f.Offset)
	}
}
