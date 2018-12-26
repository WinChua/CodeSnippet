#include <unistd.h>
#include <stdio.h>
#include <stdlib.h>
#include <fcntl.h>
#include <sys/types.h>
#include <string.h>

int main(int argc, char * argv[]) {
    if(argc < 2) {
        printf("Usage %s chars\n", argv[0]);
        exit(1);
    }
    printf("the pid is %d\n", getpid());
    umask(0);
    int fd = open("hello", O_WRONLY | O_CREAT | O_TRUNC, S_IRUSR | S_IWUSR | S_IRGRP | S_IWGRP);
    struct flock lock;
    lock.l_type = F_WRLCK;
    lock.l_start = 0;
    lock.l_whence = SEEK_SET;
    lock.l_len = 1;
    while(fcntl(fd, F_SETLK, &lock)) {
    }
    write(0, "get the lock\n", sizeof "get the lock\n");
    write(fd, argv[1], strlen(argv[1]));
    close(fd);
    return 0;
    
}
