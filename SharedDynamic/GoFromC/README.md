# Generating a shared library from Go code

This folder generates a libfoo.so from the Go code and links it to the C program.
Don't forget to update LD_LIBRARY_PATH and export the variable, otherwise you'll have a runtime error, i.e.,:

~~~~ 
LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/path/to/libfoo.so 
export LD_LIBRARY_PATH
~~~~
