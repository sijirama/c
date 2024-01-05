#include <iostream>
class Person{

    public:
    Person(){this->itsAge = 10;}
    ~Person(){}
    void setAge(int age) {this->itsAge = age;}
    int getAge(){ return this->itsAge;}

    private:
        int itsAge;
};


int main(){
    Person Frank;
    Frank.setAge(30);
    std::cout << "Frank is " << Frank.getAge() << "years old" << std::endl ;
}
