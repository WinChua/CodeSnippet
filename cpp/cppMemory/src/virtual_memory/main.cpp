#include "bytes.h"
#include "vptr.h"
#include <iostream>

class Base {
    public:
        virtual void hello() {
            std::cout << "Hello" << std::endl;
        }
};

class Inheri : public Base {
    public:
        void hello() {
            std::cout << "World" << std::endl;
        }
};
/*
 *  +----------------+
 *  |    Base        |
 *  +----------------+
 *  |    Base's vptr |
 *  +----------------+
 *          ^
 *          |
 *  +----------------+
 *  |    Inheri      |
 *  +----------------+
 *  |    Inheri's vpt|
 *  +----------------+
 *
 */
// typdef void (*Func)(void);
int main() {
    Inheri a;
    SHOW_ADDR(a);
    printByte(a);
    long *vptr = GET_VPTR(a);
    Func f = (Func)vptr[0];
    f();
    Base b;
    SHOW_ADDR(b);
    printByte(b);
    *(long*)&a = *(long*)&b;
    // replace the contenct of a's vptr to b's vptr
    // so that the call of &a->hello() will call the hello
    // from Base

    vptr = GET_VPTR(a);
    f = (Func)vptr[0];
    printf("--------------------------\n");
    //call hello without pointer will invoke the hello from a;
    a.hello();
    (&a)->hello();
    Base *p = &a;
    //call hello by pointer will invoke the hello from b;
    p->hello();
    printf("--------------------------\n");
    SHOW_ADDR(a);
    printByte(a);
    f();
}
