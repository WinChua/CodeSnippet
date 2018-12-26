#define _GNU_SOURCE
#include <fcntl.h>
#include <stdlib.h>
#include <stdio.h>
#include <unistd.h>
#include <string.h>
#include <errno.h>
#define BLOCKSIZE 512
char image[] =
{
	'P', '5', ' ', '2', '4', ' ', '7', ' ', '1', '5', '\n',
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 3, 3, 3, 3, 0, 0, 7, 7, 7, 7, 0, 0,11,11,11,11, 0, 0,15,15,15,15, 0,
	0, 3, 0, 0, 0, 0, 0, 7, 0, 0, 0, 0, 0,11, 0, 0, 0, 0, 0,15, 0, 0,15, 0,
	0, 3, 3, 3, 0, 0, 0, 7, 7, 7, 0, 0, 0,11,11,11, 0, 0, 0,15,15,15,15, 0,
	0, 3, 0, 0, 0, 0, 0, 7, 0, 0, 0, 0, 0,11, 0, 0, 0, 0, 0,15, 0, 0, 0, 0,
	0, 3, 0, 0, 0, 0, 0, 7, 7, 7, 7, 0, 0,11,11,11,11, 0, 0,15, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0,
};

int main(int argc, char * argv[]) {
    int fd;
    int openflag = O_CREAT | O_TRUNC | O_WRONLY;
#ifdef DIRECT
    openflag |= O_DIRECT;
#endif
    fd = open(argv[1], openflag, S_IRUSR | S_IWUSR);
    if(fd == -1) {
        perror("ERROR Open:");
        exit(1);
    }
    void * buf;
    posix_memalign(&buf, BLOCKSIZE, 1024 * BLOCKSIZE);
    void *data = buf;
    int i;
    for(i=0; i<1024; i++) {
        memcpy(data, image, BLOCKSIZE);
        data += BLOCKSIZE;
    }
    size_t size = atoi(argv[2]);
    for(i=0; i<size; i++) {
        if(-1 == write(fd, buf, 1024 * BLOCKSIZE)) {
            perror("Error Write:");
            exit(1);
        }
    }
    close(fd);
    free(buf);

}
