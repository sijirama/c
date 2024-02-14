#include "Cat.h"
#include <iostream>
#include <iterator>

int main(){
    Cat Friska ; 
    Friska.Meow();
    Friska.Purr();
    std::cout << "My name is " << Friska.getName() << "\n";
}
