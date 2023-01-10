package main

import (
	"fmt"
	client "test/client"
)

func main() {
	var clien = &client.Test{}
	var clien2 = &client.Test2{}
	object := client.NewClient(clien)
	object2 := client.NewClient(clien2)
	num1 := 1
	num2 := 2

	res, err := object.GetSum(num1, num2)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Sum of %d and %d is  %d\n", num1, num2, res)
	}
	res2, err2 := object2.GetSum(num1, num2)
	if err2 != nil {
		panic(err2)
	} else {
		fmt.Printf("Sum of %d and %d is  %d\n", num1, num2, res2)
	}
}
