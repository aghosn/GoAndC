package main

import "C"
import "fmt"

//export Foo
func Foo() {
	fmt.Println("Inside the Static Go Foo function.")
}

func main() {}
