#include <iostream>
using namespace std;


int main(){
    class Cat{
        unsigned int weigth;
        unsigned int legs;

        public:void bark(){
            cout << "bark like a dog" << endl;
        }

        public: void meow(){
            cout << "Meo like a cat" << endl;
        }
    }; 
    Cat Frisky;

    Frisky.meow();
    Frisky.bark();

    Frisky.meow();

    return 0;
}
