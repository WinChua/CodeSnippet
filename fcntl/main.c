#include <stdlib.h>
#include <stdio.h>
#include <unistd.h>
#include <string.h>
#include "filelock.h"

#define SEQNO "seqno"

int main(int argc, char ** argv) {
    if (argc < 3) {
        printf("usage %s [w|r] filename\n", argv[0]);
        exit(0);
    }
    int fd;
    fd = open(argv[2], O_RDWR);
    if (fd == -1) {
        perror("Open file error:");
        exit(1);
    }
    int r;
    if(strcmp(argv[1], "w") == 0) {
        r = write_lock(fd, 0, 0, 0);
    }
    else {
        r = read_lock(fd, 0, 0, 0);
    }
    if(r == -1) {
        perror("lock error:");
        exit(1);
    }
    else if (r == 0) {
        printf("read_lock succ\n");
    }

    sleep(4);
    r = un_lock(fd, 0, 0, 0);
    if(r == -1) {
        perror("Unlock error:");
    }
    else {
        printf("unlock succ\n");
    }
    close(fd);
}
