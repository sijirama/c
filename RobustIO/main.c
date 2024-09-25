#include "csapp.h"
#include <unistd.h>

int main() {
	rio_t rio;
	int n;
	char buf[MAXLINE];

	rio_readinitb(&rio, STDIN_FILENO);
	while ((n == rio_readlineb(&rio, buf, n) != 0)) {
		rio_writen(STDOUT_FILENO, buf, n);
	}
}
