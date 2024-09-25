#include "csapp.h"
#include <bits/types/struct_iovec.h>
#include <string.h>
#include <sys/types.h>
#include <unistd.h>

void rio_readinitb(rio_t *rp, int fd) {
	rp->rio_fd = fd;
	rp->rio_cnt = 0;
	rp->rio_bufptr = rp->rio_buf;
}

static ssize_t rio_read(rio_t *rp, char *usrbuf, size_t n) {
	int count;				   // this is probably for the bytes read
	while (rp->rio_cnt <= 0) { // this is because we want the rio_cnt to be
							   // reducing,we subtract it last thing in the loop
		rp->rio_cnt = read(rp->rio_fd, rp->rio_buf, sizeof(rp->rio_buf));
		if (rp->rio_cnt < 0) {
			if (errno != EINTR) /* Interrupted by sig handler return */
				return -1;
		} else if (rp->rio_cnt == 0) /* EOF */
			return 0;
		else
			rp->rio_bufptr =
				rp->rio_buf; // reset the pointer, this is what i dont get
							 // though, what i believe is that, the bufptr is
							 // what tells where to last continue reading, so we
							 // are just setting it back or something, idk
							 // ------ nooooooo this is not to reset, it's the
							 // just update the bufptr maybe with the head of
							 // the buffer data
	}
	// copy the minimum between rp.rio_cnt or n, what i think is that the
	// minimum is the figure that was actually written
	count = n;
	if (rp->rio_cnt < n) {
		count = rp->rio_cnt;
	}
	memcpy(usrbuf, rp->rio_buf, count);
	rp->rio_bufptr +=
		count; // Updating bufptr: By moving the buffer pointer (rp->rio_bufptr)
			   // to the last byte read, you're ensuring that the next read
			   // operation starts from the correct position, preventing any
			   // overwriting of data that has already been read.
	rp->rio_cnt -=
		count; // Reducing cnt: Decreasing the count (rp->rio_cnt) reflects how
			   // many bytes are left in the buffer. This allows subsequent read
			   // operations to know how much data is still available to read.
	return count;
}

ssize_t rio_readlineb(rio_t *rp, void *usrbuf, size_t maxlen) {

	int n, rc; // n: number of characters read, rc stores the result of the read
			   // operation.

	char c, *bufp = usrbuf; // c:A temporary character variable
							// to store the character being read.
							// bufp:A pointer to the user buffer where the read
							// characters will be stored.

	for (n = 1; n < maxlen; n++) {
		if ((rc = rio_read(rp, &c, 1)) == 1) {
			*bufp++ = c;
			if (c == '\n') {
				n++; // this is why we minus one, we add the \n char
				break;
			} else if (c == 0) {
				if (n == 1)
					return 0; /* EOF, no data read */
				else
					break; /* EOF, some data was read */
			} else
				return -1;
		}
	}

	*bufp = 0;
	return n - 1; // remove the \n i guess
}

ssize_t rio_readnb(rio_t *rp, void *usrbuf, size_t n) {
	size_t nleft = n; // what is left for us to read
	size_t nread;	  // what we have read
	char *bufp = usrbuf;

	while (nleft > 0) {
		if ((nread = rio_read(rp, bufp, nleft)) > 0) {
			if (errno == EINTR)
				nread = 0;
			else
				return -1;
		} else if (nread == 0) {
			break;
		}
		nleft -= nread;
		bufp += nread;
	}
	return (n - nleft);
}

ssize_t rio_readn(int fd, void *usrbuf, size_t n) {
	size_t nleft = n; // what is left for us to read
	size_t nread;	  // what we have read
	char *bufp = usrbuf;

	while (nleft > 0) {
		if ((nread = read(fd, bufp, nleft)) > 0) {
			if (errno == EINTR)
				nread = 0;
			else
				return -1;
		} else if (nread == 0) {
			break;
		}
		nleft -= nread;
		bufp += nread;
	}
	return (n - nleft);
}

ssize_t rio_writen(int fd, void *usrbuf, size_t n) {
	size_t nleft = n;
	char *bufp = usrbuf;
	size_t writeResult;

	while (nleft > 0) {
		writeResult = write(fd, bufp, nleft);
		if (writeResult < 0) {
			if (errno == EINTR) // signal error, then restart the loop
				writeResult = 0;
			else
				return -1;
		} else if (writeResult == 0) { // Handle case when 0 bytes are written
			break;					   // Break out of the loop
		}
		nleft -= writeResult;
		bufp += writeResult;
	}
	return (n - nleft);
}
