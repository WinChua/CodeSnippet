#include <fcntl.h>
#include <stdio.h>
#include <unistd.h>
#include <string.h>

#define BUF_SIZE 4096

int main(int argc, char * argv[]) {
    if( argc < 3 ) {
        printf("Usage %s source target\n", argv[0]);
        return -1;
    }
    int source_fd = open(argv[1], O_RDONLY);
    if( -1 == source_fd ) {
        perror(argv[1]);
        return -1;
    }
    int target_fd = open(argv[2], O_WRONLY | O_CREAT | O_TRUNC, S_IWUSR | S_IRUSR | S_IWGRP | S_IRGRP);
    if( -1 == target_fd ) {
        perror(argv[2]);
        return -1;
    }

    char data[BUF_SIZE+1] = {0};
    while( 0 < read(source_fd, data, BUF_SIZE) ) {
        write(target_fd, data, strlen(data));
    }

    
    
    close(source_fd);
    close(target_fd);
}
