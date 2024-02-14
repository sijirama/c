#include <iostream>
#include <ostream>
#include <vector>
using namespace std;

int main(){
    void swap(int & a , int & b);

    std::vector<int> hello = {1,2,3,4,5,6,7,8,9,10};
    hello.pop_back();
    cout << hello[1] << endl;
    for(int i : hello){
        std::cout << i << " " << std::endl ; 
    }

}

void swap (int & a , int & b){
    int temp = a;
    a = b;
    b = temp;
}
