#include <string>
#include <iostream>
using namespace std;

int main(){
    string color , firstName , lastName;
    cout << "what is your first Name?" << endl ;
    getline(std::cin , firstName); 
    
    cout << "what is your last Name?" << endl ;
    getline(std::cin , lastName);

    cout << "Welcome to cpp "  << lastName<< " " << firstName << endl;
}
