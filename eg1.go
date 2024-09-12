package main

import "fmt"

func main() {
 fmt.Print("hello world")
 beyond()
}

func beyond() {
	var x int
	x =1 
	y:=2

	sum,prod := calcFunc(x,y)

	fmt.Print("sum=",sum,"prod=",prod)
}


func calcFunc(x,y int) (sum,prod int) {
	return x+y,x*y
}