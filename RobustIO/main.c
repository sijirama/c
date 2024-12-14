#include "csapp.h"
#include <fcntl.h>
#include <sys/stat.h>
#include <unistd.h>

int main(int argc, char **argv) {
	// rio_t rio;
	// int n;
	// char buf[MAXLINE];
	//
	// rio_readinitb(&rio, STDIN_FILENO);
	// while ((n == rio_readlineb(&rio, buf, n) != 0)) {
	// 	rio_writen(STDOUT_FILENO, buf, n);
	// }

	Read

		struct stat fileInfo;
	char *type, *isReadable;

	stat(argv[1], &fileInfo);
	if (S_ISREG(fileInfo.st_mode)) {
		type = "regular";
	} else if (S_ISDIR(fileInfo.st_mode)) {
		type = "directory";
	} else {
		type = "other";
	}

	if (fileInfo.st_mode & S_IRUSR)
		isReadable = "yh";
	else
		isReadable = "nah";

	printf("type: %s, read: %s\n", type, isReadable);
	exit(0);
}
