#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <fcntl.h>

#define BUFSIZE 1024
void do_cat(FILE * fp) {
    char buf[BUFSIZE];
    while(fgets(buf, BUFSIZE, fp)) {
        fputs(buf, stdout);
    }
    return;
}

int main(int argc, char * argv[]) {
    if( argc == 1 ) 
        do_cat(stdin);
    else {
        while( --argc > 0) {
            FILE * fp = fopen(*++argv, "r");
            if( fp == NULL ) {
                perror("Opening File Error: ");
                exit(1);
            }
            do_cat(fp);
        }
    }
}