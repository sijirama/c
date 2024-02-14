#include <iostream>
using namespace std;

class Mammal {
    public:
        Mammal(){

        }
        ~Mammal(){

        }
        void setAge(int age){
            itsAge = age;
        }
        int getAge(){
            return itsAge;
        }
        void setWeight(int weight){
            itsWeight = weight;
        }
        int getWeight(){
            return itsWeight;
        }
        void move (){
            cout << "im moving " << endl;
        }
        virtual void speak(){cout << "Mammal is speaking" << endl;}


    private:
        int itsAge;
        int itsWeight;
};

enum BREED {RETRIEVER , CHIHUAHA , DOBERMAN , LAB};

class Dog : public Mammal {
    public:
    Dog(BREED breed){
        itsBreed = breed;
    }
    ~Dog(){}

    void setBreed(BREED breed){
        itsBreed = breed;
    }

    BREED getBreed(){
        return itsBreed;
    }
    void move (){
        cout << "Dog is moving " << endl;
    }
    void speak(){cout << "Dog is speaking" << endl;}

    private:
        BREED itsBreed;
};

int main(){
    Dog Fido(RETRIEVER);
    cout << "Fido is a " << Fido.getBreed() << endl;
    Fido.setAge(28);
    Fido.setWeight(21);
    cout << "Fido is " << Fido.getAge() << "years old and he is " << Fido.getWeight() << endl;
    Fido.setBreed(BREED(DOBERMAN));
    cout << "Fido is a " << Fido.getBreed() << endl ;
    Fido.Mammal::move();

    Mammal * pMammal = new Dog(RETRIEVER);
    pMammal ->speak();

    delete pMammal;

    return 0;
}
