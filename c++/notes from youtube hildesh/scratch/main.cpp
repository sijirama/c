
#include <iostream>
#include <cstdint>
#include <vector>

struct User {
    const int *userId;
    const char *userName;
    const char *password;
};

enum MsOffice {
    RED,
    BLUE,
    GREEN
};

int main (){

    int id = 21;
    User mickey = { &id ,"mickey" , "password"};

    std::cout << "Welcome! " << mickey.userName << std::endl;
    std::cout << "Welcome! " << mickey.userId<< std::endl;

    const int uid  = MsOffice::GREEN;

    std::cout << uid << "\n";

    return 0;
}
