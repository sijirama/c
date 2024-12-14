#include <math.h>
#include <stdio.h>
#include <string.h>

int giveUsSomething(char x, int power) {
	int n = 0;

	if (x >= '0' && x <= '9') { // is a digit,
		n = (x - '0') * pow(16, power);
	} else if (x >= 'a' && x <= 'f') {
		n = (x - 'a' + 10) * pow(16, power); // Convert 'a'-'f' to 10-15
	} else if (x >= 'A' && x <= 'F') {
		n = (x - 'A' + 10) * pow(16, power); // Convert 'A'-'F' to 10-15
	} else {
		printf("Error: Invalid hex character '%c'\n", x);
		return -1;
	}
	return n;
}

int htoi(const char *hex) {
	int n = 0;

	const char *hexValue =
		(hex[0] == '0' && (hex[1] == 'x' || hex[1] == 'X')) ? &hex[2] : hex;

	int len = strlen(hexValue);

	for (int index = len - 1, power = 0; index >= 0; index--, power++) {
		int value = giveUsSomething(hexValue[index], power);
		if (value == -1) { // Error encountered
			return -1;
		}
		n += value;
	}

	return n;
}

int main() {
	printf("%d\n", htoi("12"));		  // 18
	printf("%d\n", htoi("0x12"));	  // 18
	printf("%d\n", htoi("0x212"));	  // 530
	printf("%d\n", htoi("000012"));	  // 18
	printf("%d\n", htoi("0x7F"));	  // 127
	printf("%d\n", htoi("0xfff"));	  // 4095
	printf("%d\n", htoi("0xabcdef")); // 11259375
	printf("%d\n", htoi("G12"));	  // Error case
}
