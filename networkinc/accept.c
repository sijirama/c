#include <netdb.h>
#include <netinet/in.h>
#include <string.h>
#include <sys/socket.h>
#include <sys/types.h>
#define MYPORT "3490"
#define BACKLOG 10

int main() {
	struct addrinfo hints, *res;
	int sock_fd, new_fd;
	socklen_t addr_size;
	struct sockaddr_storage their_addr;

	memset(&hints, 0, sizeof hints);
	hints.ai_family = AF_UNSPEC; // use IPv4 or IPv6, whichever
	hints.ai_socktype = SOCK_STREAM;
	hints.ai_flags = AI_PASSIVE;

	int status = getaddrinfo(NULL, MYPORT, &hints, &res);
	sock_fd = socket(res->ai_family, res->ai_socktype, res->ai_protocol);
	int binstat = bind(sock_fd, res->ai_addr, res->ai_addrlen);
	listen(sock_fd, BACKLOG);

	addr_size = sizeof their_addr;
	new_fd = accept(sock_fd, (struct sockaddr *)&their_addr, &addr_size);
}
