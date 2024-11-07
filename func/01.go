package main

import "fmt"

func main() {

	fmt.Println("aaas")

	c := foo1("beijing",1)
	fmt.Println("c =", c)

}

// 返回值 int
func foo1 (a string,b int) int {
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	c := 100
	return c
}