package main

import "fmt"

func main() {
	city := "xian"

	switch {
	case city=="tianshui":
		fmt.Println("天水市")
	case city=="xian":
		fmt.Println("西安市")
		fallthrough
	case city=="naning":
		fmt.Println("南京市")

	}
}