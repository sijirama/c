
#include <any>
#include <iostream>
#include <memory>

// void printer(std::shared_ptr<int> value) { std::cout << *value << std::endl;
// }

void anyPrint(std::any value) {
	// std::cout << std::any_cast<std::shared_ptr<int>>(value) << std::endl;
	int converted_value =
		std::any_cast<int>(*std::any_cast<std::shared_ptr<int>>(&value));
	std::cout << converted_value << std::endl;
}

void anyPrintSmart(std::any value) {
	try {
		auto dereferenced = *std::any_cast<std::shared_ptr<int>>(&value);
		if (typeid(*dereferenced) == typeid(std::shared_ptr<int>)) {
			// Create a new shared_ptr if types match
			auto newSharedPtr = std::make_shared<int>(*dereferenced);
			// Use newSharedPtr here (e.g., call printer(newSharedPtr))
		} else {
			// Handle type mismatch
			std::cerr << "Error: Value in any is not a compatible smart pointer"
					  << std::endl;
		}
	} catch (const std::bad_any_cast &e) {
		// Handle the exception if the type is not the expected smart pointer
		// type
		std::cerr
			<< "Error: Value in any is not the expected smart pointer type"
			<< std::endl;
	}
}

int main() {
	std::variant<int, std::shared_ptr<int>> value;
	auto value = std::make_shared<int>(32);
	anyPrintSmart(value);
	// printer(value);
}
