#include "vptr.h"
#include "bytes.h"
#include <iostream>

class Base1 {
    public:
        int ibase1;
        Base1():ibase1(1) {}
        virtual void f() {
            std::cout << "Base1::f()" << std::endl;
        }
        virtual void g() {
            std::cout << "Base1::g()" << std::endl;
        }
        virtual void h() {
            std::cout << "Base1::h()" << std::endl;
        }
};
class Base2 {
    public:
        int ibase2;
        Base2():ibase2(2) {}
        virtual void f() {
            std::cout << "Base2::f()" << std::endl;
        }
        virtual void g() {
            std::cout << "Base2::g()" << std::endl;
        }
        virtual void h() {
            std::cout << "Base2::h()" << std::endl;
        }
};
class Base3 {
    public:
        int ibase3;
        Base3():ibase3(3) {}
        virtual void f() {
            std::cout << "Base3::f()" << std::endl;
        }
        virtual void g() {
            std::cout << "Base3::g()" << std::endl;
        }
        virtual void h() {
            std::cout << "Base3::h()" << std::endl;
        }
};

class Derive: public Base1, public Base2, public Base3{
    public:
        int iderive;
        Derive():iderive(4){}
        void f() { std::cout << "Derive::f()" << std::endl;}
        virtual void g1() {std::cout << "Dervie::g1()" << std::endl;}

};


int main() {
    Derive d;
    long *vptr = GET_VPTR(d);
    SHOW_ADDR(d);
    printByte(d);

    show_vptr(vptr);

    long *vptr2 = (long*)*(long*)((char*)&d + sizeof(Base1));
    show_vptr(vptr2);

    long *vptr3 = (long*)*(long*)((char*)&d + +sizeof(Base1) + sizeof(Base2));
    show_vptr(vptr3);
    printf("vptr1: %p\nvptr2: %p\nvptr3: %p\n", vptr, vptr2, vptr3);
}
