
#include <iostream>
#include <queue>
#include <map>
#include <vector>

using namespace std;

class Node {
  public:
	int data;
	Node *left, *right;
	Node(int value) {
		data = value;
		left = nullptr;
		right = nullptr;
	}
};

void printLevelOrder(Node *root) {
	if (root == nullptr) {
		return;
	}

	std::queue<Node *> Queue;
	Queue.push(root);

	while (!Queue.empty()) {
		int levelSize = Queue.size();
		for (int i = 0; i < levelSize; i++) {

			Node *current = Queue.front();
			Queue.pop();

			cout << current->data;

			if (current->left) {
				Queue.push(current->left);
			}
			if (current->right) {
				Queue.push(current->right);
			}
		}
		cout << endl;
	}
}

std::map<int, vector<int>> mapp;
void printLevelOrder2(Node *root, int depth) {

	//cout << root->data << endl;
	mapp[depth].push_back(root->data);

	if (root->left) {
		printLevelOrder2(root->left, depth+1);
	}
	if (root->right) {
		printLevelOrder2(root->right, depth+1);
	}
}

void printMap() {
	for (const auto &array : mapp) {
		for (const int num : array.second) {
			cout << num << ", ";
		}
		cout << endl;
	}
}

int main() {
	Node *root = new Node(1);
	root->left = new Node(2);
	root->right = new Node(3);
	root->left->left = new Node(4);
	root->left->right = new Node(5);
	root->left->left->left = new Node(4);
	root->left->left->right = new Node(4);
	root->left->right->left = new Node(5);
	root->left->right->right = new Node(5);
	printLevelOrder(root);
	printLevelOrder2(root, 0);
	printMap();
	return 0;
}
