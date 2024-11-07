package main

import "fmt"


func main(){
	foo1("hahaa",200)

	ret1,ret2 := foo2("yuque",4444)

	fmt.Println("ret1 =",ret1,"ret2=",ret2)

	ret3,ret4 := foo3("heheh",789)
	fmt.Println(ret3,ret4)
}

func foo1(a string , b int) int {
	fmt.Println(a)
	fmt.Println(b)

	c := 100

	return c
}

func foo2(a string, b int) (int,int) {
	
	return 666,777
}

func foo3(c string,d int) ( r1 string,r2 int) {
	// return "shanghai",111

	r1 = "shanghai"
	r2 = 120

	return
}