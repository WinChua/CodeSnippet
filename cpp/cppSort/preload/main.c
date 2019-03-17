#include <stdlib.h>

int main() {
    free(malloc(sizeof(int)));
}
