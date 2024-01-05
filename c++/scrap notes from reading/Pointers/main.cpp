#include <iostream>
#include <ostream>

int main(){
    std::cout << "-----------------------------Pointers-------------------------------" << std::endl;

    unsigned short shortVar = 5;
    unsigned long longVar = 60005;
    long sVar = -60005;
    //unsigned short int *pVar = &shortVar;

    
    unsigned short int *pVar = new unsigned short int;
    *pVar = 10;


    std::cout << "short var pointer: \t" << *pVar << std::endl;

    delete pVar;

    //pVar = 0;
    
    std::cout << "short var pointer: \t" << *pVar ;

    // std::cout << "short var: \t" << shortVar ;
    // std::cout << "address of short var: \t" << &shortVar ;
    //
    // std::cout << std::endl ;
    //
    // std::cout << "long var: \t" << longVar;
    // std::cout << "address of long var: \t" << &longVar;
    //
    // std::cout << std::endl ;
    //
    // std::cout << "signed var: \t" << sVar;
    // std::cout << "address of signed var: \t" << &sVar;
    

}
