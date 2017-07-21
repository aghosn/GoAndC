
all: libfoo.so foo

libfoo.so:
	go build -buildmode=c-shared -o libfoo.so foo.go

foo: libfoo.so foo.o
	gcc -L./ -Wall -o foo foo.c -lfoo

clean:
	rm -f libfoo.so foo libfoo.h