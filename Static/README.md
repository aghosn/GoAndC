#How to interact between Go and C without creating shared libraries.

In this folder we have example of Go code calling C functions and C functions executing Go code.
For Go calling C, the go compiler enables to compile everything in one stage with a simple 'go buil'.
The entry point of the program needs to be in the Go code.

Unfortunately, I didn't find how to make the entry point in C and instead provide an example where the go compiler generates a static archive that is passed as an argument to gcc when building the C code.

