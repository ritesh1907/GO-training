package main

import "fmt"

func sum(a int, b int) int {
	var c int
	c = a + b
	return c
}

func main() {
	var a, b int
	a = 10
	b = 111
	fmt.Println(" ANS ==> ", sum(a, b))
}
