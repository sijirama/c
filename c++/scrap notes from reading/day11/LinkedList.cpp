#include <iostream>

struct Node {
    int data;
    Node * next;
};

class List{
    private:
        Node * head;

    public:
        List(){
            std::cout << "New Linked List \n";
            head = nullptr;
        }

        ~List(){
            Node* current = head;
            while(current != nullptr){
                Node * next  = current->next;
                delete current;
                current = next;
            }
            std::cout << "New Linked Destroyed\n";
        }

        void insert (int value){

            //NOTE: create new Node
            Node * newNode = new Node;
            newNode->data = value;
            newNode->next = nullptr;

            //NOTE: make new node head if theres no head
            if(head == nullptr){
               head = newNode;
            }else{
            //NOTE: if theres head then get the last of the linkedlist and insert the new node there
                Node * current = head;
                while(current->next != nullptr){
                    current = current->next;
                }
                current->next = newNode;
            }
        }

        void display(){
            std::cout << "Printing \n";
            Node * current = head;
            while (current != nullptr){
                std::cout << current->data << "\n";
                //std::cout << "--> " << std::endl ;
                current = current->next;
            }
            std::cout << "nullptr \n";
        }
};

int main(){
    List LinkedList;
    LinkedList.insert(1);
    LinkedList.insert(2);
    LinkedList.insert(3);
    LinkedList.insert(4);

    LinkedList.display();
}
