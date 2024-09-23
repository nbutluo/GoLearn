package main

import "fmt"

// dd := 2000
var ee int = 1000
func main(){
	//方法 1 声明一个变量 默认为 0
	var a int
	fmt.Println("a=",a);
	// 查看类型
	fmt.Printf("type of a %T\n",a)

	// 方法2，声明一个变量，并且指定默认值
	var b int = 200
	fmt.Println("b=",b)
	fmt.Printf("type of b %T\n",b)

	// 方法 3.省去数据类型，根据值自动匹配当前类型
	var c = 300
	fmt.Println("c",c)
	fmt.Printf("type of %T\n",c)

	// 方法 4 ， 使用：=
	 d := 400
	fmt.Println(d)
	fmt.Printf("type of %T\n",d)

	// 打印全局变量
	fmt.Println("全局变量ee",ee)
	// fmt.Println("全局变量 dd",dd) :=不可以定义全局的变量

	// 多个变量定义
	var rr,tt = 100,200
	fmt.Println(rr,tt)

	// 多行定义多个
	var  (
		yy = 100
		uu = "varchar"
	)
	fmt.Println(yy,uu)

	var gg,jj  = 100,"abc"
	fmt.Println(gg,jj)

}