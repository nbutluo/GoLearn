package main

import "fmt"

func main() {
	const (
		a, b = iota + 1, iota + 2  // iota = 0 a=1 b=2

		c,d  // iota =1 c=2 d=3

		/**
		 iota 在第一行表达式确定，后续 iota 始终自增 1
		 
		*/
		// 表达式即时发生变动，iota 值依旧是旧的
		e,f  = iota*10,iota*20  // iota=2 e=20 f=40
	)

	// ee = iota * 20   iota 仅限在 const()体内使用

	fmt.Println(a, b,c,d,e,f)
}
