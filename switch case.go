package main

import "fmt"

func main() {
	area := "hangzhou"

	switch {
	case area == "beijing": 
		fmt.Println("这里是北京")		
		
	case area == "shanghai":
		fmt.Println("这里是上海")

	case area == "hangzhou":
		fmt.Println("这里是杭州")
	}
}