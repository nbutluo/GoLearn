package main

import "fmt"

func main () {
	var command = "gooutside"

	if(command == "gohost") {
		fmt.Println("这是正确的区间")
	} else if (command == "gooutside") {
		fmt.Println("这是outside")
	}
}