shared:
	rm -f libbim.so
	gcc -c -Wall -fPIC bim.c
	gcc -shared -o libbim.so bim.o
	rm -f bim.o
	cp libbim.so ../

.PHONY: clean

clean:
	@rm -f bim.o libbim.so
