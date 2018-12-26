#define _GNU_SOURCE
#include <fcntl.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

int main(int argc, char * argv[]) {
    int fd;
    fd = open(argv[1], O_RDONLY | O_DIRECT);
    if(fd == -1) {
        perror("Open Error:");
        exit(1);
    }
    size_t size;
    size = atoi(argv[2]);
    //char * buf = (char *)malloc(512);
    char * buf = (char *)memalign(0, 512) + 0;
    int numRead = read(fd, buf, 10);
    if(numRead == -1) {
        perror("Read Error:");
        exit(1);
    }
    close(fd);
    printf("Read %d bytes\nRead %s\n", numRead, buf);
    free(buf);
    return 0;
}
