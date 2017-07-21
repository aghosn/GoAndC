package main

import "fmt"

// int main2();
import "C"

//export MyFunction
func MyFunction() {
	fmt.Println("Hello from go.")
}

func main() {
	C.main2()
}
