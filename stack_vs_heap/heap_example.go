package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := &Person{Name: "Alice", Age: 30} // allocated on heap
	fmt.Println("Heap allocated struct:", p)
}
