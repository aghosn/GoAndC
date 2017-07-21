package main

// #cgo CFLAGS: -I./CCode
// #cgo LDFLAGS: -L. -lbim
//
// #include <bim.h>
import "C"

func main() {
	C.Bim_test()
}
