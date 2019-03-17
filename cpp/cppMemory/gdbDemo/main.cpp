#include <iostream>

class Parent {
    public:
        int iparent;
        Parent():iparent(1) {}
        virtual void Foo() {}
};

class Child: public Parent {
    public:
        int ichild;
        Child():ichild(2) {}
        virtual void Foo() {}
};

class vChild: virtual public Parent {
    public:
        int ivchild;
        vChild():ivchild(3) {}
        virtual void Foo() {}
};

int main() {
    Parent p;
    p.Foo();
    Child c;
    c.Foo();
    vChild vc;
    vc.Foo();
}
