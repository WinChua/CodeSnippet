#include <unistd.h>
#include <stdlib.h>
#include <stdio.h>
#include <fcntl.h>
#include <string.h>

#define SEQFILE "seqno"
#define MAXLINE 80

void my_lock(int omit) {
    return;
}

void my_unlock(int omit) {
    return;
}

int main(int argc, char **argv) {
    int fd;
    long i, seqno;
    pid_t pid;
    ssize_t n;
    char line[MAXLINE+1];
    pid = getpid();
    fd = open(SEQFILE, O_RDWR, S_IRUSR | S_IWUSR | S_IRGRP | S_IWGRP);
    if (fd < 0) {
        perror("Open file Error:");
        exit(1);
    }

    for (i = 0; i<20; i++) {
        my_lock(fd);
        if(lseek(fd, 0L, SEEK_SET) != 0) {
            perror("Seek file error:");
            exit(1);
        }
        n = read(fd, line, MAXLINE);
        if(n == -1) {
            perror("Read error:");
            exit(1);
        }
        line[n] = '\0';
        n = sscanf(line, "%ld\n", &seqno);
        printf("%s: pid = %ld, seq# = %ld\n", argv[0], (long) pid, seqno);
        seqno++;
        snprintf(line, sizeof(line), "%ld\n", seqno);
        if(lseek(fd, 0L, SEEK_SET) != 0) {
            perror("Seek file error:");
            exit(1);
        }
        n = write(fd, line, strlen(line));
        if (n == -1) {
            perror("Write Error:");
            exit(1);
        }
        my_unlock(fd);
    }
    //close(fd);
}
