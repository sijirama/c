#include <iostream>
using namespace std;
typedef short int USHORT ;

//int NumberofFibsCalled;

int main(){
    USHORT fib(int);
    cout << fib(5) ;
}

USHORT fib(int n){
    if(n < 3)
        return 1;
    return fib(n - 1) + fib(n - 2);
}
