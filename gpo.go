package gpo

import "fmt"

func Message(name string) string {
	result := fmt.Sprintf("Have a nice day %s", name)
	return result
}

func TestMessage(name string) string {
	result := fmt.Sprintf("Have a nice night %s", name)
	return result
}

func TestNewMessage(name string) string {
	result := fmt.Sprintf("Have a nice week %s", name)
	return result
}
