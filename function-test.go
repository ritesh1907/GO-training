package main

import "fmt"

func functiontest(aa int, bb int) int {
	var c int
	c = aa + bb
	return (c)
}
func main() {
	var a int
	a = 10
	b := 11
	fmt.Println(" A=", a)
	c := functiontest(a, b)
	fmt.Println("Ans==", c)
}
