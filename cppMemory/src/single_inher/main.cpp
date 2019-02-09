#include "vptr.h"
#include "bytes.h"
#include <iostream>

class Parent {
    public:
        int iparent;
        Parent():iparent(1) {}
        virtual void f() {
            std::cout << "Parent::f()" << std::endl;
        }
        virtual void g() {
            std::cout << "Parent::g()" << std::endl;
        }
        virtual void h() {
            std::cout << "Parent::h()" << std::endl;
        }
};
class Child : public Parent{
    public:
        int ichild;
        Child():ichild(2) {}
        virtual void f() {
            std::cout << "Child::f()" << std::endl;
        }
        virtual void g_child() {
            std::cout << "Child::g_child()" << std::endl;
        }
        virtual void h_child() {
            std::cout << "Child::h_child()" << std::endl;
        }
};

class GrandChild : public Child{
    public:
        int igrandchild;
        GrandChild():igrandchild(3) {}
        virtual void f() {
            std::cout << "GrandChild::f()" << std::endl;
        }
        virtual void g_child() {
            std::cout << "GrandChild::g_child()" << std::endl;
        }
        virtual void h_grandchild() {
            std::cout << "GrandChild::h_grandchild()" << std::endl;
        }
};

int main() {
    GrandChild gc;
    SHOW_ADDR(gc);printByte(gc);
    long *vptr = GET_VPTR(gc);
    printf("Hello World\n");
    show_vptr(vptr);
    printf("Hello World\n");
}
