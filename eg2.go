package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("这是输出")

	command := "walk outside"

	exit := strings.Contains(command,"outside")

	fmt.Println(exit)
}