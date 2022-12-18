package main

import "fmt"

func main() {
	fmt.Println("Hello world!!")
	CreateAMessage("Lucas")
}

func CreateAMessage(name string) string {
	result := fmt.Sprintf("Have a nice day %s", name)
	return result
}
