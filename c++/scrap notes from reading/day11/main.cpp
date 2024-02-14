
#include <iostream>
#include <string.h>
using namespace std;

int main(){

    int intArray[10] = {1,2,3,4,5,6,7,8,9};

    cout << *intArray << "\n" ;
    for (int i = 0 ; i < 10  ; i++){
        //cout << intArray[i] << "\n" ;
    }

    char helloworld[] = "hello world";
    cout << helloworld << "\n" ;

    char helo[50];
    
    strcpy_s(helo , helloworld);
    cout << helo<< "\n" ;


}
