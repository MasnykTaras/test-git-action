package client

import "fmt"

type Test struct {
	Client
}
type Test2 struct {
	Client
}
type Client interface {
	Test()
	GetSum(int, int) (int, error)
}
type ClientObject struct {
	provider Client
}

func NewClient(client Client) *ClientObject {
	return &ClientObject{provider: client}
}
func (input ClientObject) GetSum(num1, num2 int) (int, error) {
	input.provider.Test()
	return num1 + num2, nil
}
func (test *Test) Test() {
	fmt.Println("HEllo")
}
func (test *Test2) Test() {
	fmt.Println("Goodbuy")
}
