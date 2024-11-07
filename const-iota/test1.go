package main

import "fmt"
func main() {
	
	// 定义一个常量(常量具有只读属性)
	const length = 100

	// length = 200 常量一被定义，就只有只读属性了
	fmt.Println(length)

	// 定义多个常量 const()
	const (
		BEIJING = 1
		SHANGHAI =2
	)

	fmt.Println(BEIJING,SHANGHAI)

	
}