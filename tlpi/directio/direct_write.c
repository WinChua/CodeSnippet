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
    //printf("%d\n", EINVAL);
    //exit(1);
    int fd;
    int openflag;
    openflag = O_CREAT | O_TRUNC | O_WRONLY;
#ifdef DIRECT
    openflag |= O_DIRECT;
#endif
    fd = open(argv[1], openflag, S_IRUSR | S_IWUSR);
    if(fd == -1) {
        perror("Open Error:");
        exit(1);
    }

    size_t size;
    size = atoi(argv[2]);
    size *= 1000;
    //size *= sizeof(image);
    size_t written = 0;
    //char *buf = (char *) malloc(size);
    void *buf;
    posix_memalign(&buf, BLOCKSIZE, size*BLOCKSIZE);
    void * data = buf;
    for(written=0; written < size; written++) {
        memcpy(buf, image, sizeof(image));
        buf += sizeof(image);
    }
    
    // int res;
    // res = write(fd, data, 512);
    // printf("the res of write(fd, data, 512) is %d\n", res);
    // res = write(fd, data, 42);
    // printf("the res of write(fd, data, 42) is %d\n", res);
    // if(write(fd, data, 42) == -1) {
    //     perror("Write Error:");
    //     exit(-1);
    // }
    int i;
    for(i=0; i * 1000<size; i++) {
        write(fd, data+i * 1000 *BLOCKSIZE, 1000 * BLOCKSIZE);
    }
    // for(written=0; written < size; written +=sizeof(image)) {
    //     write(fd, buf+written, sizeof(image));
    // }
    printf("written is %d, size is %d\n", written, size);
    close(fd);
    free(data);
    return 0;
}
