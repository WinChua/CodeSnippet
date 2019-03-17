#include <iostream>

class Person {
    public:
        int iperson;
        Person():iperson(1){}
};

class Father : public virtual Person {
    public:
        Father():ifather(2){}
        int ifather;
};

class Mother : public virtual Person {
    public:
        Mother(): imother(3){}
        int imother;
};

class Child : public Father, public Mother {
    public:
        int ichild;
        Child(): ichild(4){}
};

int main() {
    Child child;
    child.iperson = 42;
    std::cout << "Hello, World" << std::endl;
}
