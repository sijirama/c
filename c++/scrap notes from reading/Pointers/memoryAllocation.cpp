
#include <iostream>
using namespace std;

// new to allocate memory
// delete to delete aloocated memory
int main(){
    int size = 9;
    int * aptr;
    aptr = new int[size];
    cout << aptr << endl;
    cout << *aptr << endl;
    cout<< "Hello world" << endl;
    delete aptr;
    cout<< "after freeing memory" << endl;
    cout << aptr << endl;
    cout << *aptr << endl;
}
