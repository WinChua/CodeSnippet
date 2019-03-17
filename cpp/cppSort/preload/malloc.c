#include <stdio.h>
void * malloc(size_t size) {
    printf("malloc");
    return NULL;
}

void free(void *p) {
    printf("free");
    return;
}
