#include <stdio.h>
#include <stdlib.h>
#include <dlfcn.h>

typedef void (*Func)(void);

int main() {
    void * handle = dlopen("libhello.so", RTLD_NOW);
    if(handle == NULL) {
        printf("Error\n");
        exit(1);
    }
    Func helloworld = NULL;
    helloworld = (Func)dlsym(handle, "helloworld");
    if(helloworld == NULL) {
        printf("helloworld Error\n");
        exit(1);
    }
    helloworld();
    dlclose(handle);
}
