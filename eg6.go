package main

import "fmt"

func main() {
	a  := true
	b := false

	// if !b | a {
	// 	fmt.Println("this is true")
	// }

	if b || !a {
		fmt.Println("this is",false)
	}
}