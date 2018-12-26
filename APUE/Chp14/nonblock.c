#include <errno.h>
#include <fcntl.h>
#include <stdlib.h>
#include <stdio.h>

char buf[500000];

void clr_fl(int fd, int flags) {
    int val;
    if((val = fcntl(fd, F_GETFL, 0)) < 0) {
        perror("clr fcntl F_GETFL error");
        exit(1);
    }
    val &= ~flags;
    if(fcntl(fd, F_SETFL, val) < 0) {
        perror("clr fcntl F_SETFL error");
        exit(1);
    }
    
}
void set_fl(int fd, int flags) {
    int val;
    if((val = fcntl(fd, F_GETFL, 0)) < 0) {
        perror("set fcntl F_GETFL error");
        exit(1);
    }
    val |= flags;
    if(fcntl(fd, F_SETFL, val) < 0) {
        perror("set fcntl F_SETFL error");
        exit(1);
    }
    
}

int main(void) {
    int ntowrite, nwrite;
    char * ptr;
    
    ntowrite = read(0, buf, sizeof(buf));
    fprintf(stderr, "read %d bytes\n", ntowrite);

    set_fl(1, O_NONBLOCK);

    ptr = buf;
    while(ntowrite > 0) {
        errno = 0;
        nwrite = write(1, ptr, ntowrite);
        fprintf(stderr, "nwrite = %d, errno = %d\n", nwrite, errno);
        if(nwrite > 0) {
            ptr += nwrite;
            ntowrite -= nwrite;
        }
    }

    clr_fl(1, O_NONBLOCK);
    exit(0);
}
