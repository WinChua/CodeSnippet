#include "vptr.h"
#include "bytes.h"

void _printByte(byte * start, size_t size) {
    int i = 0;
    printf("     ");
    for(;i < 8; i++) {
        printf("%02d ", i);
    }
    printf("\n");
    for(i=0; i<size; i++, start++){
        if(i % 8 == 0) 
            printf("0x%02x ", i);
        printf("%02x ", *start);
        if((i+1)%8 == 0)
            printf("\n");
    }
    printf("\n");
}
void show_vptr(long *vptr, bool call) {
    int i;
    for(i=0; (Func)vptr[i] != NULL; i++) {
        Func pFunc = (Func)vptr[i];
        //if(vptr[i] < 0) continue;
        //printf("%d: %p, %d\n", i, pFunc, pFunc);
        if(vptr[i] < 0) break;
        printf("[%d]->%p:", i, pFunc);
        if(call) pFunc();
        else printf("\n");
    }
    SHOW_ADDR(*vptr);
    _printByte((byte*)vptr, (1+i)*8);
}
