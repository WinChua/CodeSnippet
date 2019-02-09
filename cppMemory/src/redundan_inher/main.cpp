#include "bytes.h"
#include "vptr.h"

#include <iostream>

class B {
    public:
        int ib;
        char cb;
        B():ib(1), cb(1) {}
        virtual void f() {
            std::cout << "B::f()" << std::endl;
        }
        virtual void Bf() {
            std::cout << "B::Bf()" << std::endl;
        }
};

class B1: public B {
    public:
        int ib1;
        char cb1;
        B1():ib1(2), cb1(2) {}
        virtual void f1() {
            std::cout << "B1::f1()" << std::endl;
        }
        virtual void Bf1() {
            std::cout << "B1::Bf1()" << std::endl;
        }
};
class B2: public B {
    public:
        int ib2;
        char cb2;
        B2():ib2(3), cb2(3) {}
        virtual void f2() {
            std::cout << "B2::f2()" << std::endl;
        }
        virtual void Bf2() {
            std::cout << "B2::Bf2()" << std::endl;
        }
};
class D :public B1, public B2 {
    public:
        int id;
        char cd;
        D():id(4), cd(4){}
        virtual void f() {
            std::cout << "D::f()" << std::endl;
        }
        virtual void f1() {
            std::cout << "D::f1()" << std::endl;
        }
        virtual void f2() {
            std::cout << "D::f2()" << std::endl;
        }
        virtual void Df() {
            std::cout << "D::Df()" << std::endl;
        }
};


int main() {
    D d;
    long *vptr = GET_VPTR(d);
    show_vptr(vptr);
    printByte(d);
    //d.ib = 42; //compile error: request of member ib is ambiguous;
    d.B1::ib = 5;
    d.B2::ib = 6;
    printByte(d);
    long *vptr2 = (long*)*(long*)((char*)&d+sizeof(B1));
    show_vptr(vptr2);
    long *vptr3 = (long*)*(long*)((char*)&d + sizeof(B1) + sizeof(B2));
    show_vptr(vptr3);
    printf("World\n");
}
