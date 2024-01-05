#include "Cat.h"
#include <string>
#include <iostream>

Cat::Cat(){
    name = "mika";
}

Cat::Cat(std::string name){
    name = name;
}

Cat::~Cat(){}

void Cat::Meow() const {
    std::cout << "Meow, my name is" << name << "\n";
}
void Cat::Purr() const {
    std::cout << "Purr, my name is" << name << "\n";
}

std::string Cat::getName(){

    return name;
}
