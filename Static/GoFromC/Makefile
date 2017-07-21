
all: libfoo.a foo

libfoo.a:
	go build -buildmode=c-archive -o libfoo.a foo.go

foo: libfoo.a foo.o
	gcc -L./ -Wall -o foo foo.c -l:libfoo.a -lpthread

clean:
	rm -f libfoo.a libfoo.h foo