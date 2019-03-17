#ifndef H_BYTES
#define H_BYTES

#include <stdio.h>

void HelloByte();

typedef unsigned char byte;

void _printByte(byte*, size_t);
void show_vptr(long *, bool call=true);

template <class T>
void printByte(const T &b) {
    byte * start = (byte*)&b;
    size_t size = sizeof(T);
    _printByte(start, size);
}

#endif
