package main

import "fmt"
import "unsafe"

func main() {
	var a = 10
	a, b := 10, 10
	fmt.Println("A=", a, "B=", b)
	fmt.Printf("A=%T", a)
	fmt.Println("SIZE===", unsafe.Sizeof(a))
}
