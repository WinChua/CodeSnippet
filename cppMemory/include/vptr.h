#ifndef H_VPTR
#define H_VPTR

#include <stdio.h>

#define ADD_QUO(a) #a
#define CONCAT_W(a, b) ADD_QUO(a##b)
#define GET_VPTR(obj) ((long*)*(long*)&obj)
#define SHOW_MSG_ADDR(msg, addr) printf("%s: %p\n", (msg), (&addr)) 
#define SHOW_ADDR(addr) SHOW_MSG_ADDR("The addr of "#addr, addr)

typedef void (*Func)(void);

#endif
