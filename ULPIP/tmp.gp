#include <fcntl.h>
#include <unistd.h>
#include <stdio.h>
#include <stdlib.h>

#define BUF_SIZE 1024
#define COPYMODE 0644

int main(int argc, char * argv[]) {
    int in_fd, out_fd, n_chars;
    char buf[BUF_SIZE];
    if( argc != 3 ) {
        fprintf(stderr, "usage : %s should have 3 args.\n", *argv);
        exit(1);
    }
    if( (in_fd = open(argv[1], O_RDONLY)) == -1 ) {
        perror("Error Open:");
        exit(1);
    }
    if( (out_fd = creat(argv[2], COPYMODE)) == -1) {
        perror("Error Creat:");
        exit(1);
    }
    while( (n_chars = read(in_fd, buf, BUF_SIZE)) > 0 ) {
        if( (write(out_fd, buf, n_chars)) != n_chars ) {
             perror("Error Copy:"); 
             exit(1);
        }
    }

    if( close(in_fd) == -1 || close(out_fd) == -1) {
        perror("Error closing:");
        exit(1);
    }
}
