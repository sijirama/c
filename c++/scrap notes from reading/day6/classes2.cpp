#include <iostream>
using namespace std;

class Cat{
    unsigned int itsWeigth;
    unsigned int itsHeight;
    unsigned int itsAge;

    public:
        int GetAge() const;
        void SetAge(unsigned int age);
        void Meow();
        Cat(unsigned int age);
        ~Cat();
};

int Cat::GetAge() const {
    return itsAge;
};

Cat::Cat(unsigned int age){
    itsAge = age;
}

Cat::~Cat(){}

void Cat::SetAge(unsigned int age){
    itsAge = age;
};
void Cat::Meow(){
    cout <<"Meow." << endl;
};

int main(){
    Cat Frisky(30);
    unsigned int FriskysAge = Frisky.GetAge();
    cout << FriskysAge << endl;
    Frisky.Meow();
}
