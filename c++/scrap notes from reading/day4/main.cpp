#include <iostream>
using namespace std;

// Testing subtraction between unsigned numbers
int main(){
    unsigned int difference;
    unsigned int bigNumber = 100;
    unsigned int smallNumber = 120;

    difference = ++bigNumber; //bigNumber = 100
    cout << "The difference is " << difference << endl ; //100
    cout << "The bignumber is " << bigNumber ; //101
}
