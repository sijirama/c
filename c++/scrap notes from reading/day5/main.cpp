typedef unsigned int USHORT;
#include <iostream>
using namespace std;
USHORT area (int length, int breadth);

/*
    Function calls
    How do function works in c++
    function prototype
*/

int main(){

    USHORT a , b , c;

    cout << "Input the length :";
    cin >> a;
    cout << endl;
    
    cout << "Input the breadth :";
    cin >> b;
    cout << endl;
    
    c = area(a, b);
    cout << "The output is:" << c << endl;
    return 0;
}

USHORT area (int l , int b){
    return l * b;
}
