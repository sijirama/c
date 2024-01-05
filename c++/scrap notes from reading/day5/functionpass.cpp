typedef unsigned int USHORT;
#include <iostream>

int main(){

    USHORT changeValue (int valuetobechanged);
    USHORT value = 10;
    
    changeValue(value);
    USHORT c = changeValue(value);
    
    std::cout << value << std::endl;
    std::cout << c << std::endl;

}

USHORT changeValue (int x){
    return x = 20;
}
