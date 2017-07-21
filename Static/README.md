# How to interact between Go and C without creating shared libraries.

This folder contains C and Go code that interact statically.

The **go** compiler is able to compile everything in one stage as long as the program's entry point is the Go code (simple 'go build' works).
In the *Static/CFromGo* folder, the code sample calls C functions from Go, and Go functions from C.

To have an entry point in the C code instead, I had to generate a static archive from the Go code and link it to the C application (see *Static/GoFromC*).

**TODO**: Find if it is possible to import Go code inside C, with the program entry point in C.

