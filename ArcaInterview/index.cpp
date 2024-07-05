
/*
	Write a function that takes a list of integers and returns a new list where
   each integer is replaced by the square of the integer if it is even, and by
   the cube of the integer if it is odd.
	([1, 2, 3, 4, 5]) == [1, 4, 27, 16, 125]
	([-3, -2, 0, 7, 8]) == [-27, 4, 0, 343, 64]
*/

#include <iostream>
#include <ostream>
#include <vector>
std::vector<int> replaceWithSquare(std::vector<int> arr) {
	for (int num = 0; num < arr.size(); num++) {

		if (arr[num] % 2 == 0) {
			arr[num] = arr[num] * arr[num];
		} else {
			arr[num] = arr[num] * arr[num] * arr[num];
		}

		std::cout << arr[num] << std::endl;
	}
	return arr;
}

int main() {
	std::vector<int> arr = {-3, -2, 0, 7, 8};
	arr = replaceWithSquare(arr);
}
