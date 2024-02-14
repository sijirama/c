#include <iostream>
//#include <iostream>

int main(){
    int intOne = 90;
    int &rSomeRef = intOne;
    //std::cout << rSomeRef << std::endl;
    //std::cout << &rSomeRef << std::endl;
    //std::cout << &intOne<< std::endl;
    int * pInt = new int ;
    *pInt = 90;
    std::cout << *pInt << std::endl;
    int & rPref = *pInt;
    std::cout << rPref << std::endl;
    rPref = 100;
    std::cout << *pInt << std::endl;

}
