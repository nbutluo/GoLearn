package main

import "fmt"

func main() {
	var fileld = "name"

	switch {

	case fileld == "age":
		fmt.Println("这是年龄")
	
		// 执行下一个case的body部分
		fallthrough
	case fileld == "sex":
		fmt.Println("这是性别")
	case fileld =="head":
		fmt.Println("这是头部˜")
	case fileld == "name":
		fmt.Println("这是名称")
	}
}