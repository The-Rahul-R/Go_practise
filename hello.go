package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter Your First Name: ")

	var first string

	fmt.Scanln(&first)

	fmt.Println("Enter Second Last Num: ")

	var second string

	fmt.Scanln(&second)

	fmt.Printf("hello %s", first+" "+second)
}
