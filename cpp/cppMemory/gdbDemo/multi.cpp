#include <iostream>

class Mother {
    public:
        Mother(): mother_data(1) {}
        virtual void MotherMethod() {}
        int mother_data;
};

class Father {
    public:
        Father(): father_data(2) {}
        virtual void FatherMethod() {}
        int father_data;
};

class Child: public Mother, public Father {
    public:
        Child(): child_data(3) {}
        virtual void ChildMethod() {}
        int child_data;
};

int main() {
    Child child;
    child.ChildMethod();
    std::cout << "Hello, World" << std::endl;
}
