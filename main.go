package gpo

import "fmt"

type Resultado struct {
	result string
}

func (r *Resultado) Message(name string) string {
	result := fmt.Sprintf("Have a nice day %s", name)
	return result
}

func (r *Resultado) TestMessage(name string) string {
	result := fmt.Sprintf("Have a nice night %s", name)
	return result
}

func (r *Resultado) TestNewMessage(name string) string {
	result := fmt.Sprintf("Have a nice week %s", name)
	return result
}
