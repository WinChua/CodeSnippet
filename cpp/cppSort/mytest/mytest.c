#include <stdio.h>
#include <stdlib.h>

int main(int argc, char * argv[]) {
    if ( argc < 3 ) {
        printf("Usage %s number size", argv[0]);
        return 0;
    }
    int times = atoi(argv[1]);
    int i = 0;
    for(; i < times; i++) {
        free(malloc(4096 * atoi(argv[2])));
    }
}
