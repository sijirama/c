
#include <iostream>
#include <stdio.h>
#include <string>
using namespace std;

class Employee {

    private:
        string _name;
        string _company;
        int _age;
        int _salary;
        Employee();

    public:
        Employee(string name, string company , int age , int salary = 10000){
            _name = name;
            _company = company;
            _age = age;
            _salary = salary;
        };
        void setName(string name) { _name = name; }
        void setCompany(string company) { _company = company; }
        string getName() { return _name; }
        string getCompany() { return _company; }
        int getAge() { return _age; }
        bool isOlderThan (Employee other) {return (this->getAge() > other.getAge());}
        bool isRicherThan (Employee other) {return (this->_salary > other._salary);}
};

int main(){

    Employee jacob("jacob" , "Jacob Ltd" , 23 , 20000);

    jacob.setName("Jacob statham");

    Employee Mark ("Mark jones" , "Yamaha" , 23);

    if(jacob.isOlderThan(Mark)){
        puts("jacob is older than Mark");
    }else{
        puts("Mark is older than jacob");
    }
    
    if(jacob.isRicherThan(Mark)){
        puts("jacob is richer than Mark");
    }else{
        puts("Mark is richer than jacob");
    }

    cout << jacob.getName() << endl;
}
