
all: libbim.so interactTrial

libbim.so:
	make -C CCode/ clean
	make -C CCode/

interactTrial: libbim.so
	go build interactTrial.go

clean:
	make -C CCode/ clean
	rm -f libbim.so interactTrial