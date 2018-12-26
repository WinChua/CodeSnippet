#include <sys/mman.h>
#include <fcntl.h>
#include <stdio.h>
#include <stdlib.h>

#define COPYINCR (1024 * 1024 * 1024) / * 1 GB * /

int main(int argc, char * argv[]) {
    int fdin, fdout;
    void * src, * dst;
    size_t copysz;
    struct stat sbuf;
    off_t fsz = 0;

    if(argc != 3) {
        printf("usage: %s fromfile tofile\n", argv[0]);
        exit(1);
    } 

    if((fdin = open(argv[1], O_RDONLY)) < 0) {
        perror("fromfile error: ");
        exit(1);
    }
    if((fdout = open(argv[2], O_RDWR | O_CREAT | O_TRUNC, 
        S_IRUSR | S_IWUSR | S_IRGRP | S_IWGRP)) < 0) {
        perror("tofile error: ");
        exit(1);
    }
    if(fstat(fdin, &sbuf) < 0) {
        
    }
}
