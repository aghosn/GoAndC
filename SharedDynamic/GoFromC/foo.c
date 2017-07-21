#include <stdio.h>
#include "libfoo.h"

extern void MyFunction(void);

int main(int argc, char const *argv[])
{
	MyFunction();
	printf("Finished!\n");
	return 0;
}