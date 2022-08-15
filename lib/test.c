#include <stdio.h>
#include "add.h"

int main(int argc, char *argv[])
{
    char* aa = "Hello CGO, ";
    printf("%s\n", Add(aa, 8));
    return 0;
}
