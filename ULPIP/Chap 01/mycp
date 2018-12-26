#include <fcntl.h>
#include <unistd.h>
#include <stdio.h>
#include <stdlib.h>

#define BUFFERSIZE 4096
#define COPYMODE 0644

void oops(char *, char *);
/*
源码实现了 cp 的简单demo功能:
首先通过 open(fcntl.h中), 打开一个文件，返回文件描述符， 如果返回 -1 表明失败了
接着 通过 access 测试 文件 argv[2] 是否存在 F_OK 标志， 如果存在选则是否覆盖它
如果选择覆盖则 : creat 文件， 使用 COPYMODE 的打开模式
调用 read : int read(int fd, char * buf, size_t size);
将得到的 buf 写入 out_fd 中 : write(int fd, char * buf, size_t size) 返回 写入的字节数
最后 close
*/
int main(int argc, char * argv[]) {
    int in_fd, out_fd, n_chars;
    char buf[BUFFERSIZE];
    if(argc != 3) {
        fprintf(stderr, "usage: %s source destination\n", *argv);
        exit(1);
    }
    if( (in_fd = open(argv[1], O_RDONLY)) == -1)
        oops("Cannot open ", argv[1]);
    if( (access(argv[2], F_OK)) == 0 ) {
        int c;
        printf("The target file %s exists.\n Cover it?(Y/N)", argv[2]);
        while( (c = getchar()) != EOF ) {
            if (c == 'Y') break;
            if (c == 'N') exit(0);
            printf("Please choose Y or N. ");
        }
    }
    if( (out_fd = creat(argv[2], COPYMODE)) == -1 )
        oops("Cannote creat ", argv[2]);

    while( (n_chars = read(in_fd, buf, BUFFERSIZE)) > 0) {
        if( write(out_fd, buf, n_chars) != n_chars ) 
            oops("Write error to ", argv[2]);
    }
    if( n_chars == -1 )
        oops("Read error from ", argv[1]);

    if(close(in_fd) == -1 || close(out_fd) == -1) 
        oops("Error closing files", "");
    return 0;
}

void oops(char * s1, char * s2) {
    fprintf(stderr, "Error: %s ", s1);
    perror(s2);
    exit(1);
}
