package gpo

import "fmt"

func Message(name string) string {
	result := fmt.Sprintf("Have a nice day %s", name)
	return result
}
