package main

import "fmt"

type CacheLine struct {
	X int64
	_ [56]byte // padding to separate cache lines
	Y int64
}

func main() {
	var c CacheLine
	c.X = 10
	c.Y = 20
	fmt.Println("Cache line struct:", c.X, c.Y)
}
