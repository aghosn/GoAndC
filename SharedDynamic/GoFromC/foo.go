package main

import "fmt"
import "C"

//export MyFunction
func MyFunction() {
	fmt.Println("AGoFunction():")
}

func main() {}
