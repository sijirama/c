#include <arpa/inet.h>
#include <netdb.h>
#include <netinet/in.h>
#include <stdio.h>
#include <string.h>
#include <sys/socket.h>

int main(int argc, char *argv[]) {
	struct addrinfo hints, *res, *p;
	int status;

	char ipstr[INET6_ADDRSTRLEN];

	if (argc != 2) {
		fprintf(stderr, "usage: showip hostname\n");
		return 1;
	}

	memset(&hints, 0, sizeof hints);
	hints.ai_family = AF_UNSPEC; // AF_INET or AF_INET6 to force version
	hints.ai_socktype = SOCK_STREAM;

	status = getaddrinfo(argv[1], NULL, &hints, &res);
	if (status != 0) {
		fprintf(stderr, "getaddrinfo: %s\n", gai_strerror(status));
		return 2;
	}

	printf("IP addresses for %s:\n\n", argv[1]);

	int socketfd = socket(res->ai_family, res->ai_socktype, res->ai_protocol);
	int bindfd = bind(socketfd, res->ai_addr, res->ai_addrlen);

	freeaddrinfo(res);
}
