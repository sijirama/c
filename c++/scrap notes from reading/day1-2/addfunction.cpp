#include <iostream>
using namespace std;

int add(int a , int b){
    cout << "From function Add()" << endl;
    return (a + b);
};

int main(){
    
    cout << "In main function" << endl;
    int a, b, c;
    cout << "Enter 2 numbers" << endl;
    cin >> a;
    cin >> b;
    c = add(a, b);
    cout << "The sum of your input was " << c << endl;

    return 0;
}
