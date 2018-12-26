#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <dirent.h>
#include <string.h>
#include <sys/stat.h>
#include <time.h>

int stringcmp(const void * s1, const void * s2) {
    return strcmp(*(char**) s1, *(char**) s2);
}
void do_ls(char * dir) {
    DIR * path;
    if( (path = opendir(dir)) == NULL ) {
        perror("Error when opening dir: ");
        return;
    }
    struct dirent * p = NULL;
    char *filenames[1000];
    int count = 0;
    while( (p = readdir(path)) ) {
        if(strcmp(p->d_name, ".") == 0 || strcmp(p->d_name, "..") == 0) continue;
        filenames[count++] = p->d_name;
    }
    qsort(filenames, count, sizeof(char*), stringcmp);
    for(int i=0; i < count/5; ++i) {
        for(int j=0; j<5; ++j) {
            if(i*5+j < count)
                printf("%8.8s ", filenames[i*5+j]);
            else break;
        }
        printf("\n");
    }
    if( closedir(path) == -1 ) 
        perror("Error when closing dir: ");
    return;
}

void do_ls_a(char * dir) {
    DIR * path = opendir(dir);
    if(path == NULL) {
        perror("Error when opening dir: ");
        return;
    }
    struct dirent * p = NULL;
    while( (p = readdir(path)) ) {
        struct stat info;
        if(stat(p->d_name, &info) == 0) {
            printf("%8.8s\t%8d\t%d\t%s", p->d_name, info.st_size, info.st_nlink, ctime(&info.st_atime));
        }
    }
    if( closedir(path) == -1 ) {
        perror("Error when closing dir: ");
    }
    return;
}

int main(int argc, char * argv[]) {
    if( argc == 1 ) do_ls_a(".");
    else 
        while(--argc) {
            printf("%s: \n", *++argv);
            do_ls(*argv);
        }
}