# Calls from Go to C and the other way around.

In this folder, a Go main function calls a C function that itself calls a go function.
Everything gets automatically built by the go compiler (one stage only).
The only limitation is that the entry point must be inside the Go part of the program.
