#include <iostream>

struct Node {
    int data;
    Node* next;
};

class LinkedList {
private:
    Node* head;

public:
    LinkedList() {
        head = nullptr;
    }

    void insert(int value) {
        // Insertion logic...
    }

    void display() {
        Node* current = head;
        while (current != nullptr) {
            std::cout << current->data << " ";
            current = current->next;
        }
        std::cout << std::endl;
    }

    void reverse(){
        Node * current = head;
        Node * prev = nullptr;
        Node * next = nullptr;

        while(current != nullptr){

            next = current->next;
            if(current == head){
                prev = nullptr;
            }else{
                prev = nullptr;
            }
        }
    }

    void removeNthNode(int value){
        Node * fast = head;
        Node * slow = head;

        while{fast != nullptr}{
            
        }

    }
};

int main (){

}
