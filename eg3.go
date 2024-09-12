package main

func main() {
	var command = "this is outside world"
	var exit = strings.Contains(command,"outside")

	fmt.Println(exit)
}