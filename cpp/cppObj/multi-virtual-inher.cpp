#include <iostream>

using namespace std;

class Person {
    public:
        Person():iPerson(1){}
        int iPerson;
        virtual void helloPerson() { cout << "Person::helloPerson()" << endl; }
};
class Father: virtual public Person {
    public:
        Father():iFather(2){}
        int iFather;
        virtual void helloFather() { cout << "Father::helloFather()" << endl; }
};
class Mother: virtual public Person {
    public:
        Mother():iMother(3){}
        int iMother;
        virtual void helloMother() { cout << "Mother::helloMother()" << endl; }
};
class Child: public Father, public Mother {
    public:
        Child():iChild(4){}
        int iChild;
        virtual void helloChild() { cout << "Child::helloChild()" << endl; }
};

int main() {
    Person person;
    person.helloPerson();
    Father father;
    father.helloFather();
    Mother mother;
    mother.helloMother();
    Child child;
    child.helloChild();
    // child.helloPerson(); // compile error: request for member "helloPerson()" is ambiguous;
}
